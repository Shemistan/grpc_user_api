package user

import (
	"context"

	"github.com/Masterminds/squirrel"

	"github.com/Shemistan/grpc_user_api/internal/client/db"
	"github.com/Shemistan/grpc_user_api/internal/model"
	def "github.com/Shemistan/grpc_user_api/internal/storage"
	"github.com/Shemistan/grpc_user_api/internal/storage/user/converter"
	storageModel "github.com/Shemistan/grpc_user_api/internal/storage/user/model"
)

type storage struct {
	db        db.Client
	txManager db.TxManager
}

// NewStorage - новый storage
func NewStorage(db db.Client, txManager db.TxManager) def.User {
	return &storage{
		db:        db,
		txManager: txManager,
	}
}

// Create - создать пользователя
func (s *storage) Create(ctx context.Context, req model.User, passwordHash string) (int64, error) {
	user := converter.ServiceUserToStorageUser(req, passwordHash)

	query := `INSERT INTO users( name,email, password, role) VALUES ( $1, $2,$3,$4) RETURNING(id);`

	var id int64
	err := s.db.DB().QueryRowContext(ctx, db.Query{
		Name:     "create_user",
		QueryRaw: query,
	}, user.Name, user.Email, user.Password, user.Role).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *storage) Update(ctx context.Context, req model.UpdateUser, passwordHash *string) error {
	user := converter.ServiceUpdateUserToStorageUpdateUser(req, passwordHash)

	qb := squirrel.Update("users").Set("role", user.Role)

	if user.Name != nil {
		qb = qb.Set("name", *user.Name)
	}

	if user.Email != nil {
		qb = qb.Set("email", *user.Email)
	}

	if user.Password != nil {
		qb = qb.Set("password", *user.Password)
	}

	qb = qb.Where(squirrel.Eq{"id": user.ID}).PlaceholderFormat(squirrel.Dollar)

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
func (s *storage) GetUser(ctx context.Context, id int64) (model.User, error) {
	query := `SELECT  name, email, password, role, created_at, updated_at FROM users WHERE id = $1`

	var user storageModel.User
	err := s.db.DB().ScanOneContext(ctx, &user, db.Query{
		Name:     "get_user",
		QueryRaw: query,
	}, id)
	if err != nil {
		return model.User{}, err
	}

	return converter.StorageUserToServiceUser(user), nil
}

// GetPasswordHash - получить hash пароля
func (s *storage) GetPasswordHash(ctx context.Context, id int64) (string, error) {
	query := `SELECT password FROM users WHERE id = $1`

	var password string
	err := s.db.DB().QueryRowContext(ctx, db.Query{
		Name:     "get_password_hash",
		QueryRaw: query,
	}, id).Scan(&password)
	if err != nil {
		return "", err
	}

	return password, nil
}

// Delete - удалить пользователя
func (s *storage) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM users WHERE id =$1`

	_, err := s.db.DB().ExecContext(ctx, db.Query{
		Name:     "delete_user",
		QueryRaw: query,
	}, id)
	if err != nil {
		return err
	}

	return nil
}
