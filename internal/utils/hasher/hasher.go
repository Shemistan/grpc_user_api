package hasher

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/Shemistan/grpc_user_api/internal/utils"
)

// hasher структура, реализующая интерфейс Hasher.
type hasher struct {
	secretKey string
}

// New создает новый экземпляр service с заданным секретным ключом.
func New(secretKey string) utils.Hasher {
	return &hasher{secretKey: secretKey}
}

// GetPasswordHash генерирует хеш пароля с использованием bcrypt.
func (s *hasher) GetPasswordHash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password+s.secretKey), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPassword проверяет, соответствует ли пароль хешу.
func (s *hasher) CheckPassword(hash, password string) bool {
	// Сравниваем предоставленный пароль с хешем.
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password+s.secretKey))
	return err == nil
}
