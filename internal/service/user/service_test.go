package user

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/Shemistan/grpc_user_api/internal/model"
	serviceErrors "github.com/Shemistan/grpc_user_api/internal/model/service_errors"
	"github.com/Shemistan/grpc_user_api/internal/storage/mocks"
	storageModel "github.com/Shemistan/grpc_user_api/internal/storage/user/model"
	mocksHasher "github.com/Shemistan/grpc_user_api/internal/utils/mocks"
)

var (
	password     = "password"
	passwordHash = "hash"
)

func TestNewService(t *testing.T) {
	storage := new(mocks.User)
	hasher := new(mocksHasher.Hasher)
	s := NewService(storage, hasher)

	assert.NotNil(t, s)
}

func TestCreate(t *testing.T) {
	storage := new(mocks.User)
	hasher := new(mocksHasher.Hasher)
	s := NewService(storage, hasher)

	ctx := context.Background()
	crateAt := time.Now()

	user := model.User{
		Name:            "Name",
		Email:           "Email",
		Password:        "Password",
		PasswordConfirm: "Password",
		Role:            1,
		CreatedAt:       crateAt,
	}

	testError := errors.New("test error")

	t.Run("passwords not equal", func(t *testing.T) {
		_, err := s.Create(ctx, model.User{
			Name:            "Name",
			Email:           "Email",
			Password:        "Password",
			PasswordConfirm: "PasswordConfirm",
			Role:            0,
			CreatedAt:       crateAt,
		})

		assert.ErrorIs(t, err, serviceErrors.ErrPasswordMismatch)
	})

	t.Run("password hash error", func(t *testing.T) {
		hasher.On("GetPasswordHash", user.Password).
			Return("", testError).Once()

		_, err := s.Create(ctx, user)

		assert.ErrorIs(t, err, testError)
	})

	t.Run("storage error", func(t *testing.T) {
		hasher.On("GetPasswordHash", user.Password).
			Return(passwordHash, nil).Once()

		storage.On("Create", ctx, user, passwordHash).
			Return(int64(0), testError).Once()

		_, err := s.Create(ctx, user)

		assert.ErrorIs(t, err, testError)
	})

	t.Run("success", func(t *testing.T) {
		hasher.On("GetPasswordHash", user.Password).
			Return(passwordHash, nil).Once()

		storage.On("Create", ctx, user, passwordHash).
			Return(int64(1), nil).Once()

		expect := int64(1)
		actual, err := s.Create(ctx, user)

		assert.NoError(t, err)
		assert.Equal(t, expect, actual)
	})
}

