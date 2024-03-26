package converter

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/Shemistan/grpc_user_api/internal/model"
	pb "github.com/Shemistan/grpc_user_api/pkg/user_api_v1"
)

func TestRPCCreateUserToModelUser(t *testing.T) {
	t.Run("req equal nil", func(t *testing.T) {
		expected := model.User{}
		actual := RPCCreateUserToModelUser(nil)
		assert.Equal(t, expected, actual)
	})

	t.Run("not name in req", func(t *testing.T) {
		expected := model.User{
			ID:              0,
			Email:           "Email",
			Password:        "Password",
			PasswordConfirm: "Password",
			Role:            0,
		}
		actual := RPCCreateUserToModelUser(&pb.CreateRequest{
			Email:           "Email",
			Password:        "Password",
			PasswordConfirm: "Password",
			Role:            0,
		})
		assert.Equal(t, expected, actual)
	})

	t.Run("req is valid", func(t *testing.T) {
		expected := model.User{
			Name:            "Name",
			Email:           "Email",
			Password:        "Password",
			PasswordConfirm: "PasswordConfirm",
			Role:            1,
		}
		actual := RPCCreateUserToModelUser(&pb.CreateRequest{
			Name:            "Name",
			Email:           "Email",
			Password:        "Password",
			PasswordConfirm: "PasswordConfirm",
			Role:            1,
		})
		assert.Equal(t, expected, actual)
	})
}

