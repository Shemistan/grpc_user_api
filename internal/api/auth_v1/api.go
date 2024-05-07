package user_v1

import (
	"github.com/Shemistan/grpc_user_api/internal/service"
	pb "github.com/Shemistan/grpc_user_api/pkg/auth_api_v1"
)

// Auth - структура реализующая методы АПИ
type Auth struct {
	pb.UnimplementedAuthV1Server

	service service.Auth
}

// New - создать новую структуру реализующая методы АПИ
func New(service service.Auth) *Auth {
	return &Auth{
		service: service,
	}
}
