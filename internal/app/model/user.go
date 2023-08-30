package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID                int
	Email             string
	Password          string
	EncryptedPassword string
}

func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
