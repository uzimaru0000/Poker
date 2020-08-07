package lib

import "github.com/google/uuid"

type UUIDGenerator interface {
	Generate() (string, error)
}

type UUID struct{}

func (*UUID) Generate() (string, error) {
	id, err := uuid.NewUUID()

	if err != nil {
		return "", err
	}

	return id.String(), nil
}
