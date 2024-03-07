package user_finder

import "example_app/sample_auth_app/domain/user/user"

//go:generate moq -out ./$GOFILE.moq.go . Finder
type Finder interface {
	Find(fo FilteringOptions, so SortingOptions) ([]user.ID, error)
}
