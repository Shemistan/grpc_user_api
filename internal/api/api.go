package api

import (
	pb "github.com/Shemistan/grpc_user_api/pkg/user_api_v1"
)

// User - структура реализующая методы АПИ
type User struct {
	Secret string
	pb.UnimplementedUserV1Server
}
