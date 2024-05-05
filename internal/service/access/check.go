package auth

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/jackc/pgx/v4"
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
	accessesMap := s.cache.GetAccessesForRole(role)
	if accessesMap != nil {
		if isAccess, okResource := accessesMap[resource]; okResource {
			return isAccess, nil
		}
	}

	access, err := s.accessStorage.GetAccess(ctx, model.AccessRequest{
		Role:     role,
		Resource: resource,
	})
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return false, err
		}

		go func(ctx context.Context, role int64, resource string) {
			req := model.AccessRequest{
				Role:     role,
				Resource: resource,
			}
			errAdd := s.accessStorage.AddAccess(ctx, req)
			if errAdd != nil {
				log.Println(fmt.Sprintf("failed  to add new access for role(%d) and resource(%s)", role, resource), err)
				return
			}

			s.cache.AddInCache(req)
		}(context.Background(), role, resource)

		return false, nil
	}

	s.cache.AddInCache(access)

	return access.IsAccess, nil
}
