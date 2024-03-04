package model

import "time"

type User struct {
	ID         int64     `db:"id"`
	Name       string    `db:"name"`
	Email      string    `db:"email"`
	Password   string    `db:"password"`
	Role       int64     `db:"role"`
	CreateAt   time.Time `db:"create_at"`
	UpdateDate time.Time `db:"update_at"`
}
