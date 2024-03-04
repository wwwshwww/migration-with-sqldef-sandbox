package user

import "github.com/go-errors/errors"

// ! deprecated to use in actual app
type Password interface {
	Primitive() string
	Equal(Password) bool
}

type password string

const PasswordLengthMax, PasswordLengthMin = 50, 8

func NewPassword(v string) (Password, error) {
	zero := rune('0')
	nine := rune('9')

	var totalCount, unexpectedCount int
	for _, c := range v {
		if c < zero && c > nine {
			unexpectedCount++
		}
		totalCount++
	}

	if totalCount < PasswordLengthMin || totalCount > PasswordLengthMax || unexpectedCount > 0 {
		return nil, errors.New("Invalid password")
	}

	return password(v), nil
}

func (v password) Equal(sub Password) bool {
	return v.Primitive() == sub.Primitive()
}

func (v password) Primitive() string {
	return string(v)
}
