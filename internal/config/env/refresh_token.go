package env

import (
	"errors"
	"os"

	"github.com/Shemistan/grpc_user_api/internal/config"
)

const (
	refreshSecretEnvName = "REFRESH_KEY"
)

type secretRefreshTokenConfig struct {
	secret string
}

// NewSecretRefreshTokenConfig - получить кофиг токена
func NewSecretRefreshTokenConfig() (config.SecretRefreshTokenConfig, error) {
	secret := os.Getenv(refreshSecretEnvName)
	if len(secret) == 0 {
		return nil, errors.New("refresh secret key not found")
	}

	return &secretRefreshTokenConfig{
		secret: secret,
	}, nil
}

func (cfg *secretRefreshTokenConfig) SecretKey() string {
	return cfg.secret
}
