package env

import (
	"errors"
	"os"

	"github.com/Shemistan/grpc_user_api/internal/config"
)

const (
	accessSecretEnvName = "ACCESS_KEY"
)

type secretAccessTokenConfig struct {
	secret string
}

// NewSecretAccessTokenConfig - получить кофиг токена
func NewSecretAccessTokenConfig() (config.SecretAccessTokenConfig, error) {
	secret := os.Getenv(accessSecretEnvName)
	if len(secret) == 0 {
		return nil, errors.New("access secret key not found")
	}

	return &secretAccessTokenConfig{
		secret: secret,
	}, nil
}

func (cfg *secretAccessTokenConfig) SecretKey() string {
	return cfg.secret
}
