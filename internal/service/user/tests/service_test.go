package tests

import (
	utilsMocks "github.com/Shemistan/grpc_user_api/internal/utils/mocks"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Shemistan/grpc_user_api/internal/service/user"
	storageMocks "github.com/Shemistan/grpc_user_api/internal/storage/mocks"
)

func TestNewService(t *testing.T) {
	storage := new(storageMocks.User)
	hasher := new(utilsMocks.Hasher)
	service := user.NewService(storage, hasher)

	assert.NotNil(t, service)
}
