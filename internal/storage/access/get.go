package user

import (
	"context"
	"fmt"

	"github.com/Shemistan/platform_common/pkg/db"

	"github.com/Shemistan/grpc_user_api/internal/model"
)

// GetAccess - получить информацию о доступе
func (s *storage) GetAccess(ctx context.Context, req model.AccessRequest) (model.AccessRequest, error) {
	query := fmt.Sprintf(`SELECT role, url, is_access FROM %s WHERE role = $1 AND url=$2`, tableUrlAccess)

	var res model.AccessRequest
	err := s.db.DB().ScanOneContext(ctx, &res, db.Query{
		Name:     "get_access",
		QueryRaw: query,
	}, req.Role, req.URL)
	if err != nil {
		return res, err
	}

	return res, nil
}

// GetAllAccess - получить все доступы
func (s *storage) GetAllAccess(ctx context.Context) ([]model.AccessRequest, error) {
	query := fmt.Sprintf(`SELECT role, url, is_access FROM %s `, tableUrlAccess)

	var res []model.AccessRequest
	err := s.db.DB().ScanOneContext(ctx, &res, db.Query{
		Name:     "get_access",
		QueryRaw: query,
	})
	if err != nil {
		return res, err
	}

	return res, nil
}
