package services

import (
	"fmt"
	"regexp"

	"github.com/cakeru/kanago-backend/internal/auth"
	"github.com/cakeru/kanago-backend/internal/config"
	"github.com/cakeru/kanago-backend/internal/models"
	"github.com/cakeru/kanago-backend/internal/repository"
)

// AuthService handles authentication-related operations
type AuthService struct {
	userRepo *repository.UserRepository
	cfg      *config.Config
}

// NewAuthService creates a new AuthService
func NewAuthService(cfg *config.Config) *AuthService {
	return &AuthService{
		userRepo: repository.NewUserRepository(),
		cfg:      cfg,
	}
}

type RegisterRequest struct {
	Email   string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	User        *models.User `json:"user"`
	AccessToken string        `json:"access_token"`
	RefreshToken string       `json:"refresh_token"`
}

//Register creates a new user and returns auth tokens
func (s *AuthService) Register(req *RegisterRequest) (*AuthResponse, error) {
	// Validate input
	if err := s.validateRegisterInput(req); err != nil {
		return nil, err
	}

	// Check if email already exists
	exists, err := s.userRepo.CheckEmailExists(req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to check email existence: %w", err)
	}
	if exists {
		return nil, fmt.Errorf("email already registered")
	}

	// Hash password
	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// Create user model
	user := &models.User{
		Email:       req.Email,
		Username:    req.Username,
		PasswordHash: hashedPassword,
	}

	//Save user to database
	if err := s.userRepo.CreateUser(user); err != nil {
		return nil, err
	}

	//Generate tokens
	accessToken, refreshToken, err := auth.GenerateTokens(user, s.cfg)
	if err != nil {
		return nil, err
	}

	// Return response (don't include password hash!)
    return &AuthResponse{
        User: &models.User{
            ID:       user.ID,
            Email:    user.Email,
            Username: user.Username,
        },
        AccessToken:  accessToken,
        RefreshToken: refreshToken,
    }, nil
}

//Login authenticates a user and returns auth tokens
func (s *AuthService) Login(req *LoginRequest) (*AuthResponse, error){
	//Validate Input
	if err := s.validateLoginInput(req); err != nil {
		return nil, err
	}

	//Get user from database
	user, err := s.userRepo.GetUserByEmail(req.Email)
	if err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	//Check password
	if !auth.CheckPassword(user.PasswordHash, req.Password){
		return nil, fmt.Errorf("invalid email or password")
	}
	//Generate tokens
	accessToken, refreshToken, err := auth.GenerateTokens(user, s.cfg)
	if err != nil {
		return nil, err
	}
	
	return &AuthResponse{
		User: &models.User{
			ID:       user.ID,
			Email:    user.Email,
			Username: user.Username,
		},
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

//RefreshAccessToken generates a new access token using a valid refresh token
func (s *AuthService) RefreshAccessToken(refreshToken string) (string, error){
	//Verify refresh token
	claims, err := auth.VerifyToken(refreshToken, s.cfg)
	if err != nil {
		return "", fmt.Errorf("invalid refresh token: %w", err)
	}

	//Get user from database
	user, err := s.userRepo.GetUserByID(claims.UserID)
	if err != nil {
		return "", fmt.Errorf("failed to get user: %w", err)
	}
	
	//Generate new access token
	accessToken, _, err := auth.GenerateTokens(user, s.cfg)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func (s *AuthService) validateRegisterInput(req *RegisterRequest) error {
    if req.Email == "" {
        return fmt.Errorf("email is required")
    }
    if !s.isValidEmail(req.Email) {
        return fmt.Errorf("invalid email format")
    }

    if req.Username == "" {
        return fmt.Errorf("username is required")
    }
    if len(req.Username) < 3 || len(req.Username) > 50 {
        return fmt.Errorf("username must be between 3 and 50 characters")
    }

    if req.Password == "" {
        return fmt.Errorf("password is required")
    }
    if len(req.Password) < 8 {
        return fmt.Errorf("password must be at least 8 characters")
    }

    return nil
}

func (s *AuthService) validateLoginInput(req *LoginRequest) error {
	if req.Email == "" {
		return fmt.Errorf("email is required")
	}
	if req.Password == "" {
		return fmt.Errorf("password is required")
	}
	return nil
}

func (s *AuthService) isValidEmail(email string) bool {
    re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
    return re.MatchString(email)
}