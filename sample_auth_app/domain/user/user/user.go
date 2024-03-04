package user

import "github.com/go-errors/errors"

type User interface {
	ID() ID
	Name() Name
	Password() Password

	ChangeName(Name) error
	ChangePassword(Password) error
}

type user struct {
	id       ID
	name     Name
	password Password
}

func New(i ID, n Name, p Password) (User, error) {
	return &user{id: i, name: n, password: p}, nil
}

func (e *user) ID() ID {
	return e.id
}

func (e *user) Name() Name {
	return e.name
}

func (e *user) Password() Password {
	return e.password
}

func (e *user) ChangeName(v Name) error {
	e.name = v
	return nil
}

func (e *user) ChangePassword(v Password) error {
	if e.password.Equal(v) {
		return errors.New("New password must be different from old one")
	}
	e.password = v
	return nil
}
