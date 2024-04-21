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

// Cache - кэш для хранения доступов
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

	s.initCache()

	go func(s *service) {
		err := s.initAccessibleRoles(context.Background())
		if err != nil {
			log.Println(err)
		}
	}(s)

	return s
}

func (s *service) initCache() {
	if s.cache == nil {
		// Ролей предполагается не так много, по этому эффективнее сделать такой кэш, что бы не плодить много хэштаблиц
		cache := make(map[int64]map[string]bool)

		s.cache = &Cache{
			Mutex:           &sync.Mutex{},
			accessibleRoles: cache,
		}
	}
}
