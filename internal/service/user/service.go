package user

import (
	"context"
	"log"

	serviceErrors "github.com/Shemistan/grpc_user_api/internal/constants/errors"
	"github.com/Shemistan/grpc_user_api/internal/model"
	def "github.com/Shemistan/grpc_user_api/internal/service"
	"github.com/Shemistan/grpc_user_api/internal/storage"
	"github.com/Shemistan/grpc_user_api/internal/storage/user/converter"
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
		return 0, serviceErrors.ErrorPasswordMismatch
	}

	passwordHash, err := s.hasher.GetPasswordHash(req.Password)
	if err != nil {
		log.Println("failed to hashing password:", err.Error())
		return 0, err
	}

	res, err := s.userStorage.Create(ctx, converter.ServiceUserToStorageUser(req, passwordHash))
	if err != nil {
		return 0, err
	}

	return res, nil
}

// Update - - обновление пользователя
func (s *service) Update(ctx context.Context, req model.UpdateUser) error {
	var passwordHash *string

	if req.NewPassword != nil && req.NewPasswordConfirm != nil {
		switch {
		case *req.NewPassword != *req.NewPasswordConfirm:
			return serviceErrors.ErrorPasswordMismatch
		case req.OldPassword == nil:
			return serviceErrors.ErrorPasswordMismatch
		}

		ok, err := s.checkUserPassword(ctx, req.ID, *req.OldPassword)
		if err != nil {
			return err
		}

		if !ok {
			return serviceErrors.ErrorsOldPasswordNotValid
		}

		hash, err := s.hasher.GetPasswordHash(*req.NewPassword)
		if err != nil {
			return err
		}

		passwordHash = &hash
	}

	err := s.userStorage.Update(ctx, converter.ServiceUpdateUserToStorageUpdateUser(req, passwordHash))
	if err != nil {
		return err
	}

	return nil
}

// GetUser - полуучение пользователя
func (s *service) GetUser(ctx context.Context, id int64) (model.User, error) {
	res, err := s.userStorage.GetUser(ctx, id)
	if err != nil {
		return model.User{}, err
	}

	return converter.StorageUserToServiceUser(res), nil
}

// Delete - удаление пользователя
func (s *service) Delete(ctx context.Context, id int64) error {
	return s.userStorage.Delete(ctx, id)
}

func (s *service) checkUserPassword(ctx context.Context, userID int64, password string) (bool, error) {
	passwordHash, err := s.userStorage.GetPasswordHash(ctx, userID)
	if err != nil {
		return false, err
	}

	res := s.hasher.CheckPassword(passwordHash, password)
	return res, nil
}
