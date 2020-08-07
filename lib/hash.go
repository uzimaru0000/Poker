package lib

import "golang.org/x/crypto/bcrypt"

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
