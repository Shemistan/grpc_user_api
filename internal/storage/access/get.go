package user

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/Shemistan/platform_common/pkg/db"

	"github.com/Shemistan/grpc_user_api/internal/model"
)

// GetAccess - получить информацию о доступе
func (s *storage) GetAccess(ctx context.Context, req model.AccessRequest) (model.AccessRequest, error) {
	var res model.AccessRequest

	qb := squirrel.
		Select(role, resource, isAccess).
		From(tableResourceAccess).
		Where(squirrel.Eq{role: req.Role, resource: req.Resource})

	query, args, err := qb.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return res, err
	}

	err = s.db.DB().ScanOneContext(ctx, &res, db.Query{
		Name:     "get_access",
		QueryRaw: query,
	}, args...)
	if err != nil {
		return res, err
	}

	return res, nil
}

// GetAllAccess - получить все доступы
func (s *storage) GetAllAccess(ctx context.Context) ([]model.AccessRequest, error) {
	qb := squirrel.
		Select(role, resource, isAccess).
		From(tableResourceAccess)

	query, _, err := qb.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	var res []model.AccessRequest
	err = s.db.DB().ScanAllContext(ctx, &res, db.Query{
		Name:     "get_all_access",
		QueryRaw: query,
	})
	if err != nil {
		return res, err
	}

	return res, nil
}
