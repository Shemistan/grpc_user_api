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
	passwordHash = "test"
)

func TestServiceUserToStorageUser(t *testing.T) {
	createAt := time.Now()
	updateAt := time.Now().Add(1 * time.Hour)

	t.Run("updateAt is null", func(t *testing.T) {
		expected := storageModel.User{
			ID:       1,
			Name:     "Name",
			Email:    "Email",
			Password: passwordHash,
			Role:     2,
			CreatedAt: sql.NullTime{
				Time:  createAt,
				Valid: true,
			},
			UpdatedAt: sql.NullTime{},
		}
		req := serviceModel.User{
			ID:              1,
			Name:            "Name",
			Email:           "Email",
			Password:        "Password",
			PasswordConfirm: "PasswordConfirm",
			Role:            2,
			CreateAt:        createAt,
			UpdateAt:        nil,
		}
		actual := ServiceUserToStorageUser(req, passwordHash)

		assert.Equal(t, expected, actual)
	})

	t.Run("with updateAt", func(t *testing.T) {
		expected := storageModel.User{
			ID:       1,
			Name:     "Name",
			Email:    "Email",
			Password: passwordHash,
			Role:     2,
			CreatedAt: sql.NullTime{
				Time:  createAt,
				Valid: true,
			},
			UpdatedAt: sql.NullTime{
				Time:  updateAt,
				Valid: true,
			},
		}
		req := serviceModel.User{
			ID:              1,
			Name:            "Name",
			Email:           "Email",
			Password:        "Password",
			PasswordConfirm: "PasswordConfirm",
			Role:            2,
			CreateAt:        createAt,
			UpdateAt:        &updateAt,
		}
		actual := ServiceUserToStorageUser(req, passwordHash)

		assert.Equal(t, expected, actual)
	})
}

func TestStorageUserToServiceUser(t *testing.T) {
	createAt := time.Now()
	updateAt := time.Now().Add(1 * time.Hour)

	t.Run("updateAt is null", func(t *testing.T) {
		req := storageModel.User{
			ID:       1,
			Name:     "Name",
			Email:    "Email",
			Password: passwordHash,
			Role:     2,
			CreatedAt: sql.NullTime{
				Time:  createAt,
				Valid: true,
			},
			UpdatedAt: sql.NullTime{},
		}
		expected := serviceModel.User{
			ID:       1,
			Name:     "Name",
			Email:    "Email",
			Password: passwordHash,
			Role:     2,
			CreateAt: createAt,
			UpdateAt: nil,
		}
		actual := StorageUserToServiceUser(req)

		assert.Equal(t, expected, actual)
	})

	t.Run("with updateAt", func(t *testing.T) {
		req := storageModel.User{
			ID:       1,
			Name:     "Name",
			Email:    "Email",
			Password: passwordHash,
			Role:     2,
			CreatedAt: sql.NullTime{
				Time:  createAt,
				Valid: true,
			},
			UpdatedAt: sql.NullTime{
				Time:  updateAt,
				Valid: true,
			},
		}
		expected := serviceModel.User{
			ID:       1,
			Name:     "Name",
			Email:    "Email",
			Password: passwordHash,
			Role:     2,
			CreateAt: createAt,
			UpdateAt: &updateAt,
		}
		actual := StorageUserToServiceUser(req)

		assert.Equal(t, expected, actual)
	})

}

func TestServiceUpdateUserToStorageUpdateUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		name := "name"
		email := "email"

		expected := storageModel.UpdateUser{
			ID:       1,
			Name:     &name,
			Email:    &email,
			Password: &passwordHash,
			Role:     2,
		}
		req := serviceModel.UpdateUser{
			ID:    1,
			Name:  &name,
			Email: &email,
			Role:  2,
		}
		actual := ServiceUpdateUserToStorageUpdateUser(req, &passwordHash)

		assert.Equal(t, expected, actual)
	})

}
