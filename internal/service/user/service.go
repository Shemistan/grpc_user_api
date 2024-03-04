package user

import (
	"context"
	"github.com/Shemistan/grpc_user_api/internal/utils"

	"github.com/Shemistan/grpc_user_api/internal/model"
	def "github.com/Shemistan/grpc_user_api/internal/service"
	"github.com/Shemistan/grpc_user_api/internal/storage"
)

type service struct {
	storage storage.User
	hasher  utils.Hasher
}

// NewService - новый сервис
func NewService(storage storage.User) def.User {
	return &service{storage: storage}
}

// Create - пользователя
func (s *service) Create(ctx context.Context, req model.User) (int64, error) {
	return s.storage.Create(ctx, req)
}

// Update - пользователя
func (s *service) Update(ctx context.Context, req model.User) error {
	return s.storage.Update(ctx, req)
}

// GetUser - пользователя
func (s *service) GetUser(ctx context.Context, id int64) (model.User, error) {
	return s.storage.GetUser(ctx, id)
}

// Delete - пользователя
func (s *service) Delete(ctx context.Context, id int64) error {
	return s.storage.Delete(ctx, id)
}
