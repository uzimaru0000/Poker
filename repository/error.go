package repository

import (
	"fmt"
)

type NotFound struct {
	ID string
}

func (e *NotFound) Error() string {
	return fmt.Sprintf("%s is NotFound", e.ID)
}

func (e *NotFound) Is(target error) bool {
	ee, ok := target.(*NotFound)
	return ok && ee.ID == e.ID
}

type AlreadyRegistered struct {
	ID string
}

func (e *AlreadyRegistered) Error() string {
	return fmt.Sprintf("%s is Already registered", e.ID)
}

func (e *AlreadyRegistered) Is(target error) bool {
	ee, ok := target.(*AlreadyRegistered)
	return ok && ee.ID == e.ID
}
