package env

import (
	"errors"
	"os"

	"github.com/Shemistan/grpc_user_api/internal/config"
)

const (
	secretEnvName = "PASSWORD_HASH_KEY"
)

var _ config.SecretHashConfig = (*secretHashConfig)(nil)

type secretHashConfig struct {
	secret string
}

// NewSecretHashConfig - получить
func NewSecretHashConfig() (*secretHashConfig, error) {
	secret := os.Getenv(secretEnvName)
	if len(secret) == 0 {
		return nil, errors.New("password hash key not found")
	}

	return &secretHashConfig{
		secret: secret,
	}, nil
}

func (cfg *secretHashConfig) PasswordHashKey() string {
	return cfg.secret
}
