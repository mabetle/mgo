package mpasswd

import (
	"golang.org/x/crypto/bcrypt"
)

// GenerateHashPassword
func GenerateHashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	logger.Debugf("Length Hashed Password: %d ", len(hashedPassword))
	return string(hashedPassword), nil
}

// CompareHashAndPassword
func CompareHashAndPassword(encryptPasswd, passwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(encryptPasswd), []byte(passwd))
}
