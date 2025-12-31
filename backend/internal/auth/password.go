package auth

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a password using bcrypt
func HashPassword(password string) (string, error) {
    // GenerateFromPassword hashes the password
    // 10 is the "cost" (higher = slower = more secure but slower to compute)
    hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
    if err != nil {
        return "", fmt.Errorf("failed to hash password: %w", err)
    }
    return string(hash), nil
}

// CheckPassword compares a password with its hash
func CheckPassword(hashedPassword, password string) bool {
    // CompareHashAndPassword returns nil if they match
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    return err == nil
}