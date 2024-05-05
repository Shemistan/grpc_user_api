package storage

import (
	"context"

	"github.com/Shemistan/grpc_user_api/internal/model"
)

// User - для работы с пользователями
type User interface {
	Create(ctx context.Context, req model.User, passwordHash string) (int64, error)
	Update(ctx context.Context, req model.UpdateUser, passwordHash *string) error
	GetUser(ctx context.Context, req model.GetUserRequest) (model.User, error)
	Delete(ctx context.Context, id int64) error
	GetPasswordHash(ctx context.Context, id int64) (string, error)
}

// Access - для работы с доступами
type Access interface {
	AddAccess(ctx context.Context, req model.AccessRequest) error
	UpdateAccess(ctx context.Context, req model.AccessRequest) error
	GetAccess(ctx context.Context, req model.AccessRequest) (model.AccessRequest, error)
	GetAllAccess(ctx context.Context) ([]model.AccessRequest, error)
	UpsertAccess(ctx context.Context, req model.AccessRequest) error
}

// Cache - кэш для хранения доступов
type Cache interface {
	AddInCache(req model.AccessRequest)
	GetAccessesForRole(role int64) map[string]bool
}
