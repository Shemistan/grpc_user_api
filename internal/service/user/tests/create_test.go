package tests

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/Shemistan/grpc_user_api/internal/model"
	"github.com/Shemistan/grpc_user_api/internal/model/service_errors"
	"github.com/Shemistan/grpc_user_api/internal/service/user"
	storageMocks "github.com/Shemistan/grpc_user_api/internal/storage/mocks"
	utilsMocks "github.com/Shemistan/grpc_user_api/internal/utils/mocks"
)

func TestCreate(t *testing.T) {
	ctx := context.Background()
	storage := new(storageMocks.User)
	hasher := new(utilsMocks.Hasher)
	service := user.NewService(storage, hasher)
	errorTest := errors.New("test error")

	createdAt := time.Now()
	updatedAt := time.Now().Add(10 * time.Hour)

	hash := "hash"

	req := model.User{
		ID:              10,
		Name:            "Name",
		Email:           "Email",
		Password:        "Password",
		PasswordConfirm: "Password",
		Role:            1,
		CreateAt:        createdAt,
		UpdateAt:        &updatedAt,
	}

	t.Run("password mismatch", func(t *testing.T) {
		actual, err := service.Create(ctx, model.User{
			ID:              10,
			Name:            "Name",
			Email:           "Email",
			Password:        "Password",
			PasswordConfirm: "Password2",
			Role:            1,
			CreateAt:        createdAt,
			UpdateAt:        &updatedAt,
		})
		assert.ErrorIs(t, err, service_errors.PasswordMismatch)
		assert.Equal(t, int64(0), actual)
	})

	t.Run("hasher error", func(t *testing.T) {
		hasher.On("GetPasswordHash", req.Password).
			Return("", errorTest).Once()

		actual, err := service.Create(ctx, req)
		assert.ErrorIs(t, err, errorTest)
		assert.Equal(t, int64(0), actual)
	})

	t.Run("storage error", func(t *testing.T) {
		hasher.On("GetPasswordHash", req.Password).
			Return(hash, nil).Once()

		storage.On("Create", ctx, req, hash).
			Return(int64(0), errorTest).Once()

		actual, err := service.Create(ctx, req)

		assert.ErrorIs(t, err, errorTest)
		assert.Equal(t, int64(0), actual)
	})

	t.Run("success", func(t *testing.T) {
		hasher.On("GetPasswordHash", req.Password).
			Return(hash, nil).Once()

		storage.On("Create", ctx, req, hash).
			Return(int64(10), nil).Once()

		actual, err := service.Create(ctx, req)

		assert.NoError(t, err)
		assert.Equal(t, int64(10), actual)
	})
}
