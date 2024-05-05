package access_cache

import (
	"github.com/Shemistan/grpc_user_api/internal/model"
)

// AddInCache - добавить новое значение в кэш
func (s *storage) AddInCache(req model.AccessRequest) {
	if _, ok := s.cache[req.Role]; !ok {
		ResourcesMap := make(map[string]bool)

		s.Lock()
		s.cache[req.Role] = ResourcesMap
		s.cache[req.Role][req.Resource] = req.IsAccess
		s.Unlock()
		return
	}

	s.Lock()
	s.cache[req.Role][req.Resource] = req.IsAccess
	s.Unlock()
}

// GetAccessesForRole - получить словарь с доступами для роли
func (s *storage) GetAccessesForRole(role int64) map[string]bool {
	if v, ok := s.cache[role]; ok {
		return v
	}

	return nil
}
