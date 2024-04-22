package auth

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/Shemistan/grpc_user_api/internal/model"
	serviceErrors "github.com/Shemistan/grpc_user_api/internal/model/service_errors"
)

func (s *service) AddOrUpdateAccess(ctx context.Context, req model.AccessRequest) error {
	access, err := s.accessStorage.GetAccess(ctx, req)
	if err != nil {
		if !strings.HasPrefix(err.Error(), serviceErrors.ErrNoRows.Error()) {
			log.Println(fmt.Sprintf("failed  to add new access for role(%d) and resource(%s)", req.Role, req.Resource), err)
			return err
		}

		err = s.accessStorage.AddAccess(ctx, req)
		if err != nil {
			log.Println(fmt.Sprintf("failed  to add new access for role(%d) and resource(%s)", req.Role, req.Resource), err)
			return err
		}

		access = req
	}

	if access.IsAccess != req.IsAccess {
		err = s.accessStorage.UpdateAccess(ctx, req)
		if err != nil {
			log.Println(fmt.Sprintf("failed  to update access for role(%d) and resource(%s)", req.Role, req.Resource), err)
			return err
		}
	}

	s.addInCache(req)

	return nil
}
