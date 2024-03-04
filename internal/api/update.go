package api

import (
	"context"
	"github.com/Shemistan/grpc_user_api/internal/converter"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/Shemistan/grpc_user_api/pkg/user_api_v1"
)

// Update - редактировать имя, почту и роль пользователя
func (u *User) Update(ctx context.Context, req *pb.UpdateRequest) (*emptypb.Empty, error) {
	err := u.service.Update(ctx, converter.RPCUpdateUserToModelUser(req))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
