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
	HashedToken() string // 秘匿化されたセキュリティトークン。復号化不可。平文との比較は可能
	CreatedAt() time.Time
	ActivitiesAt() time.Time
	ExpiresAt() time.Time

	UpdateActivity(activitiesAt time.Time)
	Invalidate(now time.Time)

	IsValid(now time.Time) bool
}

type session struct {
	id           ID
	userId       user.ID
	hashedToken  string
	createdAt    time.Time
	activitiesAt time.Time
	expiresAt    time.Time
}

func New(i ID, u user.ID, hashedToken string, createdAt, activitiesAt, expiresAt time.Time) (Session, error) {
	if activitiesAt.Before(createdAt) {
		return nil, errors.New("Invalid period")
	}
	if expiresAt.Sub(createdAt) > LifeSpan {
		return nil, errors.New("too long period")
	}
	return &session{
		id:           i,
		userId:       u,
		hashedToken:  hashedToken,
		createdAt:    createdAt,
		activitiesAt: activitiesAt,
		expiresAt:    expiresAt,
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

	s, err := New(id, u, hashedToken, now, now, now.Add(ExpiresDuration))
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
	return e.expiresAt
}

func (e *session) UpdateActivity(activitiesAt time.Time) {
	e.activitiesAt = activitiesAt

	expectedExpiresAt := e.activitiesAt.Add(ExpiresDuration)
	if total := expectedExpiresAt.Sub(e.createdAt); total > LifeSpan {
		e.expiresAt = e.ActivitiesAt().Add(LifeSpan - e.activitiesAt.Sub(e.createdAt))
	} else {
		e.expiresAt = expectedExpiresAt
	}
}
func (e *session) Invalidate(now time.Time) {
	e.expiresAt = now
	e.activitiesAt = now
}
func (e *session) IsValid(now time.Time) bool {
	return now.Before(e.expiresAt)
}
