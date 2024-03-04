package session

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/go-errors/errors"
)

type Token interface {
	Primitive() string
}

type token string

const TokenByteLength = 32

func NewToken(v string) Token {
	return token(v)
}

func GenerateToken() (Token, error) {
	b := make([]byte, TokenByteLength)
	if _, err := rand.Read(b); err != nil {
		return nil, errors.New("Failed to sample byte row")
	}
	return NewToken(base64.URLEncoding.EncodeToString(b)), nil
}

func (v token) Primitive() string {
	return string(v)
}
