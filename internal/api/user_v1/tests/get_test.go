package tests

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/Shemistan/grpc_user_api/internal/api/user_v1"
	"github.com/Shemistan/grpc_user_api/internal/model"
	"github.com/Shemistan/grpc_user_api/internal/service/mocks"
	pb "github.com/Shemistan/grpc_user_api/pkg/user_api_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestGet(t *testing.T) {
	service := new(mocks.User)
	ctx := context.Background()
	api := user_v1.New(service)
	assert.NotNil(t, api)

	errorTest := errors.New("test error")
	userID := int64(1)

	reqAPI := &pb.GetRequest{
		Id: userID,
	}

	createAt := time.Now()
	updateAt := time.Now().Add(2 * time.Hour)
	expected := pb.GetResponse{
		Id:        userID,
		Name:      "Name",
		Email:     "Email",
		Role:      1,
		CreatedAt: timestamppb.New(createAt),
		UpdatedAt: timestamppb.New(updateAt),
	}

	t.Run("service error", func(t *testing.T) {
		service.On("GetUser", ctx, userID).
			Return(model.User{}, errorTest).Once()

		actual, err := api.Get(ctx, reqAPI)
		assert.ErrorIs(t, err, errorTest)
		assert.Nil(t, actual)
	})

	t.Run("success", func(t *testing.T) {
		service.On("GetUser", ctx, userID).
			Return(model.User{
				ID:       userID,
				Name:     "Name",
				Email:    "Email",
				Role:     1,
				CreateAt: createAt,
				UpdateAt: &updateAt,
			}, nil).Once()

		actual, err := api.Get(ctx, reqAPI)

		assert.Nil(t, err)
		assert.Equal(t, expected.Id, actual.Id)
		assert.Equal(t, expected.Name, actual.Name)
		assert.Equal(t, expected.Email, actual.Email)
		assert.Equal(t, expected.Role, actual.Role)
		assert.Equal(t, expected.CreatedAt, actual.CreatedAt)
		assert.Equal(t, expected.UpdatedAt, actual.UpdatedAt)

	})
}
