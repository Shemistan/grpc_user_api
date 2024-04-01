package tests

//
//import (
//	"context"
//	"github.com/Shemistan/grpc_user_api/internal/service/user"
//	utilsMocks "github.com/Shemistan/grpc_user_api/internal/utils/mocks"
//	"testing"
//
//	"github.com/stretchr/testify/assert"
//
//	"github.com/Shemistan/grpc_user_api/internal/storage/mocks"
//)
//
//func TestGetUser(t *testing.T) {
//	ctx := context.Background()
//	storage := new(mocks.User)
//	hasher := new(utilsMocks.Hasher)
//	service := user.NewService(storage, hasher)
//
//	t.Run("storage error", func(t *testing.T) {
//		storage.On("AddMessage", ctx, model.Message{}).
//			Return(errorTest).Once()
//
//		err := service.AddMessage(ctx, model.Message{})
//		assert.ErrorIs(t, err, errorTest)
//	})
//
//	t.Run("storage valid", func(t *testing.T) {
//		storage.On("AddMessage", ctx, model.Message{}).
//			Return(nil).Once()
//
//		err := service.AddMessage(ctx, model.Message{})
//		assert.Nil(t, err)
//	})
//}
