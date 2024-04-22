package user

import (
	"github.com/Shemistan/platform_common/pkg/db"

	def "github.com/Shemistan/grpc_user_api/internal/storage"
)

type storage struct {
	db        db.Client
	txManager db.TxManager
}

const (
	tableResourceAccess = "resource_access"
)

// NewStorage - новый storage
func NewStorage(db db.Client, txManager db.TxManager) def.Access {
	return &storage{
		db:        db,
		txManager: txManager,
	}
}
