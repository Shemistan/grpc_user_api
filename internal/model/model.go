package model

import "time"

// User - основная модель пользователя
type User struct {
	ID       int64     `db:"id"`
	Name     string    `db:"name"`
	Email    string    `db:"email"`
	Password string    `db:"password"`
	Role     int64     `db:"role"`
	CreateAt time.Time `db:"create_at"`
	UpdateAt time.Time `db:"update_at"`
}

// UpdateUser - модель пользователя для редактирования
type UpdateUser struct {
	ID          int64
	Name        *string
	Email       *string
	OldPassword *string
	NewPassword *string
	Role        int64
}
