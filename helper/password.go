package helper

import "golang.org/x/crypto/bcrypt"

func HashingPassword(password string) (string, error) {
	hashByte, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	if err != nil {
		return "", err
	}
	return string(hashByte), nil
}

func VerifyPassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}
