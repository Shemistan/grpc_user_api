package api

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/Shemistan/grpc_user_api/pkg/user_api_v1"
)

// Update - редактировать имя, почту и роль пользователя
func (u *User) Update(_ context.Context, req *pb.UpdateRequest) (*emptypb.Empty, error) {
	log.Printf("request: %+v", req)
	return &emptypb.Empty{}, nil
}
