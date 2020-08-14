package lib

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type HashGenerator interface {
	Create(string) (string, error)
	Verify(string, string) error
}

type Hash struct{}

func (*Hash) Create(str string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), err
}

func (*Hash) Verify(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func TokenValidation(token string) bool {
	return strings.HasPrefix(token, "Bearer")
}

func ExtractingToken(raw string) (string, error) {
	splited := strings.Split(raw, " ")

	if len(splited) < 2 {
		return "", errors.New("Invalid Token")
	}

	return splited[1], nil
}

