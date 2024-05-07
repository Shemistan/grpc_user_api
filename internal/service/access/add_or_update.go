package auth

import (
	"context"

	"github.com/Shemistan/grpc_user_api/internal/model"
)

// AddOrUpdateAccess - добавить или редактировать доступ к ресурсу
func (s *service) AddOrUpdateAccess(ctx context.Context, req model.AccessRequest) error {
	err := s.accessStorage.UpsertAccess(ctx, req)
	if err != nil {
		return err
	}

	s.cache.AddInCache(req)

	return nil
}
