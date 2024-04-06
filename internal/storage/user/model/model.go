package model

import (
	"database/sql"
)

// User - основная модель пользователя
type User struct {
	ID        int64        `db:"id"`
	Name      string       `db:"name"`
	Email     string       `db:"email"`
	Password  string       `db:"password"`
	Role      int64        `db:"role"`
	CreatedAt sql.NullTime `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

// UpdateUser - модель пользователя для редактирования
type UpdateUser struct {
	ID       int64   `db:"id"`
	Name     *string `db:"name"`
	Email    *string `db:"email"`
	Password *string `db:"password"`
	Role     int64   `db:"role"`
}
