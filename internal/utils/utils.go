package utils

import (
	"github.com/Shemistan/grpc_user_api/internal/model"

	"time"
)

// Hasher - интерфейс хэшера паролей
type Hasher interface {
	GetPasswordHash(password string) (string, error)
	CheckPassword(hash, password string) bool
}

// TokenProvider - интерфейс обработчика токена
type TokenProvider interface {
	GenerateToken(info model.UserInfo, secretKey []byte, duration time.Duration) (string, error)
	VerifyToken(tokenStr string, secretKey []byte) (*model.UserClaims, error)
}
