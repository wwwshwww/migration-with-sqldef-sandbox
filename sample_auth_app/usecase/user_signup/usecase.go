package user_signup

import (
	"context"
	"example_app/sample_auth_app/domain/user/user"

	"github.com/go-errors/errors"
)

type Usecase interface {
	SignUp(InputDTO) error
}

type usecase struct {
	ctx            context.Context
	userRepository user.Repository
}

func New(ctx context.Context, userRepository user.Repository) Usecase {
	return &usecase{
		ctx:            ctx,
		userRepository: userRepository,
	}
}

func (uc *usecase) SignUp(input InputDTO) error {
	userId := user.GenerateID()
	name, err := user.NewName(input.Name)
	if err != nil {
		return errors.Wrap(err, 1)
	}
	password, err := user.NewPassword(input.Password)
	if err != nil {
		return errors.Wrap(err, 1)
	}
	u, err := user.New(userId, name, password)
	if err != nil {
		return errors.Wrap(err, 1)
	}
	if err := uc.userRepository.BulkSave([]user.User{u}); err != nil {
		return errors.Wrap(err, 1)
	}
	return nil
}
