package auth

import (
	"context"

	"github.com/Shemistan/grpc_user_api/internal/model"
)

// В идеале нужна очистка кэша от переполнения, но по идее при обновлении он чистится и актуализируется
func (s *service) addInCache(req model.AccessRequest) {
	if _, ok := s.cache.accessibleRoles[req.Role]; !ok {
		urlsMap := make(map[string]bool)

		s.cache.Lock()
		s.cache.accessibleRoles[req.Role] = urlsMap
		s.cache.accessibleRoles[req.Role][req.URL] = req.IsAccess
		s.cache.Unlock()
		return
	}

	s.cache.Lock()
	s.cache.accessibleRoles[req.Role][req.URL] = req.IsAccess
	s.cache.Unlock()
}

func (s *service) initAccessibleRoles(ctx context.Context) error {
	roles, err := s.accessStorage.GetAllAccess(ctx)
	if err != nil {
		return err
	}

	for _, v := range roles {
		s.addInCache(v)
	}

	return nil
}
