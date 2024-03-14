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
	storage storage.User
	hasher  utils.Hasher
}

// NewService - новый сервис
func NewService(storage storage.User, hasher utils.Hasher) def.User {
	return &service{
		storage: storage,
		hasher:  hasher,
	}
}

// Create - пользователя
func (s *service) Create(ctx context.Context, req model.User) (int64, error) {
	var err error

	req.Password, err = s.hasher.GetPasswordHash(req.Password)
	if err != nil {
		log.Println("failed to hashing password:", err.Error())
		return 0, err
	}

	return s.storage.Create(ctx, req)
}

// Update - пользователя
func (s *service) Update(ctx context.Context, req model.UpdateUser) error {
	if req.OldPassword != nil {
		ok, err := s.checkUserPassword(ctx, req.ID, *req.OldPassword)
		if err != nil {
			return err
		}

		if !ok {
			return errors.New("failed update: old password not valid")
		}

		hash, err := s.hasher.GetPasswordHash(*req.NewPassword)
		if err != nil {
			return err
		}

		req.NewPassword = &hash
	}

	return s.storage.Update(ctx, req)
}

// GetUser - пользователя
func (s *service) GetUser(ctx context.Context, id int64) (model.User, error) {
	return s.storage.GetUser(ctx, id)
}

// Delete - пользователя
func (s *service) Delete(ctx context.Context, id int64) error {
	return s.storage.Delete(ctx, id)
}

func (s *service) checkUserPassword(ctx context.Context, userID int64, password string) (bool, error) {
	passwordHash, err := s.storage.GetPasswordHash(ctx, userID)
	if err != nil {
		return false, err
	}

	return s.hasher.CheckPassword(passwordHash, password), nil
}
