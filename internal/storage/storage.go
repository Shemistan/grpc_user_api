package storage

import (
	"context"

	"github.com/Shemistan/grpc_user_api/internal/model"
)

// User - для работы с пользователями
type User interface {
	Create(ctx context.Context, req model.User, passwordHash string) (int64, error)
	Update(ctx context.Context, req model.UpdateUser, passwordHash *string) error
	GetUser(ctx context.Context, id int64) (model.User, error)
	Delete(ctx context.Context, id int64) error
	GetPasswordHash(ctx context.Context, id int64) (string, error)
}
