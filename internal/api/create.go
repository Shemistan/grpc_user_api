package api

import (
	"context"
	"errors"

	"github.com/Shemistan/grpc_user_api/internal/converter"
	pb "github.com/Shemistan/grpc_user_api/pkg/user_api_v1"
)

// Create - создать пользователя
func (u *User) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	if req.GetPassword() != req.GetPasswordConfirm() {
		return nil, errors.New("password mismatch")
	}

	res, err := u.Service.Create(ctx, converter.RPCCreateUserToModelUser(req))
	if err != nil {
		return nil, err
	}

	return &pb.CreateResponse{Id: res}, nil
}
