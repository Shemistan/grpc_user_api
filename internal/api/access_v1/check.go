package user_v1

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/Shemistan/grpc_user_api/pkg/access_api_v1"
)

// Check - проверка доступа к ресурсу
func (a *Access) Check(ctx context.Context, req *pb.CheckRequest) (*emptypb.Empty, error) {
	err := a.service.Check(ctx, req.GetEndpointAddress())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
