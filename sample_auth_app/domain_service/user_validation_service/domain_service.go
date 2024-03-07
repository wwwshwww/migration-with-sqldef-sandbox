package user_validation_service

import "example_app/sample_auth_app/domain/user/user"

type Service interface {
	CheckDuplicatedInExtSource(targets []user.User) (map[user.ID]bool, error)
}

type service struct {
	port Port
}

func New(port Port) Service {
	return service{port}
}

type dupModel struct {
	Name string
}

func (s service) CheckDuplicatedInExtSource(targets []user.User) (map[user.ID]bool, error) {
	return nil, nil
}
