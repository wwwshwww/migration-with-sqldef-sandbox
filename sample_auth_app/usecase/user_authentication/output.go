package user_authentication

import "time"

type SessionInfo struct {
	Id        string
	Token     string
	CreatedAt time.Time
	ExpiresAt time.Time
}
