package user

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/Shemistan/platform_common/pkg/db"

	"github.com/Shemistan/grpc_user_api/internal/model"
)

// AddAccess - добавить информацию о доступе
func (s *storage) AddAccess(ctx context.Context, req model.AccessRequest) error {
	query := fmt.Sprintf(`INSERT INTO %s( role ,resource, is_access) VALUES ( $1, $2,$3);`, tableResourceAccess)

	_, err := s.db.DB().ExecContext(ctx, db.Query{
		Name:     "add_access",
		QueryRaw: query,
	}, req.Role, req.Resource, req.IsAccess)
	if err != nil {
		return err
	}

	return nil
}

// UpdateAccess - изменить информацию о доступе
func (s *storage) UpdateAccess(ctx context.Context, req model.AccessRequest) error {
	qb := squirrel.Update(tableResourceAccess).
		Set("role", req.Role).
		Set("resource", req.Resource).
		Set("is_access", req.IsAccess)

	qb = qb.Where(squirrel.And{
		squirrel.Eq{"role": req.Role},
		squirrel.Eq{"resource": req.Resource},
	}).
		PlaceholderFormat(squirrel.Dollar)

	query, args, err := qb.ToSql()
	if err != nil {
		return err
	}

	_, err = s.db.DB().ExecContext(ctx, db.Query{
		Name:     "update_user",
		QueryRaw: query,
	}, args...)

	return err
}
