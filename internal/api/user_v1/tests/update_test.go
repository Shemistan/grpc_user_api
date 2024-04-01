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
	api := user_v1.New(service)
	ctx := context.Background()
	errorTest := errors.New("test error")

	reqAPI := &pb.UpdateRequest{
		Id:                 1,
		Name:               &wrapperspb.StringValue{Value: "name"},
		Email:              &wrapperspb.StringValue{Value: "email"},
		Role:               1,
		OldPassword:        nil,
		NewPassword:        nil,
		NewPasswordConfirm: nil,
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
		assert.NoError(t, err)
		assert.Equal(t, &emptypb.Empty{}, actual)
	})
}
