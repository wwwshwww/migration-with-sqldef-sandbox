package session_finder

import "example_app/sample_auth_app/domain/session/session"

//go:generate moq -out ./$GOFILE.moq.go . Finder
type Finder interface {
	Find(fo FilteringOptions, so SortingOptions) ([]session.ID, error)
}
