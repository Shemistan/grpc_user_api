package tests

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/Shemistan/grpc_user_api/internal/api/user_v1"
	"github.com/Shemistan/grpc_user_api/internal/converter"
	"github.com/Shemistan/grpc_user_api/internal/model"
	"github.com/Shemistan/grpc_user_api/internal/service/mocks"
	pb "github.com/Shemistan/grpc_user_api/pkg/user_api_v1"
)

func TestGet(t *testing.T) {
	service := new(mocks.User)
	api := user_v1.New(service)
	ctx := context.Background()
	errorTest := errors.New("test error")

	reqAPI := &pb.GetRequest{
		Id: 10,
	}

	createdAt := time.Now()
	updatedAt := time.Now().Add(10 * time.Hour)

	user := model.User{
		ID:              10,
		Name:            "Name",
		Email:           "Email",
		Password:        "Password",
		PasswordConfirm: "Password",
		Role:            1,
		CreatedAt:       createdAt,
		UpdatedAt:       &updatedAt,
	}

	expect := converter.ModelUserToRPCGetUserResponse(user)

	t.Run("service error", func(t *testing.T) {
		service.On("GetUser", ctx, reqAPI.GetId()).
			Return(model.User{}, errorTest).Once()

		actual, err := api.Get(ctx, reqAPI)
		assert.ErrorIs(t, err, errorTest)
		assert.Nil(t, actual)
	})

	t.Run("service valid", func(t *testing.T) {
		service.On("GetUser", ctx, reqAPI.GetId()).
			Return(user, nil).Once()

		actual, err := api.Get(ctx, reqAPI)
		assert.NoError(t, err)
		assert.Equal(t, expect, actual)
	})
}
