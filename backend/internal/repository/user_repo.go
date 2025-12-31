package repository

import (
	"fmt"

	"github.com/cakeru/kanago-backend/internal/db"
	"github.com/cakeru/kanago-backend/internal/models"
)

type UserRepository struct {
	// Empty - we use the global DB from db package
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	query := `INSERT INTO users (email, username, password_hash, created_at, updated_at)
			  VALUES ($1, $2, $3, NOW(), NOW())
			  RETURNING id, created_at, updated_at`


	err := db.DB.QueryRow(query, user.Email, user.Username, user.PasswordHash).
		Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}

	query := `SELECT id, email, username, password_hash, created_at, updated_at
			  FROM users WHERE email = $1`
	
	err := db.DB.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Username, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}
	return user, nil
}

func (r *UserRepository) GetUserByID(id int) (*models.User, error) {
	user := &models.User{}

	query := `SELECT id, email, username, password_hash, created_at, updated_at FROM users WHERE id = $1`

	err := db.DB.QueryRow(query, id).Scan(&user.ID, &user.Email, &user.Username, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by ID: %w", err)
	}

	return user, nil
}

func (r *UserRepository) CheckEmailExists(email string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`

	err := db.DB.QueryRow(query, email).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check email existence: %w", err)
	}
	return exists, nil
}