package auth

import (
	"context"
	"database/sql"
	"errors"
)

// AddActualValuesInCache - заполнить кэш актуальными значениями из БД
func (s *service) AddActualValuesInCache(ctx context.Context) error {
	roles, err := s.accessStorage.GetAllAccess(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		return err
	}

	for _, v := range roles {
		s.cache.AddInCache(v)
	}

	return nil
}
