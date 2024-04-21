package config

import (
	"github.com/joho/godotenv"
	"time"
)

const (
	RefreshTokenExpiration = 3 * 24 * time.Hour
	AccessTokenExpiration  = 5 * time.Hour
)

// GRPCConfig конфиг
type GRPCConfig interface {
	Address() string
}

// GRPCAuthConfig конфиг
type GRPCAuthConfig interface {
	Address() string
}

// GRPCAccessConfig конфиг
type GRPCAccessConfig interface {
	Address() string
}

// PGConfig конфиг
type PGConfig interface {
	DSN() string
}

// SecretHashConfig конфиг
type SecretHashConfig interface {
	PasswordHashKey() string
}

// Load загрузить кофиги
func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}

// HTTP конфиг
type HTTP interface {
	Address() string
}

// Swagger конфиг
type Swagger interface {
	Address() string
}

// SecretRefreshTokenConfig конфиг
type SecretRefreshTokenConfig interface {
	SecretKey() string
}

// SecretAccessTokenConfig конфиг
type SecretAccessTokenConfig interface {
	SecretKey() string
}

type TokenServiceConfig struct {
	RefreshTokenSecretKey  string
	AccessTokenSecretKey   string
	RefreshTokenExpiration time.Duration
	AccessTokenExpiration  time.Duration
}
