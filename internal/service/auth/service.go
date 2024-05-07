package auth

import (
	"github.com/Shemistan/grpc_user_api/internal/config"
	def "github.com/Shemistan/grpc_user_api/internal/service"
	"github.com/Shemistan/grpc_user_api/internal/storage"
	"github.com/Shemistan/grpc_user_api/internal/utils"
)

type service struct {
	userStorage   storage.User
	hasher        utils.Hasher
	tokenProvider utils.TokenProvider
	config        *config.TokenServiceConfig
}

// NewService - новый сервис
func NewService(
	storage storage.User,
	hasher utils.Hasher,
	tokenProvider utils.TokenProvider,
	config *config.TokenServiceConfig,
) def.Auth {
	return &service{
		userStorage:   storage,
		hasher:        hasher,
		tokenProvider: tokenProvider,
		config:        config,
	}
}
