package service_errors

import "errors"

var (
	// ErrPasswordMismatch - пароли не совпадают
	ErrPasswordMismatch = errors.New("password mismatch")
	// ErrOldPasswordNotValid - прежний пароль не правильный
	ErrOldPasswordNotValid = errors.New("old password not valid")
	// ErrOldPasswordNotFound - прежний пароль не найден
	ErrOldPasswordNotFound = errors.New("old password not found")
)
