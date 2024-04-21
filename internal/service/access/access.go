package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"

	"google.golang.org/grpc/metadata"

	"github.com/Shemistan/grpc_user_api/internal/model"
	serviceErrors "github.com/Shemistan/grpc_user_api/internal/model/service_errors"
)

// Check - проверить доступ
func (s *service) Check(ctx context.Context, url string) error {
	claims, err := s.checkTokenAndGetClaims(ctx)
	if err != nil {
		return err
	}

	ok, err := s.checkAccessibleRoles(ctx, claims.Role, url)
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

func (s *service) checkAccessibleRoles(ctx context.Context, role int64, url string) (bool, error) {
	if s.cache == nil {
		err := s.initAccessibleRoles(ctx)
		if err != nil {
			log.Println("failed to init accessible roles", err)
			return false, serviceErrors.ErrCheckAccess
		}
	}

	if v, okRole := s.cache.accessibleRoles[role]; okRole {
		if isAccess, okUrl := v[url]; okUrl {
			return isAccess, nil
		}
	}

	access, err := s.accessStorage.GetAccess(ctx, model.AccessRequest{
		Role: role,
		URL:  url,
	})
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return false, err
		}

		go func(ctx context.Context, role int64, url string) {
			errAdd := s.addNewAccess(ctx, model.AccessRequest{
				Role: role,
				URL:  url,
			})

			if errAdd != nil {
				log.Println(fmt.Sprintf("failed  to add new access for role(%d) and url(%s)", role, url), err)
			}

			s.cache.Lock()
			s.cache.accessibleRoles[role][url] = false
			s.cache.Unlock()
		}(ctx, role, url)

		return false, nil
	}

	return access.IsAccess, nil
}

func (s *service) initAccessibleRoles(ctx context.Context) error {
	roles, err := s.accessStorage.GetAllAccess(ctx)
	if err != nil {
		return err
	}

	// Ролей предполагается не так много, по этому эффективнее сделать такой кэш, что бы не плодить много хэштаблиц
	rolesMap := make(map[int64]map[string]bool)

	for _, v := range roles {
		if _, ok := rolesMap[v.Role]; !ok {
			urlsMap := make(map[string]bool)
			rolesMap[v.Role] = urlsMap
		}

		rolesMap[v.Role][v.URL] = v.IsAccess
	}

	s.cache = &Cache{
		Mutex:           &sync.Mutex{},
		accessibleRoles: rolesMap,
	}

	return nil
}

func (s *service) addNewAccess(ctx context.Context, req model.AccessRequest) error {
	err := s.accessStorage.AddAccess(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
