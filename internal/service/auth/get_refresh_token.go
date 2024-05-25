package auth

import (
	"context"
	"github.com/Shemistan/grpc_user_api/internal/logger"
	"go.uber.org/zap"

	"github.com/Shemistan/grpc_user_api/internal/model"
	serviceErrors "github.com/Shemistan/grpc_user_api/internal/model/service_errors"
)

// GetRefreshToken - выписать новый refresh token
func (s *service) GetRefreshToken(_ context.Context, req string) (string, error) {
	claims, err := s.tokenProvider.VerifyToken(req, []byte(s.config.RefreshTokenSecretKey))
	if err != nil {
		logger.Error("failed to get refresh token:", zap.String("error", serviceErrors.ErrRefreshTokenInvalid.Error()))

		return "", serviceErrors.ErrRefreshTokenInvalid
	}

	refreshToken, err := s.tokenProvider.GenerateToken(model.UserInfo{
		Login: claims.Login,
		Role:  claims.Role,
	},
		[]byte(s.config.RefreshTokenSecretKey),
		s.config.RefreshTokenExpiration,
	)
	if err != nil {
		logger.Error("failed to get refresh token:", zap.String("error", serviceErrors.ErrAccessTokenInvalid.Error()))

		return "", err
	}

	return refreshToken, nil
}
