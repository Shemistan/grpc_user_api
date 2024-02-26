package api

import (
	"context"

	"github.com/brianvoe/gofakeit"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/Shemistan/grpc_user_api/pkg/user_api_v1"
)

// Get - получить пользователя по id
func (u *User) Get(_ context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	return &pb.GetResponse{
		Id:        req.GetId(),
		Name:      gofakeit.Name(),
		Email:     gofakeit.Email(),
		Role:      0,
		CreatedAt: &timestamppb.Timestamp{},
		UpdatedAt: &timestamppb.Timestamp{},
	}, nil
}
