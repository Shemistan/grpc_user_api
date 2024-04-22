package auth

import (
	"context"
	"fmt"
	"log"
	"strings"

	"google.golang.org/grpc/metadata"

	"github.com/Shemistan/grpc_user_api/internal/model"
	serviceErrors "github.com/Shemistan/grpc_user_api/internal/model/service_errors"
)

// Check - проверить доступ
func (s *service) Check(ctx context.Context, resource string) error {
	claims, err := s.checkTokenAndGetClaims(ctx)
	if err != nil {
		return err
	}

	ok, err := s.checkAccessibleRoles(ctx, claims.Role, resource)
	if err != nil {
		return err
	}
	if !ok {
		return serviceErrors.ErrAccessDenied
	}

	return nil
}

func (s *service) checkTokenAndGetClaims(ctx context.Context) (*model.UserClaims, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, serviceErrors.ErrMetadataIsNotProvided
	}

	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return nil, serviceErrors.ErrAuthHeaderIsNotProvided
	}

	if !strings.HasPrefix(authHeader[0], model.TokenAuthPrefix) {
		return nil, serviceErrors.ErrAuthHeaderFormat
	}

	accessToken := strings.TrimPrefix(authHeader[0], model.TokenAuthPrefix)

	claims, err := s.tokenProvider.VerifyToken(accessToken, []byte(s.accessTokenSecretKey))
	if err != nil {
		return nil, serviceErrors.ErrAccessTokenInvalid
	}

	return claims, nil
}

func (s *service) checkAccessibleRoles(ctx context.Context, role int64, resource string) (bool, error) {
	if v, okRole := s.cache.accessibleRoles[role]; okRole {
		if isAccess, okResource := v[resource]; okResource {
			return isAccess, nil
		}
	}

	access, err := s.accessStorage.GetAccess(ctx, model.AccessRequest{
		Role:     role,
		Resource: resource,
	})
	if err != nil {
		if !strings.HasPrefix(err.Error(), serviceErrors.ErrNoRows.Error()) {
			return false, err
		}

		go func(ctx context.Context, role int64, resource string) {
			errAdd := s.addNewAccess(ctx, model.AccessRequest{
				Role:     role,
				Resource: resource,
			})

			if errAdd != nil {
				log.Println(fmt.Sprintf("failed  to add new access for role(%d) and resource(%s)", role, resource), err)
			}

			s.addInCache(model.AccessRequest{
				Role:     role,
				Resource: resource,
				IsAccess: false,
			})
		}(ctx, role, resource)

		return false, nil
	}

	return access.IsAccess, nil
}

func (s *service) addNewAccess(ctx context.Context, req model.AccessRequest) error {
	err := s.accessStorage.AddAccess(ctx, req)
	if err != nil {
		return err
	}

	s.addInCache(req)

	return nil
}
