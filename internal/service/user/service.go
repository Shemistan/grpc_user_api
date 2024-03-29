package user

import (
	"context"
	"errors"
	"log"

	"github.com/Shemistan/grpc_user_api/internal/model"
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
		return 0, errors.New("password mismatch")
	}

	passwordHash, err := s.hasher.GetPasswordHash(req.Password)
	if err != nil {
		log.Println("failed to hashing password:", err.Error())
		return 0, err
	}

	req.Password = passwordHash

	return s.userStorage.Create(ctx, req)
}

// Update - - обновление пользователя
func (s *service) Update(ctx context.Context, req model.UpdateUser) error {
	if req.NewPassword != nil && req.NewPasswordConfirm != nil {
		switch {
		case req.NewPassword != req.NewPasswordConfirm:
			return errors.New("password mismatch")
		case req.OldPassword == nil:
			return errors.New("password mismatch")
		}

		ok, err := s.checkUserPassword(ctx, req.ID, *req.OldPassword)
		if err != nil {
			return err
		}

		if !ok {
			return errors.New("old password not valid")
		}

		hash, err := s.hasher.GetPasswordHash(*req.NewPassword)
		if err != nil {
			return err
		}

		req.NewPassword = &hash
	}

	return s.userStorage.Update(ctx, req)
}

// GetUser - полуучение пользователя
func (s *service) GetUser(ctx context.Context, id int64) (model.User, error) {
	return s.userStorage.GetUser(ctx, id)
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

	return s.hasher.CheckPassword(passwordHash, password), nil
}
