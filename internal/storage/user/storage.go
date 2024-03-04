package user

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/Shemistan/grpc_user_api/internal/model"
	def "github.com/Shemistan/grpc_user_api/internal/storage"
)

type storage struct {
	db *sqlx.DB
}

func NewStorage(db *sqlx.DB) def.User {
	return &storage{db: db}
}

// Create - создать пользователя
func (s *storage) Create(ctx context.Context, req model.User) (int64, error) {
	query := `INSERT INTO users( name,email, password, role) VALUES ( $1, $2);`

	var id int64
	err := s.db.QueryRowContext(ctx, query, req.Name, req.Email, req.Password, req.Role).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Create - редактировать пользователя
func (s *storage) Update(ctx context.Context, req model.User) error {
	query := `UPDATE users SET name=$1 ,email= $2, password=$3, role=$4, updated_at=now() WHERE id=$5`

	_, err := s.db.ExecContext(ctx, query, req.Name, req.Email, req.Password, req.Role, req.ID)
	if err != nil {
		return err
	}

	return nil
}

// Create - получить пользователя
func (s *storage) GetUser(ctx context.Context, id int64) (model.User, error) {
	query := `SELECT name, email, password, role, created_at, updated_at FROM users WHERE id = $1`

	var res model.User
	err := s.db.SelectContext(ctx, &res, query, id)
	if err != nil {
		return res, err
	}

	return res, nil
}

// Create - удалить пользователя
func (s *storage) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM users WHERE id =$1`

	_, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
