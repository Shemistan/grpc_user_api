package user_v1

import (
	"context"

	pb "github.com/Shemistan/grpc_user_api/pkg/auth_api_v1"
)

// GetRefreshToken -  получть Refresh токен
func (u *Auth) GetRefreshToken(ctx context.Context, req *pb.GetRefreshTokenRequest) (*pb.GetRefreshTokenResponse, error) {
	res, err := u.service.GetRefreshToken(ctx, req.GetRefreshToken())
	if err != nil {
		return nil, err
	}

	return &pb.GetRefreshTokenResponse{RefreshToken: res}, nil
}
