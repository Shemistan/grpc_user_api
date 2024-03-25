package converter

import (
	"database/sql"
	"time"

	serviceModel "github.com/Shemistan/grpc_user_api/internal/model"
	storageModel "github.com/Shemistan/grpc_user_api/internal/storage/user/model"
)

// ServiceUserToStorageUser - конвертер из модели сервиса в модель хранилища
func ServiceUserToStorageUser(req serviceModel.User, passwordHash string) storageModel.User {
	var updateAt sql.NullTime
	if req.UpdateAt != nil {
		updateAt = sql.NullTime{
			Time:  *req.UpdateAt,
			Valid: true,
		}
	}

	return storageModel.User{
		ID:        req.ID,
		Name:      req.Name,
		Email:     req.Email,
		Password:  passwordHash,
		Role:      req.Role,
		UpdatedAt: updateAt,
	}
}

// StorageUserToServiceUser - конвертер из модели хранилища в модель сервиса
func StorageUserToServiceUser(req storageModel.User) serviceModel.User {
	var updateAt *time.Time
	if req.UpdatedAt.Valid {
		updateAt = &req.UpdatedAt.Time
	}

	return serviceModel.User{
		ID:       req.ID,
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
		CreateAt: req.CreatedAt.Time,
		UpdateAt: updateAt,
	}
}

// ServiceUpdateUserToStorageUpdateUser - конвертер из модели сервиса в модель хранилища
func ServiceUpdateUserToStorageUpdateUser(req serviceModel.UpdateUser, passwordHash *string) storageModel.UpdateUser {
	return storageModel.UpdateUser{
		ID:       req.ID,
		Name:     req.Name,
		Email:    req.Email,
		Password: passwordHash,
		Role:     req.Role,
	}
}
