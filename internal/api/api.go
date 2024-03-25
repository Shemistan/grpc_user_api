package api

import (
	"github.com/Shemistan/grpc_user_api/internal/service"
	pb "github.com/Shemistan/grpc_user_api/pkg/user_api_v1"
)

// User - структура реализующая методы АПИ
type User struct {
	pb.UnimplementedUserV1Server

	Service service.User
}
