package user_authentication

import (
	"example_app/sample_auth_app/common/mycontext"
	"example_app/sample_auth_app/domain/session/session"
	"example_app/sample_auth_app/domain/user/user"
	"example_app/sample_auth_app/domain/user/user_finder"
	"example_app/sample_auth_app/domain_service/secure_hasher"
	"time"

	"github.com/go-errors/errors"
)

type Usecase interface {
	SignUp(SignUpInput) error
	SignIn(SignInInput) (SessionInfo, error)
	SignOut() error
}

type usecase struct {
	ctx                 mycontext.Context
	userRepository      user.Repository
	userFinder          user_finder.Finder
	sessionRepository   session.Repository
	secureHasherService secure_hasher.SecureHasher
}

func New(ctx mycontext.Context, userRepository user.Repository) Usecase {
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
	hashedPassword, err := uc.secureHasherService.Hash(password.Primitive())
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
	if !uc.secureHasherService.IsSame(loginUser.HashedPassword(), inputPassword.Primitive()) {
		return SessionInfo{}, errors.New("Wrong password")
	}

	s, sToken, err := session.Generate(loginUserID, time.Now(), uc.secureHasherService)
	if err != nil {
		return SessionInfo{}, errors.WrapPrefix(err, "Failed to generate session", 1)
	}
	if err := uc.sessionRepository.BulkSave([]session.Session{s}); err != nil {
		return SessionInfo{}, errors.WrapPrefix(err, "Failed to save session", 1)
	}
	return SessionInfo{
		Id:        s.ID().Primitive(),
		Token:     sToken.Primitive(),
		CreatedAt: s.CreatedAt(),
		ExpiresAt: s.ExpiresAt(),
	}, nil
}

func (uc *usecase) SignOut() error {
	sid, ok := uc.ctx.SessionID()
	if !ok {
		return nil
	}
	ss, err := uc.sessionRepository.BulkGet([]session.ID{session.NewID(sid)})
	if err != nil {
		return errors.WrapPrefix(err, "session repository error", 1)
	}
	if len(ss) == 0 {
		return nil
	}
	targetSession := ss[0]
	stoken, ok := uc.ctx.SessionToken()
	if !ok {
		return errors.WrapPrefix(err, "missing token", 1)
	}
	if !uc.secureHasherService.IsSame(targetSession.HashedToken(), stoken) {
		// TODO: needs to be able to handle errors more correctly, using ex. generics
		return errors.New("invalid token")
	}

	now := time.Now()
	if !targetSession.IsValid(now) {
		return nil
	}

	targetSession.Invalidate(time.Now())
	if err := uc.sessionRepository.BulkSave([]session.Session{targetSession}); err != nil {
		return errors.WrapPrefix(err, "failed to save session", 1)
	}
	return nil
}
