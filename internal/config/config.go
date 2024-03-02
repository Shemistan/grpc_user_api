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

// Tes тестовый
type Tes interface {
	Secret() string
}

// Load загрузить кофиги
func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}
