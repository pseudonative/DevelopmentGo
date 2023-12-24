package api

import "testing"

func TestGenerateJWT(t *testing.T) {
	username := "testuser"
	token, err := GenerateJWT(username)
	if err != nil {
		t.Errorf("Failed to generate JWT: %v", err)
	}
	claims, isValid := ValidateToken(token)
	if !isValid || claims.Username != username {
		t.Errorf("Failed to validate JWT or incorrect username")
	}
}
