package user_authentication

import "time"

type SignUpInput struct {
	Name     string
	Password string
}

type SignInInput struct {
	Name     string
	Password string
}

type SessionInfo struct {
	Id        string
	Token     string
	CreatedAt time.Time
	ExpiresAt time.Time
}
