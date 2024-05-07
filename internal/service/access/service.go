package auth

import (
	def "github.com/Shemistan/grpc_user_api/internal/service"
	"github.com/Shemistan/grpc_user_api/internal/storage"
	"github.com/Shemistan/grpc_user_api/internal/utils"
)

type service struct {
	tokenProvider        utils.TokenProvider
	accessStorage        storage.Access
	accessTokenSecretKey string
	cache                storage.Cache
}

// NewService - новый сервис
func NewService(
	token utils.TokenProvider,
	accessStorage storage.Access,
	accessTokenSecretKey string,
	cache storage.Cache,
) def.Access {

	return &service{
		tokenProvider:        token,
		accessStorage:        accessStorage,
		accessTokenSecretKey: accessTokenSecretKey,
		cache:                cache,
	}
}
