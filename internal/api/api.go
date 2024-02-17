package api

import (
	pb "github.com/Shemistan/grpc_user_api/pkg/user_api_v1"
)

// User - структура реализующая методы АПИ
type User struct {
	pb.UnimplementedUserV1Server
}
