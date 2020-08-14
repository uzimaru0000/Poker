package impl

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/uzimaru0000/poker/usecase"
)

type tokenizer struct {
	secret string
}

type Claim struct {
	Subject string
}

func NewTokenizer(secret string) usecase.Tokenizer {
	return &tokenizer{secret: secret}
}

func (t *tokenizer) CreateToken(claim interface{}) (string, error) {
	claims, ok := claim.(Claim)
	if !ok {
		return "", errors.New("This function takes the Claim type")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Subject:   claims.Subject,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(t.secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (t *tokenizer) VerifyToken(tokenString string) (interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.secret), nil
	})

	if err != nil {
		return nil, err
	}

	return decodeClain(token.Claims)
}

func decodeClain(raw interface{}) (*Claim, error) {
	claims, ok := raw.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("Failed")
	}

	sub, ok := claims["sub"].(string)
	if !ok {
		return nil, errors.New("Not exist field for subject")
	}

	return &Claim{
		Subject: sub,
	}, nil
}
