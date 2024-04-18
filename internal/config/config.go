package config

import (
	"github.com/joho/godotenv"
)

// GRPCConfig конфиг
type GRPCConfig interface {
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
