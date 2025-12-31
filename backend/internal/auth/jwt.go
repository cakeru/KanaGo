package auth

import (
	"fmt"
	"time"

	"github.com/cakeru/kanago-backend/internal/config"
	"github.com/cakeru/kanago-backend/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

//Claims represents the JWT claims
type Claims struct {
	UserID  int    `json:"user_id"`
	Email   string `json:"email"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

//GenerateTokens creates access and refresh tokens
func GenerateTokens(user *models.User, cfg *config.Config) (string, string, error){
	// Create access token claims
	accessClaims := &Claims{
		UserID:   user.ID,
		Email:    user.Email,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(cfg.JWT.ExpiryMinutes) * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:   "kanago-backend",
		},
	}
	//create access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString([]byte(cfg.JWT.Secret))
	if err != nil {
		return "", "", fmt.Errorf("failed to sign access token: %w", err)
	}

	//Create refresh token claims
	refreshClaims := &Claims{
		UserID:  user.ID,
		Email:   user.Email,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(cfg.JWT.RefreshExpiryDays) * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:   "kanago-backend",
		},
	}
	//create refresh token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(cfg.JWT.Secret))
	if err != nil {
		return "", "", fmt.Errorf("failed to sign refresh token: %w", err)
	}
	
	return accessTokenString, refreshTokenString, nil
}

func VerifyToken(tokenString string, cfg *config.Config) (*Claims, error) {
	//Parse token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token)(interface{}, error){
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfg.JWT.Secret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token claims")
	}
	return claims, nil
}
