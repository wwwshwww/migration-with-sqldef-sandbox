package session

import (
	"example_app/sample_auth_app/domain/user/user"
	"time"

	"github.com/go-errors/errors"
)

const (
	ExpiresDuration = time.Hour * 6  // 最後のアクティビティからセッションが切れるまでの時間
	LifeSpan        = time.Hour * 48 // セッションの最大寿命
)

type Session interface {
	ID() ID
	UserID() user.ID
	Token() Token
	CreatedAt() time.Time
	ActivitiesAt() time.Time
	ExpiresAt() time.Time

	UpdateActivity(time.Time)
	IsExpired(time.Time) bool
}

type session struct {
	id           ID
	userId       user.ID
	token        Token
	createdAt    time.Time
	activitiesAt time.Time
}

func New(i ID, u user.ID, t Token, createdAt, activitiesAt time.Time) (Session, error) {
	if activitiesAt.Before(createdAt) {
		return nil, errors.New("Invalid period")
	}
	return &session{
		id:           i,
		userId:       u,
		token:        t,
		createdAt:    createdAt,
		activitiesAt: activitiesAt,
	}, nil
}

func Generate(u user.ID, now time.Time) Session {
	id := GenerateID()
	token, err := GenerateToken()
	if err != nil {
		panic(err)
	}

	s, err := New(id, u, token, now, now)
	if err != nil {
		panic(err)
	}
	return s
}

func (e *session) ID() ID {
	return e.id
}
func (e *session) UserID() user.ID {
	return e.userId
}
func (e *session) Token() Token {
	return e.token
}
func (e *session) CreatedAt() time.Time {
	return e.createdAt
}
func (e *session) ActivitiesAt() time.Time {
	return e.activitiesAt
}
func (e *session) ExpiresAt() time.Time {
	expectedExpiresAt := e.activitiesAt.Add(ExpiresDuration)
	if total := expectedExpiresAt.Sub(e.createdAt); total > LifeSpan {
		return e.ActivitiesAt().Add(LifeSpan - e.activitiesAt.Sub(e.createdAt))
	}
	return expectedExpiresAt
}
func (e *session) UpdateActivity(now time.Time) {
	e.activitiesAt = now
}
func (e *session) IsExpired(now time.Time) bool {
	return now.After(e.ExpiresAt())
}
