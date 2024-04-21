package user_v1

import (
	"context"

	"github.com/Shemistan/grpc_user_api/internal/converter"
	pb "github.com/Shemistan/grpc_user_api/pkg/auth_api_v1"
)

// Login - авторизоваться
func (u *Auth) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	res, err := u.service.Login(ctx, converter.RPCLoginToServiceLogin(req))
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		RefreshToken: res.RefreshToken,
		AccessToken:  res.AccessToken,
	}, nil
}
