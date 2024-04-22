package service_errors

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	// ErrPasswordMismatch - пароли не совпадают
	ErrPasswordMismatch = errors.New("password mismatch")
	// ErrOldPasswordNotValid - прежний пароль не правильный
	ErrOldPasswordNotValid = errors.New("old password not valid")
	// ErrOldPasswordNotFound - прежний пароль не найден
	ErrOldPasswordNotFound = errors.New("old password not found")
	// ErrPasswordNotValid - прежний пароль не правильный
	ErrPasswordNotValid = errors.New("password not valid")
	// ErrAccessDenied - нет доступа
	ErrAccessDenied = errors.New("access denied")
	// ErrMetadataIsNotProvided - метаданные не предоставлены
	ErrMetadataIsNotProvided = errors.New("metadata is not provided")
	// ErrAuthHeaderIsNotProvided - authorization header не предоставлены
	ErrAuthHeaderIsNotProvided = errors.New("authorization header is not provided")
	// ErrAuthHeaderFormat - неверный формат хэдера авторизации
	ErrAuthHeaderFormat = errors.New("invalid authorization header format")
	// ErrAccessTokenInvalid - невалидный токен
	ErrAccessTokenInvalid = status.Errorf(codes.Aborted, "invalid access token")
	// ErrRefreshTokenInvalid - невалидный токен
	ErrRefreshTokenInvalid = status.Errorf(codes.Aborted, "invalid refresh token")
	// ErrGenerateToken - не удалось создать токен
	ErrGenerateToken = errors.New("failed to generate token")
	// ErrNoRows - нет соотвествующих данных в хранилище
	ErrNoRows = errors.New("no rows in result")
)
