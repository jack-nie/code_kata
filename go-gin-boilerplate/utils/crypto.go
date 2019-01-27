package utils

import "golang.org/x/crypto/bcrypt"

func HashAndSalt(password []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePasswords(hashedPassword string, plainPassword []byte) (bool, error) {
	byteHash := []byte(hashedPassword)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		return false, err
	}
	return true, nil
}
