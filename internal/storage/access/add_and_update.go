package user

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/Shemistan/platform_common/pkg/db"

	"github.com/Shemistan/grpc_user_api/internal/model"
)

// AddAccess - добавить информацию о доступе
func (s *storage) AddAccess(ctx context.Context, req model.AccessRequest) error {
	qb := squirrel.
		Insert(tableResourceAccess).
		Columns("role", "resource", "is_access").
		Values(req.Role, req.Resource, req.IsAccess)

	query, args, err := qb.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return err
	}

	_, err = s.db.DB().ExecContext(ctx, db.Query{
		Name:     "add_access",
		QueryRaw: query,
	}, args...)
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

// UpsertAccess - добавить или редактировать доступ
func (s *storage) UpsertAccess(ctx context.Context, req model.AccessRequest) error {
	qb := squirrel.
		Insert(tableResourceAccess).
		Columns("role", "resource", "is_access").
		Values(req.Role, req.Resource, req.IsAccess).
		Suffix("ON CONFLICT (role, resource) DO UPDATE SET role = EXCLUDED.role, resource = EXCLUDED.resource, is_access = EXCLUDED.is_access")

	query, args, err := qb.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return err
	}

	_, err = s.db.DB().ExecContext(ctx, db.Query{
		Name:     "upsert_access",
		QueryRaw: query,
	}, args...)
	if err != nil {
		return err
	}

	return nil
}
