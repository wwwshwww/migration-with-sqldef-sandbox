package session_finder

import "example_app/sample_auth_app/domain/session/session"

type Finder interface {
	Find(fo FilteringOptions, so SortingOptions) ([]session.ID, error)
}
