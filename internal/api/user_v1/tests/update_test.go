package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/Shemistan/grpc_user_api/internal/api/user_v1"
	"github.com/Shemistan/grpc_user_api/internal/converter"
	"github.com/Shemistan/grpc_user_api/internal/service/mocks"
	pb "github.com/Shemistan/grpc_user_api/pkg/user_api_v1"
)

func TestUpdate(t *testing.T) {
	service := new(mocks.User)
	ctx := context.Background()
	api := user_v1.New(service)
	assert.NotNil(t, api)

	userID := int64(1)
	errorTest := errors.New("test error")
	expected := &emptypb.Empty{}

	reqAPI := &pb.UpdateRequest{
		Id:                 userID,
		Name:               &wrapperspb.StringValue{Value: "name"},
		Email:              &wrapperspb.StringValue{Value: "email"},
		Role:               1,
		OldPassword:        &wrapperspb.StringValue{Value: "old_password"},
		NewPassword:        &wrapperspb.StringValue{Value: "password"},
		NewPasswordConfirm: &wrapperspb.StringValue{Value: "password"},
	}

	t.Run("service error", func(t *testing.T) {
		service.On("Update", ctx, converter.RPCUpdateUserToModelUpdateUser(reqAPI)).
			Return(errorTest).Once()

		actual, err := api.Update(ctx, reqAPI)
		assert.ErrorIs(t, err, errorTest)
		assert.Nil(t, actual)
	})

	t.Run("service valid", func(t *testing.T) {
		service.On("Update", ctx, converter.RPCUpdateUserToModelUpdateUser(reqAPI)).
			Return(nil).Once()

		actual, err := api.Update(ctx, reqAPI)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})
}
