package utils

import (
	_ "bytes"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func GetPasswordHash(pwd []byte) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, 14)
	return hash, err
}

func ComparePassowrds(hashedPassword, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	fmt.Println(err == nil)
	return err == nil
}
