package user

import "github.com/go-errors/errors"

type Name interface {
	Primitive() string
}

type name string

const NameLengthMax, NameLengthMin = 12, 1

func NewName(v string) (Name, error) {
	if l := len([]rune(v)); l < NameLengthMin || l > NameLengthMax {
		return nil, errors.New("Invalid Name")
	}
	return name(v), nil
}

func (v name) Primitive() string {
	return string(v)
}
