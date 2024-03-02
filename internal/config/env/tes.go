package env

import (
	"os"

	"github.com/Shemistan/grpc_user_api/internal/config"
)

const (
	secret = "PASSWORD_HASH_KEY"
)

var _ config.Tes = (*tesConfig)(nil)

type tesConfig struct {
	secret string
}

// NewTesConfig - получить
func NewTesConfig() (*tesConfig, error) {
	sec := os.Getenv(secret)
	//if len(sec) == 0 {
	//	return nil, errors.New("secret not found")
	//}

	return &tesConfig{
		secret: sec,
	}, nil
}

func (cfg *tesConfig) Secret() string {
	return secret
}
