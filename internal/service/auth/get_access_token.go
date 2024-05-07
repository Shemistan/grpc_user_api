package auth

import (
	"context"

	"github.com/Shemistan/grpc_user_api/internal/model"
	serviceErrors "github.com/Shemistan/grpc_user_api/internal/model/service_errors"
)

// GetAccessToken - выписать новый access token
func (s *service) GetAccessToken(_ context.Context, req string) (string, error) {
	claims, err := s.tokenProvider.VerifyToken(req, []byte(s.config.AccessTokenSecretKey))
	if err != nil {
		return "", serviceErrors.ErrAccessTokenInvalid
	}

	accessToken, err := s.tokenProvider.GenerateToken(model.UserInfo{
		Login: claims.Login,
		Role:  claims.Role,
	},
		[]byte(s.config.AccessTokenSecretKey),
		s.config.AccessTokenExpiration,
	)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