func TestRPCUpdateUserToModelUpdateUser(t *testing.T) {
	name := "name"
	email := "email"
	oldPassword := "oldPassword"
	newPassword := "newPassword"
	newPasswordConfirm := "newPasswordConfirm"
	role := int64(1)

	pbName := &wrapperspb.StringValue{Value: "name"}
	pbEmail := &wrapperspb.StringValue{Value: "email"}
	pbOldPassword := &wrapperspb.StringValue{Value: "oldPassword"}
	pbNewPassword := &wrapperspb.StringValue{Value: "newPassword"}
	pbNewPasswordConfirm := &wrapperspb.StringValue{Value: "newPasswordConfirm"}
	pbRole := pb.Role(1)

	id := int64(123)

	t.Run("req equal nil", func(t *testing.T) {
		expected := model.UpdateUser{}
		actual := RPCUpdateUserToModelUpdateUser(nil)
		assert.Equal(t, expected, actual)
	})

	t.Run("not name in req", func(t *testing.T) {
		expected := model.UpdateUser{
			ID:                 id,
			Email:              &email,
			OldPassword:        &oldPassword,
			NewPassword:        &newPassword,
			NewPasswordConfirm: &newPasswordConfirm,
			Role:               role,
		}
		actual := RPCUpdateUserToModelUpdateUser(&pb.UpdateRequest{
			Id:                 id,
			Email:              pbEmail,
			Role:               pbRole,
			OldPassword:        pbOldPassword,
			NewPassword:        pbNewPassword,
			NewPasswordConfirm: pbNewPasswordConfirm,
		})
		assert.Equal(t, expected, actual)
	})

	t.Run("not email in req", func(t *testing.T) {
		expected := model.UpdateUser{
			ID:                 id,
			Name:               &name,
			OldPassword:        &oldPassword,
			NewPassword:        &newPassword,
			NewPasswordConfirm: &newPasswordConfirm,
			Role:               role,
		}
		actual := RPCUpdateUserToModelUpdateUser(&pb.UpdateRequest{
			Id:                 id,
			Name:               pbName,
			Role:               pbRole,
			OldPassword:        pbOldPassword,
			NewPassword:        pbNewPassword,
			NewPasswordConfirm: pbNewPasswordConfirm,
		})
		assert.Equal(t, expected, actual)
	})

	t.Run("not oldPassword in req", func(t *testing.T) {
		expected := model.UpdateUser{
			ID:                 id,
			Name:               &name,
			Email:              &email,
			NewPassword:        &newPassword,
			NewPasswordConfirm: &newPasswordConfirm,
			Role:               role,
		}
		actual := RPCUpdateUserToModelUpdateUser(&pb.UpdateRequest{
			Id:                 id,
			Name:               pbName,
			Email:              pbEmail,
			Role:               pbRole,
			NewPassword:        pbNewPassword,
			NewPasswordConfirm: pbNewPasswordConfirm,
		})
		assert.Equal(t, expected, actual)
	})

	t.Run("not newPassword in req", func(t *testing.T) {
		expected := model.UpdateUser{
			ID:                 id,
			Name:               &name,
			Email:              &email,
			OldPassword:        &oldPassword,
			NewPasswordConfirm: &newPasswordConfirm,
			Role:               role,
		}
		actual := RPCUpdateUserToModelUpdateUser(&pb.UpdateRequest{
			Id:                 id,
			Name:               pbName,
			Email:              pbEmail,
			Role:               pbRole,
			OldPassword:        pbOldPassword,
			NewPasswordConfirm: pbNewPasswordConfirm,
		})
		assert.Equal(t, expected, actual)
	})

	t.Run("not newPasswordConfirm in req", func(t *testing.T) {
		expected := model.UpdateUser{
			ID:          id,
			Name:        &name,
			Email:       &email,
			OldPassword: &oldPassword,
			NewPassword: &newPassword,
			Role:        role,
		}
		actual := RPCUpdateUserToModelUpdateUser(&pb.UpdateRequest{
			Id:          id,
			Name:        pbName,
			Email:       pbEmail,
			Role:        pbRole,
			OldPassword: pbOldPassword,
			NewPassword: pbNewPassword,
		})
		assert.Equal(t, expected, actual)
	})

	t.Run("success request", func(t *testing.T) {
		expected := model.UpdateUser{
			ID:                 id,
			Name:               &name,
			Email:              &email,
			OldPassword:        &oldPassword,
			NewPassword:        &newPassword,
			NewPasswordConfirm: &newPasswordConfirm,
			Role:               role,
		}
		actual := RPCUpdateUserToModelUpdateUser(&pb.UpdateRequest{
			Id:                 id,
			Name:               pbName,
			Email:              pbEmail,
			Role:               pbRole,
			OldPassword:        pbOldPassword,
			NewPassword:        pbNewPassword,
			NewPasswordConfirm: pbNewPasswordConfirm,
		})
		assert.Equal(t, expected, actual)
	})
}

func TestModelUserToRPCGetUserResponse(t *testing.T) {
	createAT := time.Now()
	updateAt := time.Now().Add(10 * time.Hour)
	id := int64(123)
	role := int64(1)
	pbRole := pb.Role(1)

	t.Run("not name in req", func(t *testing.T) {
		expected := &pb.GetResponse{
			Id:        id,
			Email:     "Email",
			Role:      pbRole,
			CreatedAt: timestamppb.New(createAT),
			UpdatedAt: timestamppb.New(updateAt),
		}
		actual := ModelUserToRPCGetUserResponse(model.User{
			ID:       id,
			Email:    "Email",
			Role:     role,
			CreateAt: createAT,
			UpdateAt: &updateAt,
		})
		assert.Equal(t, expected, actual)
	})

	t.Run("req is valid", func(t *testing.T) {
		expected := &pb.GetResponse{
			Id:        id,
			Name:      "Name",
			Email:     "Email",
			Role:      pbRole,
			CreatedAt: timestamppb.New(createAT),
			UpdatedAt: timestamppb.New(updateAt),
		}
		actual := ModelUserToRPCGetUserResponse(model.User{
			ID:       id,
			Name:     "Name",
			Email:    "Email",
			Role:     role,
			CreateAt: createAT,
			UpdateAt: &updateAt,
		})
		assert.Equal(t, expected, actual)
	})
}
