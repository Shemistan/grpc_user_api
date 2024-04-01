package tests

//import (
//	"context"
//	"errors"
//	"testing"
//
//	"github.com/stretchr/testify/assert"
//
//	"github.com/Shemistan/chat_server/internal/service/chat"
//	"github.com/Shemistan/chat_server/internal/service/mocks"
//)
//
//func TestDeactivateChat(t *testing.T) {
//	ctx := context.Background()
//	storage := new(mocks.Chat)
//	service := chat.NewService(storage)
//	errorTest := errors.New("test error")
//
//	t.Run("storage error", func(t *testing.T) {
//		storage.On("DeactivateChat", ctx, int64(1)).
//			Return(errorTest).Once()
//
//		err := service.DeactivateChat(ctx, int64(1))
//		assert.ErrorIs(t, err, errorTest)
//	})
//
//	t.Run("storage valid", func(t *testing.T) {
//		storage.On("DeactivateChat", ctx, int64(1)).
//			Return(nil).Once()
//
//		err := service.DeactivateChat(ctx, int64(1))
//		assert.Nil(t, err)
//	})
//}
