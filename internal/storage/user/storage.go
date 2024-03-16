package user

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/Shemistan/grpc_user_api/internal/client/db"
	def "github.com/Shemistan/grpc_user_api/internal/storage"
	"github.com/Shemistan/grpc_user_api/internal/storage/user/model"
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
func (s *storage) Create(ctx context.Context, req model.User) (int64, error) {
	query := `INSERT INTO users( name,email, password, role) VALUES ( $1, $2,$3,$4) RETURNING(id);`

	var id int64
	err := s.db.DB().QueryRowContext(ctx, db.Query{
		Name:     "create_user",
		QueryRaw: query,
	}, req.Name, req.Email, req.Password, req.Role).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *storage) Update(ctx context.Context, req model.UpdateUser) error {
	query := `UPDATE users SET role=$1 %s,  updated_at=now() WHERE id=$2`
	args := []interface{}{req.Role, req.ID}
	placeholderIdx := 3
	var clauses []string

	addClause := func(clause string, value interface{}) {
		clauses = append(clauses, clause)
		args = append(args, value)
		placeholderIdx++
	}

	if req.Name != nil {
		addClause(", name=$"+strconv.Itoa(placeholderIdx), *req.Name)
	}

	if req.Email != nil {
		addClause(", email=$"+strconv.Itoa(placeholderIdx), *req.Email)
	}

	if req.Password != nil {
		addClause(", password=$"+strconv.Itoa(placeholderIdx), *req.Password)
	}

	query = fmt.Sprintf(query, strings.Join(clauses, " "))

	_, err := s.db.DB().ExecContext(ctx, db.Query{
		Name:     "update_user",
		QueryRaw: query,
	}, args...)
	return err
}

// GetUser - получить пользователя
func (s *storage) GetUser(ctx context.Context, id int64) (model.User, error) {
	query := `SELECT  name, email, password, role, created_at, updated_at FROM users WHERE id = $1`

	var user model.User
	err := s.db.DB().ScanOneContext(ctx, &user, db.Query{
		Name:     "get_user",
		QueryRaw: query,
	}, id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *storage) GetPasswordHash(ctx context.Context, id int64) (string, error) {
	query := `SELECT   password FROM users WHERE id = $1`

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
