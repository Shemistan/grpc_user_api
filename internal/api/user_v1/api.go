package user_v1

import (
	"github.com/Shemistan/grpc_user_api/internal/service"
	pb "github.com/Shemistan/grpc_user_api/pkg/user_api_v1"
)

// User - структура реализующая методы АПИ
type User struct {
	pb.UnimplementedUserV1Server

	service service.User
}

// New - создать новую структуру реализующая методы АПИ
func New(service service.User) *User {
	return &User{
		service: service,
	}
}
