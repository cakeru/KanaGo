package auth

import (
	"testing"

	"github.com/cakeru/kanago-backend/internal/config"
	"github.com/cakeru/kanago-backend/internal/models"
)

func TestGenerateTokens(t *testing.T) {
    // Create test user
    user := &models.User{
        ID:       1,
        Email:    "test@example.com",
        Username: "testuser",
    }

    // Create test config
    cfg := &config.Config{
        JWT: config.JWTConfig{
            Secret:            "test-secret",
            ExpiryMinutes:     15,
            RefreshExpiryDays: 7,
        },
    }

    // Generate tokens
    accessToken, refreshToken, err := GenerateTokens(user, cfg)
    if err != nil {
        t.Fatalf("Failed to generate tokens: %v", err)
    }

    if accessToken == "" {
        t.Fatal("Access token is empty")
    }

    if refreshToken == "" {
        t.Fatal("Refresh token is empty")
    }

    // Verify tokens
    accessClaims, err := VerifyToken(accessToken, cfg)
    if err != nil {
        t.Fatalf("Failed to verify access token: %v", err)
    }

    if accessClaims.UserID != user.ID {
        t.Fatalf("Expected UserID %d, got %d", user.ID, accessClaims.UserID)
    }

    t.Log("✓ Tokens generated and verified successfully")
}

func TestCheckPassword(t *testing.T) {
    password := "mySecurePassword123"

    // Hash password
    hash, err := HashPassword(password)
    if err != nil {
        t.Fatalf("Failed to hash password: %v", err)
    }

    // Check password
    if !CheckPassword(hash, password) {
        t.Fatal("Password check failed for correct password")
    }

    // Check wrong password
    if CheckPassword(hash, "wrongpassword") {
        t.Fatal("Password check passed for wrong password")
    }

    t.Log("✓ Password hashing and verification working")
}