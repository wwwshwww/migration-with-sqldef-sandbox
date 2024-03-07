package user

import (
	"example_app/sample_auth_app/domain_service/secure_hasher"

	"github.com/go-errors/errors"
)

type User interface {
	ID() ID
	Name() Name
	HashedPassword() string // 秘匿化されたパスワード。復号化不可。平文との比較は可能

	ChangeName(Name) error
	ChangePassword(Password, secure_hasher.SecureHasher) error
}

type user struct {
	id             ID
	name           Name
	hashedPassword string
}

func New(i ID, n Name, hashedPassword string) (User, error) {
	return &user{id: i, name: n, hashedPassword: hashedPassword}, nil
}

func (e *user) ID() ID {
	return e.id
}

func (e *user) Name() Name {
	return e.name
}

func (e *user) HashedPassword() string {
	return e.hashedPassword
}

func (e *user) ChangeName(v Name) error {
	e.name = v
	return nil
}

func (e *user) ChangePassword(newPassword Password, passwordHasher secure_hasher.SecureHasher) error {
	hashedPassword, err := passwordHasher.Hash(newPassword.Primitive())
	if err != nil {
		return errors.Wrap(err, 1)
	}
	e.hashedPassword = hashedPassword
	return nil
}
