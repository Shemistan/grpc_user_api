package user

import (
	"context"
	"log"

	"github.com/Shemistan/grpc_user_api/internal/model"
	serviceErrors "github.com/Shemistan/grpc_user_api/internal/model/service_errors"
	def "github.com/Shemistan/grpc_user_api/internal/service"
	"github.com/Shemistan/grpc_user_api/internal/storage"
	"github.com/Shemistan/grpc_user_api/internal/utils"
)

type service struct {
	userStorage storage.User
	hasher      utils.Hasher
}

// NewService - новый сервис
func NewService(storage storage.User, hasher utils.Hasher) def.User {
	return &service{
		userStorage: storage,
		hasher:      hasher,
	}
}

// Create - создание пользователя
func (s *service) Create(ctx context.Context, req model.User) (int64, error) {
	if req.Password != req.PasswordConfirm {
		return 0, serviceErrors.ErrPasswordMismatch
	}

	passwordHash, err := s.hasher.GetPasswordHash(req.Password)
	if err != nil {
		log.Println("failed to hashing password:", err.Error())
		return 0, err
	}

	userID, err := s.userStorage.Create(ctx, req, passwordHash)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

// Update - обновление пользователя
func (s *service) Update(ctx context.Context, req model.UpdateUser) error {
	var passwordHash *string

	if req.NewPassword != nil && req.NewPasswordConfirm != nil {
		if *req.NewPassword != *req.NewPasswordConfirm {
			return serviceErrors.ErrPasswordMismatch
		}

		if req.OldPassword == nil {
			return serviceErrors.ErrOldPasswordNotFound
		}

		ok, err := s.checkUserPassword(ctx, req.ID, *req.OldPassword)
		if err != nil {
			return err
		}

		if !ok {
			return serviceErrors.ErrOldPasswordNotValid
		}

		hash, err := s.hasher.GetPasswordHash(*req.NewPassword)
		if err != nil {
			return err
		}

		passwordHash = &hash
	}

	err := s.userStorage.Update(ctx, req, passwordHash)
	if err != nil {
		return err
	}

	return nil
}

// GetUser - полуучение пользователя
func (s *service) GetUser(ctx context.Context, id int64) (model.User, error) {
	user, err := s.userStorage.GetUser(ctx, model.GetUserRequest{
		ID:    &id,
		Email: nil,
	})
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

// Delete - удаление пользователя
func (s *service) Delete(ctx context.Context, id int64) error {
	err := s.userStorage.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) checkUserPassword(ctx context.Context, userID int64, password string) (bool, error) {
	passwordHash, err := s.userStorage.GetPasswordHash(ctx, userID)
	if err != nil {
		return false, err
	}

	return s.hasher.CheckPassword(passwordHash, password), nil
}
