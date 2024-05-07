package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	//TokenAuthPrefix - префикс для токена в хэдере
	TokenAuthPrefix = "Bearer "
)

// User - основная модель пользователя
type User struct {
	ID              int64  `db:"id"`
	Name            string `db:"name"`
	Email           string `db:"email"`
	Password        string `db:"password"`
	PasswordConfirm string
	Role            int64      `db:"role"`
	CreatedAt       time.Time  `db:"create_at"`
	UpdatedAt       *time.Time `db:"update_at"`
}

// UpdateUser - модель пользователя для редактирования
type UpdateUser struct {
	ID                 int64
	Name               *string
	Email              *string
	OldPassword        *string
	NewPassword        *string
	NewPasswordConfirm *string
	Role               int64
}

// UserInfo - информация о пользователе для авторизации
type UserInfo struct {
	Login string `json:"login"`
	Role  int64  `json:"role"`
}

// UserClaims - структура для работы токена
type UserClaims struct {
	jwt.StandardClaims
	Login string `json:"login"`
	Role  int64  `json:"role"`
}

// LoginRequest - запрос для авторизации
type LoginRequest struct {
	Login    string
	Password string
}

// LoginResponse - ответ при авторизации
type LoginResponse struct {
	RefreshToken string
	AccessToken  string
}

// AccessRequest - запрос для обработки доступов
type AccessRequest struct {
	Role     int64  `db:"role"`
	Resource string `db:"resource"`
	IsAccess bool   `db:"is_access"`
}

// GetUserRequest - запрос для получения пользователя
type GetUserRequest struct {
	ID    *int64
	Email *string
}
