package api

import (
	"context"
	"log"

	pb "github.com/Shemistan/grpc_user_api/pkg/user_api_v1"
)

// Create - создать пользователя
func (u *User) Create(_ context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	log.Printf("request: %+v", req)
	return &pb.CreateResponse{Id: 1}, nil
}
