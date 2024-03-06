package secure_hasher

import "golang.org/x/crypto/bcrypt"

type SecureHasher interface {
	Hash(plaintext string, options ...hashOption) (string, error)
	IsSame(hash, password string) bool
}

type secureHasher struct{}

func New() SecureHasher {
	return &secureHasher{}
}

type hashConfig struct {
	Cost int
}

type hashOption func(*hashConfig)

func WithCost(cost int) hashOption {
	return func(c *hashConfig) {
		c.Cost = cost
	}
}

func (p *secureHasher) Hash(plaintext string, options ...hashOption) (string, error) {
	config := hashConfig{
		Cost: bcrypt.DefaultCost,
	}
	for _, opt := range options {
		opt(&config)
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(plaintext), config.Cost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func (p *secureHasher) IsSame(hash, plaintext string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plaintext)) == nil
}
