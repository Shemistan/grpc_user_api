package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Shemistan/grpc_user_api/internal/api/user_v1"
	"github.com/Shemistan/grpc_user_api/internal/service/mocks"
)

func TestNew(t *testing.T) {
	service := new(mocks.User)

	api := user_v1.New(service)
	assert.NotNil(t, api)
}
