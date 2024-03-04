package api

import (
	"context"
	"github.com/Shemistan/grpc_user_api/internal/converter"
	pb "github.com/Shemistan/grpc_user_api/pkg/user_api_v1"
)

// Get - получить пользователя по id
func (u *User) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	res, err := u.service.GetUser(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return converter.ModelUserToRPCGetUserResponse(res), nil
}
