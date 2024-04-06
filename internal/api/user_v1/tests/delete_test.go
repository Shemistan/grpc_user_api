package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Shemistan/grpc_user_api/internal/api/user_v1"
	"github.com/Shemistan/grpc_user_api/internal/service/mocks"
	pb "github.com/Shemistan/grpc_user_api/pkg/user_api_v1"
)

func TestDelete(t *testing.T) {
	service := new(mocks.User)
	api := user_v1.New(service)
	ctx := context.Background()
	errorTest := errors.New("test error")

	reqAPI := &pb.DeleteRequest{
		Id: 1,
	}

	t.Run("service error", func(t *testing.T) {
		service.On("Delete", ctx, reqAPI.GetId()).
			Return(errorTest).Once()

		actual, err := api.Delete(ctx, reqAPI)
		assert.ErrorIs(t, err, errorTest)
		assert.Nil(t, actual)
	})

	t.Run("service valid", func(t *testing.T) {
		service.On("Delete", ctx, reqAPI.GetId()).
			Return(nil).Once()

		actual, err := api.Delete(ctx, reqAPI)
		assert.NoError(t, err)
		assert.Equal(t, &emptypb.Empty{}, actual)
	})
}
