package user

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/Shemistan/grpc_user_api/internal/model"
	def "github.com/Shemistan/grpc_user_api/internal/storage"
)

type storage struct {
	//db *sqlx.DB
	db *pgxpool.Pool
}

// NewStorage - новый storage
func NewStorage(db *pgxpool.Pool) def.User {
	return &storage{db: db}
}

// Create - создать пользователя
func (s *storage) Create(ctx context.Context, req model.User) (int64, error) {
	query := `INSERT INTO users( name,email, password, role) VALUES ( $1, $2,$3,$4) RETURNING(id);`

	var id int64
	err := s.db.QueryRow(ctx, query, req.Name, req.Email, req.Password, req.Role).Scan(&id)
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

	if req.NewPassword != nil {
		addClause(", password=$"+strconv.Itoa(placeholderIdx), *req.NewPassword)
	}

	query = fmt.Sprintf(query, strings.Join(clauses, " "))

	_, err := s.db.Exec(ctx, query, args...)
	return err
}

// GetUser - получить пользователя
func (s *storage) GetUser(ctx context.Context, id int64) (model.User, error) {
	query := `SELECT  name, email, password, role, created_at, updated_at FROM users WHERE id = $1`

	var createdAt, updatedAt sql.NullTime
	var name, email, password string
	var role int64
	err := s.db.QueryRow(ctx, query, id).Scan(&name, &email, &password, &role, &createdAt, &updatedAt)
	if err != nil {
		return model.User{}, err
	}

	return model.User{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: password,
		Role:     role,
		CreateAt: createdAt.Time,
		UpdateAt: updatedAt.Time,
	}, nil
}

func (s *storage) GetPasswordHash(ctx context.Context, id int64) (string, error) {
	query := `SELECT   password FROM users WHERE id = $1`

	var password string
	err := s.db.QueryRow(ctx, query, id).Scan(&password)
	if err != nil {
		return "", err
	}

	return password, nil

}

// Delete - удалить пользователя
func (s *storage) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM users WHERE id =$1`

	_, err := s.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
