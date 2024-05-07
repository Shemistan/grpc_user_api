package converter

import (
	"github.com/Shemistan/grpc_user_api/internal/model"
	pbAccess "github.com/Shemistan/grpc_user_api/pkg/access_api_v1"
	pbAuth "github.com/Shemistan/grpc_user_api/pkg/auth_api_v1"
	pbUser "github.com/Shemistan/grpc_user_api/pkg/user_api_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// RPCCreateUserToModelUser конвертер из rpc в model
func RPCCreateUserToModelUser(req *pbUser.CreateRequest) model.User {
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
func RPCUpdateUserToModelUpdateUser(req *pbUser.UpdateRequest) model.UpdateUser {
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
func ModelUserToRPCGetUserResponse(req model.User) *pbUser.GetResponse {
	var updateAt *timestamppb.Timestamp

	if req.UpdatedAt != nil {
		updateAt = timestamppb.New(*req.UpdatedAt)
	}

	res := &pbUser.GetResponse{
		Id:        req.ID,
		Name:      req.Name,
		Email:     req.Email,
		Role:      pbUser.Role(req.Role),
		CreatedAt: timestamppb.New(req.CreatedAt),
		UpdatedAt: updateAt,
	}

	if req.UpdatedAt != nil {
		res.UpdatedAt = timestamppb.New(*req.UpdatedAt)
	}

	return res
}

// RPCAccessRequestToServiceAccessRequest - конвертер из rpc в model
func RPCAccessRequestToServiceAccessRequest(req *pbAccess.AddOrUpdateAccessRequest) model.AccessRequest {
	if req == nil {
		return model.AccessRequest{}
	}

	return model.AccessRequest{
		Role:     req.GetRole(),
		Resource: req.GetResource(),
		IsAccess: req.GetIsAccess(),
	}
}

// RPCLoginToServiceLogin - конвертер из rpc в model
func RPCLoginToServiceLogin(req *pbAuth.LoginRequest) model.LoginRequest {
	if req == nil {
		return model.LoginRequest{}
	}

	return model.LoginRequest{
		Login:    req.GetLogin(),
		Password: req.GetPassword(),
	}
}
