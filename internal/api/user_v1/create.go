package user_v1

import (
	"context"

	"github.com/Shemistan/grpc_user_api/internal/converter"
	pb "github.com/Shemistan/grpc_user_api/pkg/user_api_v1"
)

// Create - создать пользователя
func (u *User) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	res, err := u.service.Create(ctx, converter.RPCCreateUserToModelUser(req))
	if err != nil {
		return nil, err
	}

	return &pb.CreateResponse{Id: res}, nil
}
