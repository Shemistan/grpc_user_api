package access_cache

import (
	"sync"

	def "github.com/Shemistan/grpc_user_api/internal/storage"
)

type storage struct {
	*sync.Mutex
	cache map[int64]map[string]bool
}

// NewCache - новый кэш
func NewCache() def.Cache {
	cache := make(map[int64]map[string]bool)

	return &storage{
		Mutex: &sync.Mutex{},
		cache: cache,
	}
}
