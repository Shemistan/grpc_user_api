package user_v1

import (
	"github.com/Shemistan/grpc_user_api/internal/service"
	pb "github.com/Shemistan/grpc_user_api/pkg/access_api_v1"
)

// Access - структура реализующая методы АПИ
type Access struct {
	pb.UnimplementedAccessV1Server

	service service.Access
}

// New - создать новую структуру реализующая методы АПИ
func New(service service.Access) *Access {
	return &Access{
		service: service,
	}
}
