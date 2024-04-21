package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	ExamplePath     = "/auth_v1.AuthV1/Get"
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

type UserInfo struct {
	Login string `json:"login"`
	Role  int64  `json:"role"`
}

type UserClaims struct {
	jwt.StandardClaims
	Login string `json:"login"`
	Role  int64  `json:"role"`
}

type LoginRequest struct {
	Login    string
	Password string
}

type LoginResponse struct {
	RefreshToken string
	AccessToken  string
}

type AccessRequest struct {
	Role     int64  `db:"role"`
	URL      string `db:"url"`
	IsAccess bool   `db:"is_access"`
}
