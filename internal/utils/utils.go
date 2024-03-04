package utils

// Hasher - интерфейс хэшера паролей
type Hasher interface {
	GetPasswordHash(password string) (string, error)
	CheckPassword(hash, password string) bool
}
