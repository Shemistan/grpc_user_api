package converter

import (
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	serviceModel "github.com/Shemistan/grpc_user_api/internal/model"
	storageModel "github.com/Shemistan/grpc_user_api/internal/storage/user/model"
)

var (
	hash = "hash"
)

func TestServiceUserToStorageUser(t *testing.T) {
	//hash := "hash"
	createdAt := time.Now()
	updatedAt := time.Now().Add(10 * time.Hour)

	t.Run("updatedAt equal nil", func(t *testing.T) {
		expected := storageModel.User{
			ID:       1,
			Name:     "name",
			Email:    "email",
			Password: hash,
			Role:     1,
			CreatedAt: sql.NullTime{
				Time:  createdAt,
				Valid: true,
			},
			UpdatedAt: sql.NullTime{},
		}

		actual := ServiceUserToStorageUser(serviceModel.User{
			ID:              1,
			Name:            "name",
			Email:           "email",
			Password:        "Password",
			PasswordConfirm: "PasswordConfirm",
			Role:            1,
			CreatedAt:       createdAt,
			UpdatedAt:       nil,
		}, hash)
		assert.Equal(t, expected, actual)
	})

	t.Run("updatedAt not nil", func(t *testing.T) {
		expected := storageModel.User{
			ID:       1,
			Name:     "name",
			Email:    "email",
			Password: hash,
			Role:     1,
			CreatedAt: sql.NullTime{
				Time:  createdAt,
				Valid: true,
			},
			UpdatedAt: sql.NullTime{
				Time:  updatedAt,
				Valid: true,
			},
		}
		actual := ServiceUserToStorageUser(serviceModel.User{
			ID:              1,
			Name:            "name",
			Email:           "email",
			Password:        "Password",
			PasswordConfirm: "PasswordConfirm",
			Role:            1,
			CreatedAt:       createdAt,
			UpdatedAt:       &updatedAt,
		}, hash)
		assert.Equal(t, expected, actual)
	})

}

func TestStorageUserToServiceUser(t *testing.T) {
	//hash := "hash"
	createdAt := time.Now()
	updatedAt := time.Now().Add(10 * time.Hour)

	t.Run("updatedAt is false", func(t *testing.T) {
		expected := serviceModel.User{
			ID:        1,
			Name:      "name",
			Email:     "email",
			Password:  hash,
			Role:      1,
			CreatedAt: createdAt,
			UpdatedAt: nil,
		}

		actual := StorageUserToServiceUser(storageModel.User{
			ID:       1,
			Name:     "name",
			Email:    "email",
			Password: hash,
			Role:     1,
			CreatedAt: sql.NullTime{
				Time:  createdAt,
				Valid: true,
			},
			UpdatedAt: sql.NullTime{
				Time:  updatedAt,
				Valid: false,
			},
		})

		assert.Equal(t, expected, actual)

	})

	t.Run("with updatedAt", func(t *testing.T) {
		expected := serviceModel.User{
			ID:        1,
			Name:      "name",
			Email:     "email",
			Password:  hash,
			Role:      1,
			CreatedAt: createdAt,
			UpdatedAt: &updatedAt,
		}

		actual := StorageUserToServiceUser(storageModel.User{
			ID:       1,
			Name:     "name",
			Email:    "email",
			Password: hash,
			Role:     1,
			CreatedAt: sql.NullTime{
				Time:  createdAt,
				Valid: true,
			},
			UpdatedAt: sql.NullTime{
				Time:  updatedAt,
				Valid: true,
			},
		})

		assert.Equal(t, expected, actual)
	})
}

func TestServiceUpdateUserToStorageUpdateUser(t *testing.T) {
	hash := "hash"
	email := "email"
	name := "name"
	password := "password"

	t.Run("with passwordHash", func(t *testing.T) {
		expected := storageModel.UpdateUser{
			ID:       10,
			Name:     &name,
			Email:    &email,
			Password: &hash,
			Role:     1,
		}

		actual := ServiceUpdateUserToStorageUpdateUser(serviceModel.UpdateUser{
			ID:                 10,
			Name:               &name,
			Email:              &email,
			OldPassword:        &password,
			NewPassword:        &password,
			NewPasswordConfirm: &password,
			Role:               1,
		}, &hash)

		assert.Equal(t, expected, actual)

	})

	t.Run("without passwordHash", func(t *testing.T) {
		expected := storageModel.UpdateUser{
			ID:       10,
			Name:     &name,
			Email:    &email,
			Password: nil,
			Role:     1,
		}

		actual := ServiceUpdateUserToStorageUpdateUser(serviceModel.UpdateUser{
			ID:                 10,
			Name:               &name,
			Email:              &email,
			OldPassword:        &password,
			NewPassword:        &password,
			NewPasswordConfirm: &password,
			Role:               1,
		}, nil)

		assert.Equal(t, expected, actual)
	})
}
