package utilities

import "golang.org/x/crypto/bcrypt"

func Hash(value string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(value), 14)
	return string(bytes), err
}

func HashCompare(value, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(value))
	return err == nil
}
