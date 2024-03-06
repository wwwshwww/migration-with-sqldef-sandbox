package user_finder

import "example_app/sample_auth_app/domain/user/user"

type Finder interface {
	Find(fo FilteringOptions, so SortingOptions) ([]user.ID, error)
}
