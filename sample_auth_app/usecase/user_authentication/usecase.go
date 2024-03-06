package user_authentication

import (
	"context"
	"example_app/sample_auth_app/domain/session/session"
	"example_app/sample_auth_app/domain/user/user"
	"example_app/sample_auth_app/domain/user/user_finder"
	"example_app/sample_auth_app/domain_service/password_hasher"
	"time"

	"github.com/go-errors/errors"
)

type Usecase interface {
	SignUp(SignUpInput) error
	SignIn(SignInInput) (SessionInfo, error)
}

type usecase struct {
	ctx                   context.Context
	userRepository        user.Repository
	userFinder            user_finder.Finder
	sessionRepository     session.Repository
	passwordHasherService password_hasher.PasswordHasher
}

func New(ctx context.Context, userRepository user.Repository) Usecase {
	return &usecase{
		ctx:            ctx,
		userRepository: userRepository,
	}
}

func (uc *usecase) SignUp(input SignUpInput) error {
	userId := user.GenerateID()
	name, err := user.NewName(input.Name)
	if err != nil {
		return errors.Wrap(err, 1)
	}
	password, err := user.NewPassword(input.Password)
	if err != nil {
		return errors.Wrap(err, 1)
	}
	hashedPassword, err := uc.passwordHasherService.Hash(password.Primitive())
	if err != nil {
		return errors.Wrap(err, 1)
	}
	u, err := user.New(userId, name, hashedPassword)
	if err != nil {
		return errors.Wrap(err, 1)
	}
	if err := uc.userRepository.BulkSave([]user.User{u}); err != nil {
		return errors.Wrap(err, 1)
	}
	return nil
}

func (uc *usecase) SignIn(input SignInInput) (SessionInfo, error) {
	// ?: probably move auth process to domain_service

	inputPassword, err := user.NewPassword(input.Password)
	if err != nil {
		return SessionInfo{}, errors.WrapPrefix(err, "Invalid password", 1)
	}

	userIds, err := uc.userFinder.Find(
		user_finder.FilteringOptions{
			NameExact: &input.Name,
		},
		user_finder.SortingOptions{},
	)
	if err != nil {
		return SessionInfo{}, errors.WrapPrefix(err, "Failed to find user", 1)
	}
	if len(userIds) == 0 {
		return SessionInfo{}, errors.New("User not found")
	}

	loginUserID := userIds[0]
	loginUsers, err := uc.userRepository.BulkGet([]user.ID{loginUserID})
	if err != nil {
		return SessionInfo{}, errors.WrapPrefix(err, "Failed to get user", 1)
	}
	loginUser := loginUsers[0]
	if !uc.passwordHasherService.IsSame(loginUser.HashedPassword(), inputPassword.Primitive()) {
		return SessionInfo{}, errors.New("Wrong password")
	}

	sId := session.GenerateID()
	sToken, err := session.GenerateToken()
	if err != nil {
		return SessionInfo{}, errors.WrapPrefix(err, "Failed to generate session token", 1)
	}
	ss, err := session.New(sId, loginUserID, sToken, time.Now(), time.Now())
	if err != nil {
		return SessionInfo{}, errors.WrapPrefix(err, "Failed to create session", 1)
	}
	if err := uc.sessionRepository.BulkSave([]session.Session{ss}); err != nil {
		return SessionInfo{}, errors.WrapPrefix(err, "Failed to save session", 1)
	}
	return SessionInfo{
		Id:        ss.ID().Primitive(),
		Token:     ss.Token().Primitive(),
		CreatedAt: ss.CreatedAt(),
		ExpiresAt: ss.ExpiresAt(),
	}, nil
}
