package session

import "github.com/google/uuid"

type ID interface {
	Primitive() string
}

type id string

func NewID(v string) ID {
	return id(v)
}

func GenerateID() ID {
	i := uuid.New()
	return NewID(i.String())
}

func (v id) Primitive() string {
	return string(v)
}
