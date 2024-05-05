package user

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/Shemistan/platform_common/pkg/db"

	"github.com/Shemistan/grpc_user_api/internal/model"
	def "github.com/Shemistan/grpc_user_api/internal/storage"
	"github.com/Shemistan/grpc_user_api/internal/storage/user/converter"
	storageModel "github.com/Shemistan/grpc_user_api/internal/storage/user/model"
)

type storage struct {
	db        db.Client
	txManager db.TxManager
}

const (
	tableUsers      = "users"
	columnRole      = "role"
	columnName      = "name"
	columnEmail     = "email"
	columnPassword  = "password"
	columnID        = "id"
	columnCreatedAt = "id"
	columnUpdatedAt = "id"
)

// NewStorage - новый storage
func NewStorage(db db.Client, txManager db.TxManager) def.User {
	return &storage{
		db:        db,
		txManager: txManager,
	}
}

// Update - редактировать пользователя
func (s *storage) Update(ctx context.Context, req model.UpdateUser, passwordHash *string) error {
	user := converter.ServiceUpdateUserToStorageUpdateUser(req, passwordHash)

	qb := squirrel.Update(tableUsers).Set(columnRole, user.Role)

	if user.Name != nil {
		qb = qb.Set(columnName, *user.Name)
	}

	if user.Email != nil {
		qb = qb.Set(columnEmail, *user.Email)
	}

	if user.Password != nil {
		qb = qb.Set(columnPassword, *user.Password)
	}

	qb = qb.Where(squirrel.Eq{columnID: user.ID}).PlaceholderFormat(squirrel.Dollar)

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

// GetUser - получить пользователя
func (s *storage) GetUser(ctx context.Context, req model.GetUserRequest) (model.User, error) {
	qb := squirrel.Select(columnID, columnName, columnEmail, columnPassword, columnRole, columnCreatedAt, columnUpdatedAt).
		From(tableUsers)

	if req.ID != nil {
		qb = qb.Where(squirrel.And{
			squirrel.Eq{
				columnID: *req.ID,
			},
		})
	}

	if req.Email != nil {
		qb = qb.Where(squirrel.And{
			squirrel.Eq{
				columnEmail: *req.Email,
			},
		})
	}

	query, args, err := qb.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return model.User{}, err
	}

	var user storageModel.User
	err = s.db.DB().ScanOneContext(ctx, &user, db.Query{
		Name:     "get_user",
		QueryRaw: query,
	}, args...)
	if err != nil {
		return model.User{}, err
	}

	return converter.StorageUserToServiceUser(user), nil
}

// Create - создать пользователя
func (s *storage) Create(ctx context.Context, req model.User, passwordHash string) (int64, error) {
	user := converter.ServiceUserToStorageUser(req, passwordHash)

	qb := squirrel.
		Insert(tableUsers).
		Columns(columnName, columnEmail, columnPassword, columnRole).
		Values(user.Name, user.Email, user.Password, user.Role).
		Suffix("RETURNING id")

	query, args, err := qb.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return 0, err
	}

	var id int64
	err = s.db.DB().QueryRowContext(ctx, db.Query{
		Name:     "create_user",
		QueryRaw: query,
	}, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// GetPasswordHash - получить hash пароля
func (s *storage) GetPasswordHash(ctx context.Context, id int64) (string, error) {
	qb := squirrel.
		Select(columnPassword).
		From(tableUsers).
		Where(squirrel.Eq{columnID: id})

	query, args, err := qb.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return "", err
	}

	var password string
	err = s.db.DB().QueryRowContext(ctx, db.Query{
		Name:     "get_password_hash",
		QueryRaw: query,
	}, args...).Scan(&password)
	if err != nil {
		return "", err
	}

	return password, nil
}

// Delete - удалить пользователя
func (s *storage) Delete(ctx context.Context, id int64) error {
	qb := squirrel.
		Delete(tableUsers).
		Where(squirrel.Eq{columnID: id})

	query, args, err := qb.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return err
	}

	_, err = s.db.DB().ExecContext(ctx, db.Query{
		Name:     "delete_user",
		QueryRaw: query,
	}, args...)
	if err != nil {
		return err
	}

	return nil
}
