package session

import (
	"example_app/sample_auth_app/domain/user/user"
	"example_app/sample_auth_app/domain_service/secure_hasher"
	"time"

	"github.com/go-errors/errors"
)

const (
	ExpiresDuration = time.Hour * 6  // 最後のアクティビティからセッションが切れるまでの時間
	LifeSpan        = time.Hour * 48 // セッションの最大寿命
	DefaultHashCost = 3
)

type Session interface {
	ID() ID
	UserID() user.ID
	HashedToken() string // 不可逆圧縮情報、平文との比較は可能
	CreatedAt() time.Time
	ActivitiesAt() time.Time
	ExpiresAt() time.Time

	UpdateActivity(time.Time)
	IsExpired(time.Time) bool
}

type session struct {
	id           ID
	userId       user.ID
	hashedToken  string
	createdAt    time.Time
	activitiesAt time.Time
}

func New(i ID, u user.ID, hashedToken string, createdAt, activitiesAt time.Time) (Session, error) {
	if activitiesAt.Before(createdAt) {
		return nil, errors.New("Invalid period")
	}
	return &session{
		id:           i,
		userId:       u,
		hashedToken:  hashedToken,
		createdAt:    createdAt,
		activitiesAt: activitiesAt,
	}, nil
}

func Generate(u user.ID, now time.Time, hasher secure_hasher.SecureHasher) (Session, Token, error) {
	id := GenerateID()
	token, err := GenerateToken()
	if err != nil {
		return nil, nil, errors.New("failed to generate Token")
	}
	hashedToken, err := hasher.Hash(token.Primitive(), secure_hasher.WithCost(DefaultHashCost))
	if err != nil {
		return nil, nil, errors.New("failed to hash Token")
	}

	s, err := New(id, u, hashedToken, now, now)
	if err != nil {
		return nil, nil, errors.New("failed to create session")
	}
	return s, token, nil
}

func (e *session) ID() ID {
	return e.id
}
func (e *session) UserID() user.ID {
	return e.userId
}
func (e *session) HashedToken() string {
	return e.hashedToken
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
