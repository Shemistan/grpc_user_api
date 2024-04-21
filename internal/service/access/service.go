package auth

import (
	"context"
	"log"
	"sync"

	def "github.com/Shemistan/grpc_user_api/internal/service"
	"github.com/Shemistan/grpc_user_api/internal/storage"
	"github.com/Shemistan/grpc_user_api/internal/utils"
)

type service struct {
	tokenProvider        utils.TokenProvider
	accessStorage        storage.Access
	accessTokenSecretKey string

	cache *Cache
}

type Cache struct {
	*sync.Mutex
	accessibleRoles map[int64]map[string]bool
}

// NewService - новый сервис
func NewService(
	token utils.TokenProvider,
	accessStorage storage.Access,
	accessTokenSecretKey string,
) def.Access {

	s := &service{
		tokenProvider:        token,
		accessStorage:        accessStorage,
		accessTokenSecretKey: accessTokenSecretKey,
	}

	go func(s *service) {
		err := s.initAccessibleRoles(context.Background())
		if err != nil {
			log.Println(err)
		}
	}(s)

	return s
}
