package user_authentication

import (
	"example_app/sample_auth_app/common/mycontext"
	"example_app/sample_auth_app/domain/session/session"
	"example_app/sample_auth_app/domain/user/user"
	"example_app/sample_auth_app/domain/user/user_finder"
	"example_app/sample_auth_app/domain_service/secure_hasher"
	"example_app/sample_auth_app/domain_service/user_validation_service"
	"time"

	"github.com/go-errors/errors"
)

type Usecase interface {
	SignUp(SignUpInput) error
	SignIn(SignInInput) (SessionInfo, error)
	SignOut() error
}

type usecase struct {
	ctx                   mycontext.Context
	userRepository        user.Repository
	userFinder            user_finder.Finder
	sessionRepository     session.Repository
	userValidationService user_validation_service.Service
	secureHasherService   secure_hasher.SecureHasher
}

func New(
	ctx mycontext.Context,
	userRepository user.Repository,
	userFinder user_finder.Finder,
	sessionRepository session.Repository,
	userValidationService user_validation_service.Service,
	secureHasherService secure_hasher.SecureHasher,
) Usecase {
	return &usecase{
		ctx:                   ctx,
		userRepository:        userRepository,
		userFinder:            userFinder,
		sessionRepository:     sessionRepository,
		userValidationService: userValidationService,
		secureHasherService:   secureHasherService,
	}
}

func (uc *usecase) SignUp(input SignUpInput) error {
	userId := user.GenerateID()
	name, err := user.NewName(input.Name)
	if err != nil {
		return errors.WrapPrefix(err, "Failed to create name", 1)
	}
	password, err := user.NewPassword(input.Password)
	if err != nil {
		return errors.WrapPrefix(err, "Failed to create password", 1)
	}
	u, err := user.Generate(userId, name, password, uc.secureHasherService)
	if err != nil {
		return errors.WrapPrefix(err, "Failed to generate user", 1)
	}
	if isDup, err := uc.userValidationService.CheckDuplicatedInExtSource([]user.User{u}); err != nil {
		return errors.WrapPrefix(err, "Failed to check duplication in external source", 1)
	} else if isDup[userId] {
		return errors.New("Duplicated user name")
	}
	if err := uc.userRepository.Save(u); err != nil {
		return errors.WrapPrefix(err, "Failed to save user", 1)
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
	loginUser, err := uc.userRepository.Get(loginUserID)
	if err != nil {
		return SessionInfo{}, errors.WrapPrefix(err, "Failed to get user", 1)
	}
	if !uc.secureHasherService.IsSame(loginUser.HashedPassword(), inputPassword.Primitive()) {
		return SessionInfo{}, errors.New("Wrong password")
	}

	s, sToken, err := session.Generate(loginUserID, time.Now(), uc.secureHasherService)
	if err != nil {
		return SessionInfo{}, errors.WrapPrefix(err, "Failed to generate session", 1)
	}
	if err := uc.sessionRepository.Save(s); err != nil {
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
	stoken, ok := uc.ctx.SessionToken()
	if !ok {
		return errors.New("missing token")
	}

	targetSession, err := uc.sessionRepository.Get(session.NewID(sid))
	if err != nil {
		return errors.WrapPrefix(err, "session repository error", 1)
	}
	if !uc.secureHasherService.IsSame(targetSession.HashedToken(), stoken) {
		// TODO: needs to be able to handle errors more correctly, using ex. generics
		return errors.New("invalid token")
	}

	now := time.Now()
	if !targetSession.IsValid(now) {
		return nil
	}
	targetSession.Invalidate(now)
	if err := uc.sessionRepository.Save(targetSession); err != nil {
		return errors.WrapPrefix(err, "failed to save session", 1)
	}
	return nil
}
