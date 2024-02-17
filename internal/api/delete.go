package api

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/Shemistan/grpc_user_api/pkg/user_api_v1"
)

// Delete - удалить пользователя по id
func (u *User) Delete(_ context.Context, req *pb.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("request: %+v", req)
	return &emptypb.Empty{}, nil
}
