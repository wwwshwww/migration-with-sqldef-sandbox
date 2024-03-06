package secure_hasher_test

import (
	"example_app/sample_auth_app/domain_service/secure_hasher"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService(t *testing.T) {
	s := secure_hasher.New()
	passA := "pass1234"
	passB := "another"

	hashed1, err := s.Hash(passA)
	assert.NoError(t, err)
	hashed2, err := s.Hash(passA, secure_hasher.WithCost(5))
	assert.NoError(t, err)
	hashed3, err := s.Hash(passB)
	assert.NoError(t, err)

	assert.True(t, s.IsSame(hashed1, passA))
	assert.True(t, s.IsSame(hashed2, passA))
	assert.False(t, s.IsSame(hashed3, passA))
	assert.True(t, s.IsSame(hashed3, passB))
}
