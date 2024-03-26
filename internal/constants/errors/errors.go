package model

import "errors"

// Кастомные ошибки
var (
	ErrorPasswordMismatch     = errors.New("password mismatch")
	ErrorsOldPasswordNotValid = errors.New("old password not valid")
)
