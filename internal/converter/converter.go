package converter

import (
	"github.com/Shemistan/grpc_user_api/internal/model"
	pb "github.com/Shemistan/grpc_user_api/pkg/user_api_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// RPCCreateUserToModelUser конвертер из rpc в model
func RPCCreateUserToModelUser(req *pb.CreateRequest) model.User {
	if req == nil {
		return model.User{}
	}

	return model.User{
		Name:            req.GetName(),
		Email:           req.GetEmail(),
		Password:        req.GetPassword(),
		PasswordConfirm: req.GetPasswordConfirm(),
		Role:            int64(req.GetRole().Number()),
	}
}

// RPCUpdateUserToModelUpdateUser конвертер из rpc в model
func RPCUpdateUserToModelUpdateUser(req *pb.UpdateRequest) model.UpdateUser {
	var res model.UpdateUser

	if req == nil {
		return res
	}

	var name, email, oldPassword, newPassword, newPasswordConfirm string

	if req.Name != nil {
		name = req.Name.GetValue()
		res.Name = &name
	}

	if req.Email != nil {
		email = req.Email.GetValue()
		res.Email = &email
	}

	if req.OldPassword != nil {
		oldPassword = req.OldPassword.GetValue()
		res.OldPassword = &oldPassword
	}

	if req.NewPassword != nil {
		newPassword = req.NewPassword.GetValue()
		res.NewPassword = &newPassword
	}

	if req.NewPasswordConfirm != nil {
		newPasswordConfirm = req.NewPasswordConfirm.GetValue()
		res.NewPasswordConfirm = &newPasswordConfirm
	}

	res.ID = req.GetId()
	res.Role = int64(req.GetRole())

	return res
}

// ModelUserToRPCGetUserResponse - конвертер из rpc в model
func ModelUserToRPCGetUserResponse(req model.User) *pb.GetResponse {
	return &pb.GetResponse{
		Id:        req.ID,
		Name:      req.Name,
		Email:     req.Email,
		Role:      pb.Role(req.Role),
		CreatedAt: timestamppb.New(req.CreatedAt),
		UpdatedAt: timestamppb.New(*req.UpdatedAt),
	}
}
