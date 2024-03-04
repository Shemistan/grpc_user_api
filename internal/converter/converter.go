package converter

import (
	"github.com/Shemistan/grpc_user_api/internal/model"
	pb "github.com/Shemistan/grpc_user_api/pkg/user_api_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToLoMap[T any, R any](conv func(T) R) func(T, int) R {
	return func(item T, _ int) R {
		return conv(item)
	}
}

// RPCCreateUserToModelUser конвертер из rpc в model
func RPCCreateUserToModelUser(req *pb.CreateRequest) model.User {
	if req == nil {
		return model.User{}
	}

	return model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Role:     int64(req.Role),
	}
}

// RPCUpdateUserToModelUser конвертер из rpc в model
func RPCUpdateUserToModelUser(req *pb.UpdateRequest) model.User {
	if req == nil {
		return model.User{}
	}

	req.Email.GetValue()
	return model.User{
		ID:    req.Id,
		Name:  req.Name.GetValue(),
		Email: req.Email.GetValue(),
		Role:  int64(req.Role),
	}
}

func ModelUserToRPCGetUserResponse(req model.User) *pb.GetResponse {
	return &pb.GetResponse{
		Id:        req.ID,
		Name:      req.Name,
		Email:     req.Email,
		Role:      pb.Role(req.Role),
		CreatedAt: timestamppb.New(req.CreateAt),
		UpdatedAt: timestamppb.New(req.UpdateDate),
	}
}
