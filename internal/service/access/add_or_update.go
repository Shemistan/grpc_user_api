package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/Shemistan/grpc_user_api/internal/model"
)

func (s *service) AddOrUpdateAccess(ctx context.Context, req model.AccessRequest) error {
	access, err := s.accessStorage.GetAccess(ctx, req)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Println(fmt.Sprintf("failed  to add new access for role(%d) and url(%s)", req.Role, req.URL), err)
			return err
		}
		err = s.accessStorage.AddAccess(ctx, req)
		if err != nil {
			log.Println(fmt.Sprintf("failed  to add new access for role(%d) and url(%s)", req.Role, req.URL), err)
			return err
		}

		access = req
	}

	if access.IsAccess != req.IsAccess {
		err = s.accessStorage.UpdateAccess(ctx, req)
		if err != nil {
			log.Println(fmt.Sprintf("failed  to update access for role(%d) and url(%s)", req.Role, req.URL), err)
			return err
		}
	}

	s.cache.Lock()
	s.cache.accessibleRoles[req.Role][req.URL] = req.IsAccess
	s.cache.Unlock()

	return nil
}