func TestUpdate(t *testing.T) {
	storage := new(mocks.User)
	hasher := new(mocksHasher.Hasher)
	s := NewService(storage, hasher)

	oldPassword := "password"
	name := "name"
	email := "email"
	userID := int64(1)

	ctx := context.Background()
	passHash := "hash"
	user := model.UpdateUser{
		ID:                 userID,
		Name:               &name,
		Email:              &email,
		OldPassword:        &oldPassword,
		NewPassword:        &password,
		NewPasswordConfirm: &password,
		Role:               1,
	}

	testError := errors.New("test error")

	t.Run("new passwords not equal", func(t *testing.T) {
		err := s.Update(ctx, model.UpdateUser{
			ID:                 userID,
			Name:               &name,
			Email:              &email,
			OldPassword:        &password,
			NewPassword:        &password,
			NewPasswordConfirm: &passHash,
			Role:               1,
		})

		assert.ErrorIs(t, err, serviceErrors.ErrPasswordMismatch)
	})

	t.Run("old password is nil", func(t *testing.T) {
		err := s.Update(ctx, model.UpdateUser{
			ID:                 userID,
			Name:               &name,
			Email:              &email,
			OldPassword:        nil,
			NewPassword:        &password,
			NewPasswordConfirm: &password,
			Role:               1,
		})

		assert.ErrorIs(t, err, serviceErrors.ErrOldPasswordNotFound)
	})

	t.Run("get password hash error", func(t *testing.T) {
		storage.On("GetPasswordHash", ctx, userID).
			Return("", testError).Once()

		err := s.Update(ctx, user)

		assert.ErrorIs(t, err, testError)
	})

	t.Run("password not valid", func(t *testing.T) {
		storage.On("GetPasswordHash", ctx, userID).
			Return(passHash, nil).Once()

		hasher.On("CheckPassword", passHash, password).
			Return(false).Once()

		err := s.Update(ctx, user)

		assert.ErrorIs(t, err, serviceErrors.ErrOldPasswordNotValid)
	})

	t.Run("get password hash error", func(t *testing.T) {
		storage.On("GetPasswordHash", ctx, userID).
			Return(passHash, nil).Once()

		hasher.On("CheckPassword", passHash, password).
			Return(true).Once()

		hasher.On("GetPasswordHash", password).
			Return("", testError).Once()

		err := s.Update(ctx, user)

		assert.ErrorIs(t, err, testError)
	})

	t.Run("with password update storage error", func(t *testing.T) {
		storage.On("GetPasswordHash", ctx, userID).
			Return(passHash, nil).Once()

		hasher.On("CheckPassword", passHash, password).
			Return(true).Once()

		hasher.On("GetPasswordHash", password).
			Return(passHash, nil).Once()

		storage.On("Update", ctx, user, &passHash).
			Return(testError).Once()

		err := s.Update(ctx, user)

		assert.ErrorIs(t, err, testError)
	})

	t.Run("with password update storage success", func(t *testing.T) {
		storage.On("GetPasswordHash", ctx, userID).
			Return(passHash, nil).Once()

		hasher.On("CheckPassword", passHash, password).
			Return(true).Once()

		hasher.On("GetPasswordHash", password).
			Return(passHash, nil).Once()

		storage.On("Update", ctx, user, &passHash).
			Return(nil).Once()

		err := s.Update(ctx, user)

		assert.NoError(t, err)
	})

	t.Run("without password update storage success", func(t *testing.T) {
		user2 := user
		user2.NewPassword = nil
		storage.On("Update", ctx, model.UpdateUser{
			ID:                 0,
			Name:               nil,
			Email:              nil,
			OldPassword:        nil,
			NewPassword:        nil,
			NewPasswordConfirm: nil,
			Role:               0,
		}, nil).
			Return(nil).Once()

		err := s.Update(ctx, user)

		assert.NoError(t, err)
	})

}

func TestGetUser(t *testing.T) {
	storage := new(mocks.User)
	hasher := new(mocksHasher.Hasher)
	s := NewService(storage, hasher)

	userID := int64(1)
	ctx := context.Background()
	testError := errors.New("test error")

	t.Run("storage error", func(t *testing.T) {
		storage.On("GetUser", ctx, userID).
			Return(storageModel.User{}, testError).Once()

		_, err := s.GetUser(ctx, userID)

		assert.ErrorIs(t, err, testError)
	})

	t.Run("success", func(t *testing.T) {
		crateAt := time.Now()
		updateAt := time.Now().Add(1 * time.Hour)

		storage.On("GetUser", ctx, userID).
			Return(storageModel.User{
				ID:       userID,
				Name:     "Name",
				Email:    "Email",
				Password: "Password",
				Role:     12,
				CreatedAt: sql.NullTime{
					Time:  crateAt,
					Valid: true,
				},
				UpdatedAt: sql.NullTime{
					Time:  updateAt,
					Valid: true,
				},
			}, nil).Once()

		expect := model.User{
			ID:              userID,
			Name:            "Name",
			Email:           "Email",
			Password:        "Password",
			PasswordConfirm: "",
			Role:            12,
			CreatedAt:       crateAt,
			UpdatedAt:       &updateAt,
		}

		actual, err := s.GetUser(ctx, userID)

		assert.NoError(t, err)
		assert.Equal(t, expect.ID, actual.ID)
		assert.Equal(t, expect.Name, actual.Name)
		assert.Equal(t, expect.Email, actual.Email)
		assert.Equal(t, expect.Password, actual.Password)
		assert.Equal(t, "", actual.PasswordConfirm)
		assert.Equal(t, expect.CreatedAt, actual.CreatedAt)
		assert.Equal(t, expect.UpdatedAt, actual.UpdatedAt)
	})
}

func TestDelete(t *testing.T) {
	storage := new(mocks.User)
	hasher := new(mocksHasher.Hasher)
	s := NewService(storage, hasher)

	userID := int64(1)
	ctx := context.Background()
	testError := errors.New("test error")

	t.Run("storage error", func(t *testing.T) {
		storage.On("Delete", ctx, userID).
			Return(testError).Once()

		err := s.Delete(ctx, userID)

		assert.ErrorIs(t, err, testError)
	})

	t.Run("success", func(t *testing.T) {
		storage.On("Delete", ctx, userID).
			Return(nil).Once()

		err := s.Delete(ctx, userID)

		assert.NoError(t, err)
	})
}
