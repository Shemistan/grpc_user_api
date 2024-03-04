package api

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/Shemistan/grpc_user_api/pkg/user_api_v1"
)

// Delete - удалить пользователя по id
func (u *User) Delete(ctx context.Context, req *pb.DeleteRequest) (*emptypb.Empty, error) {
	err := u.service.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
