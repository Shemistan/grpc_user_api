package service

import (
	"context"

	"github.com/Shemistan/grpc_user_api/internal/model"
)

// User - для работы с пользователями
type User interface {
	Create(ctx context.Context, req model.User) (int64, error)
	Update(ctx context.Context, req model.UpdateUser) error
	GetUser(ctx context.Context, id int64) (model.User, error)
	Delete(ctx context.Context, id int64) error
}

// Auth - сервис авторизации
type Auth interface {
	Login(ctx context.Context, req model.LoginRequest) (model.LoginResponse, error)
	GetRefreshToken(ctx context.Context, req string) (string, error)
	GetAccessToken(ctx context.Context, req string) (string, error)
}

// Access - сервис обработки доступов
type Access interface {
	Check(ctx context.Context, req string) error
	AddOrUpdateAccess(ctx context.Context, req model.AccessRequest) error
}
