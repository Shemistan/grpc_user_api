package env

import (
	"errors"
	"os"

	"github.com/Shemistan/grpc_user_api/internal/config"
)

const (
	dsnEnvName = "PG_DSN"
)

var _ config.PGConfig = (*pgConfig)(nil)

type pgConfig struct {
	dsn string
}

// NewPGConfig - получить
func NewPGConfig() (*pgConfig, error) {
	dsn := os.Getenv(dsnEnvName)
	if len(dsn) == 0 {
		return nil, errors.New("pg dsn not found")
	}

	return &pgConfig{
		dsn: dsn,
	}, nil
}

func (cfg *pgConfig) DSN() string {
	return "host=localhost port=54322 dbname=grpc user=grpc password=grpc sslmode=disable"
}
