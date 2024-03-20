package api

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Shemistan/grpc_user_api/internal/converter"
	pb "github.com/Shemistan/grpc_user_api/pkg/user_api_v1"
)

// Update - редактировать имя, почту и роль пользователя
func (u *User) Update(ctx context.Context, req *pb.UpdateRequest) (*emptypb.Empty, error) {
	err := u.Service.Update(ctx, converter.RPCUpdateUserToModelUpdateUser(req))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
