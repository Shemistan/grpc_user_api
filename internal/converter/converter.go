package converter

import (
	"time"

	"github.com/Shemistan/grpc_user_api/internal/model"
	pb "github.com/Shemistan/grpc_user_api/pkg/user_api_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

//func ToLoMap[T any, R any](conv func(T) R) func(T, int) R {
//	return func(item T, _ int) R {
//		return conv(item)
//	}
//}

// RPCCreateUserToModelUser конвертер из rpc в model
func RPCCreateUserToModelUser(req *pb.CreateRequest) model.User {
	if req == nil {
		return model.User{}
	}

	return model.User{
		Name:     req.GetName(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		Role:     int64(req.GetRole().Number()),
	}
}

// RPCUpdateUserToModelUpdateUser конвертер из rpc в model
func RPCUpdateUserToModelUpdateUser(req *pb.UpdateRequest) model.UpdateUser {
	if req == nil {
		return model.UpdateUser{}
	}

	var name, email, oldPassword, newPassword *string

	if v := req.Name.GetValue(); v != "" {
		name = &v
	}

	if v := req.Email.GetValue(); v != "" {
		email = &v
	}

	if v := req.OldPassword.GetValue(); v != "" {
		oldPassword = &v
	}

	if v := req.NewPassword.GetValue(); v != "" {
		newPassword = &v
	}

	return model.UpdateUser{
		ID:          req.GetId(),
		Name:        name,
		Email:       email,
		OldPassword: oldPassword,
		NewPassword: newPassword,
		Role:        int64(req.Role),
	}
}

// ModelUserToRPCGetUserResponse - конвертер из rpc в model
func ModelUserToRPCGetUserResponse(req model.User) *pb.GetResponse {
	var t time.Time
	if req.UpdateAt != nil {
		t = *req.UpdateAt
	}
	return &pb.GetResponse{
		Id:        req.ID,
		Name:      req.Name,
		Email:     req.Email,
		Role:      pb.Role(req.Role),
		CreatedAt: timestamppb.New(req.CreateAt),
		UpdatedAt: timestamppb.New(t),
	}
}
