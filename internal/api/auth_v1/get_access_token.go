package user_v1

import (
	"context"

	pb "github.com/Shemistan/grpc_user_api/pkg/auth_api_v1"
)

// GetAccessToken - получть Access токен
func (u *Auth) GetAccessToken(ctx context.Context, req *pb.GetAccessTokenRequest) (*pb.GetAccessTokenResponse, error) {
	res, err := u.service.GetAccessToken(ctx, req.GetAccessToken())
	if err != nil {
		return nil, err
	}

	return &pb.GetAccessTokenResponse{AccessToken: res}, nil
}
