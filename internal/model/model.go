package model

import "time"

// User - основная модель пользователя
type User struct {
	ID       int64
	Name     string
	Email    string
	Password string
	Role     int64
	CreateAt time.Time
	UpdateAt *time.Time
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
