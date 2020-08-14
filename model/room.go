package model

type Room struct {
	ID     string
	Owner  string
	Member []*User
}
