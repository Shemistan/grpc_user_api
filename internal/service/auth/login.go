package auth

import (
	"context"

	"go.uber.org/zap"

	"github.com/Shemistan/grpc_user_api/internal/logger"
	"github.com/Shemistan/grpc_user_api/internal/model"
	serviceErrors "github.com/Shemistan/grpc_user_api/internal/model/service_errors"
)

// Login - авторизоваться
func (s *service) Login(ctx context.Context, req model.LoginRequest) (model.LoginResponse, error) {
	var res model.LoginResponse

	user, err := s.userStorage.GetUser(ctx, model.GetUserRequest{
		ID:    nil,
		Email: &req.Login,
	})
	if err != nil {
		logger.Error("failed to authorize user:", zap.String("error", err.Error()))

		return res, err
	}

	ok := s.hasher.CheckPassword(user.Password, req.Password)
	if !ok {
		logger.Error("failed to authorize user:", zap.String("error", serviceErrors.ErrPasswordNotValid.Error()))

		return res, serviceErrors.ErrPasswordNotValid
	}

	res.RefreshToken, err = s.tokenProvider.GenerateToken(model.UserInfo{
		Login: req.Login,
		Role:  user.Role,
	},
		[]byte(s.config.RefreshTokenSecretKey),
		s.config.RefreshTokenExpiration,
	)
	if err != nil {
		logger.Error("failed to authorize user:", zap.String("error", serviceErrors.ErrGenerateToken.Error()))

		return res, serviceErrors.ErrGenerateToken
	}

	res.AccessToken, err = s.tokenProvider.GenerateToken(model.UserInfo{
		Login: req.Login,
		Role:  user.Role,
	},
		[]byte(s.config.AccessTokenSecretKey),
		s.config.AccessTokenExpiration,
	)
	if err != nil {
		logger.Error("failed to authorize user:", zap.String("error", serviceErrors.ErrGenerateToken.Error()))

		return res, serviceErrors.ErrGenerateToken
	}

	return res, nil
}
