package utils

import "golang.org/x/crypto/bcrypt"

// hashPassword generates a bcrypt hash for the given password
func HashPassword(password string) (string, error) {
	// Generate a hashed password using bcrypt
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// verifyPassword checks if the given password matches the hashed password
func VerifyPassword(password, hashedPassword string) error {
	// Compare the provided password with the hashed password
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
