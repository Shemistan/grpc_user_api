package hasher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPasswordHash(t *testing.T) {
	secret := "test"
	serv := New(secret)
	password := "qwer"

	hash, err := serv.GetPasswordHash(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)

	t.Run("check password hash mismatch", func(t *testing.T) {
		wrongPassword := "wrong"
		match := serv.CheckPassword(wrongPassword, hash)
		assert.False(t, match)
	})
}
