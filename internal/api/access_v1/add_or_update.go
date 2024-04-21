package user_v1

import (
	"context"

	"github.com/Shemistan/grpc_user_api/internal/converter"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/Shemistan/grpc_user_api/pkg/access_api_v1"
)

// AddOrUpdateAccess - добавить или изменить доступ
func (a *Access) AddOrUpdateAccess(ctx context.Context, req *pb.AddOrUpdateAccessRequest) (*emptypb.Empty, error) {
	err := a.service.AddOrUpdateAccess(ctx, converter.RPCAccessRequestToServiceAccessRequest(req))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
