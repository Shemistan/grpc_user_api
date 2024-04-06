package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Shemistan/grpc_user_api/internal/api/user_v1"
	"github.com/Shemistan/grpc_user_api/internal/converter"
	"github.com/Shemistan/grpc_user_api/internal/service/mocks"
	pb "github.com/Shemistan/grpc_user_api/pkg/user_api_v1"
)

func TestCreate(t *testing.T) {
	service := new(mocks.User)
	api := user_v1.New(service)
	ctx := context.Background()
	errorTest := errors.New("test error")

	reqAPI := &pb.CreateRequest{
		Name:            "Name",
		Email:           "Email",
		Password:        "Password",
		PasswordConfirm: "Password",
		Role:            1,
	}

	t.Run("service error", func(t *testing.T) {
		service.On("Create", ctx, converter.RPCCreateUserToModelUser(reqAPI)).
			Return(int64(0), errorTest).Once()

		actual, err := api.Create(ctx, reqAPI)
		assert.ErrorIs(t, err, errorTest)
		assert.Nil(t, actual)
	})

	t.Run("service valid", func(t *testing.T) {
		expect := int64(1)
		service.On("Create", ctx, converter.RPCCreateUserToModelUser(reqAPI)).
			Return(expect, nil).Once()

		actual, err := api.Create(ctx, reqAPI)
		assert.NoError(t, err)
		assert.Equal(t, expect, actual.GetId())
	})
}
