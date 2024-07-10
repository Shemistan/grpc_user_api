package app

import (
	"context"
	"log"

	"github.com/Shemistan/platform_common/pkg/closer"
	"github.com/Shemistan/platform_common/pkg/db"
	"github.com/Shemistan/platform_common/pkg/db/pg"
	"github.com/Shemistan/platform_common/pkg/db/transaction"

	accessAPI "github.com/Shemistan/grpc_user_api/internal/api/access_v1"
	authAPI "github.com/Shemistan/grpc_user_api/internal/api/auth_v1"
	userAPI "github.com/Shemistan/grpc_user_api/internal/api/user_v1"
	"github.com/Shemistan/grpc_user_api/internal/config"
	"github.com/Shemistan/grpc_user_api/internal/config/env"
	"github.com/Shemistan/grpc_user_api/internal/service"
	accessService "github.com/Shemistan/grpc_user_api/internal/service/access"
	authService "github.com/Shemistan/grpc_user_api/internal/service/auth"
	userService "github.com/Shemistan/grpc_user_api/internal/service/user"
	"github.com/Shemistan/grpc_user_api/internal/storage"
	accessStorage "github.com/Shemistan/grpc_user_api/internal/storage/access"
	"github.com/Shemistan/grpc_user_api/internal/storage/access_cache"
	userStorage "github.com/Shemistan/grpc_user_api/internal/storage/user"
	"github.com/Shemistan/grpc_user_api/internal/utils"
	"github.com/Shemistan/grpc_user_api/internal/utils/hasher"
	utilsToken "github.com/Shemistan/grpc_user_api/internal/utils/token"
)

type serviceProvider struct {
	pgConfig                 config.PGConfig
	grpcUserConfig           config.GRPCConfig
	httpConfig               config.HTTP
	swaggerConfig            config.Swagger
	secretHashConfig         config.SecretHashConfig
	secretRefreshTokenConfig config.SecretRefreshTokenConfig
	secretAccessTokenConfig  config.SecretAccessTokenConfig
	prometheusConfig         config.Prometheus
	loggerConfig             config.ZapLogger

	tokenServiceConfig *config.TokenServiceConfig

	dbClient  db.Client
	txManager db.TxManager

	userStorage storage.User
	userService service.User
	userAPI     *userAPI.User

	authService service.Auth
	authAPI     *authAPI.Auth

	accessStorage storage.Access
	accessService service.Access
	accessAPI     *accessAPI.Access

	cacheStorage storage.Cache

	passwordHasher utils.Hasher
	tokenProvider  utils.TokenProvider
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcUserConfig == nil {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcUserConfig = cfg
	}

	return s.grpcUserConfig
}

func (s *serviceProvider) SecretHashConfig() config.SecretHashConfig {
	if s.secretHashConfig == nil {
		cfg, err := env.NewSecretHashConfig()
		if err != nil {
			log.Fatalf("failed to get secret hash config: %s", err.Error())
		}

		s.secretHashConfig = cfg
	}

	return s.secretHashConfig
}

func (s *serviceProvider) SecretRefreshTokenConfig() config.SecretRefreshTokenConfig {
	if s.secretRefreshTokenConfig == nil {
		cfg, err := env.NewSecretRefreshTokenConfig()
		if err != nil {
			log.Fatalf("failed to get refresh token config: %s", err.Error())
		}

		s.secretRefreshTokenConfig = cfg
	}

	return s.secretRefreshTokenConfig
}

func (s *serviceProvider) SecretAccessTokenConfig() config.SecretAccessTokenConfig {
	if s.secretAccessTokenConfig == nil {
		cfg, err := env.NewSecretAccessTokenConfig()
		if err != nil {
			log.Fatalf("failed to get access token config: %s", err.Error())
		}

		s.secretAccessTokenConfig = cfg
	}

	return s.secretAccessTokenConfig
}

func (s *serviceProvider) HTTPConfig() config.HTTP {
	if s.httpConfig == nil {
		cfg, err := env.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) SwaggerConfig() config.Swagger {
	if s.swaggerConfig == nil {
		cfg, err := env.NewSwaggerConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}

		s.swaggerConfig = cfg
	}

	return s.swaggerConfig
}

func (s *serviceProvider) PrometheusConfig() config.HTTP {
	if s.prometheusConfig == nil {
		cfg, err := env.NewPrometheusConfig()
		if err != nil {
			log.Fatalf("failed to get prometheus config: %s", err.Error())
		}

		s.prometheusConfig = cfg
	}

	return s.prometheusConfig
}

func (s *serviceProvider) LoggerConfig() config.ZapLogger {
	if s.loggerConfig == nil {
		cfg, err := env.NewZapLoggerConfig()
		if err != nil {
			log.Fatalf("failed to get logger config: %s", err.Error())
		}

		s.loggerConfig = cfg
	}

	return s.loggerConfig

}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) UserStorage(ctx context.Context) storage.User {
	if s.userStorage == nil {
		s.userStorage = userStorage.NewStorage(s.DBClient(ctx), s.TxManager(ctx))
	}

	return s.userStorage
}

func (s *serviceProvider) AccessStorage(ctx context.Context) storage.Access {
	if s.accessStorage == nil {
		s.accessStorage = accessStorage.NewStorage(s.DBClient(ctx), s.TxManager(ctx))
	}

	return s.accessStorage
}

func (s *serviceProvider) CacheStorage(_ context.Context) storage.Cache {
	if s.cacheStorage == nil {
		s.cacheStorage = access_cache.NewCache()
	}

	return s.cacheStorage
}

func (s *serviceProvider) PasswordHasher(_ context.Context) utils.Hasher {
	if s.passwordHasher == nil {
		s.passwordHasher = hasher.New(s.SecretHashConfig().PasswordHashKey())
	}

	return s.passwordHasher
}

func (s *serviceProvider) TokenProvider(_ context.Context) utils.TokenProvider {
	if s.tokenProvider == nil {
		s.tokenProvider = utilsToken.New()
	}

	return s.tokenProvider
}

func (s *serviceProvider) TokenServiceConfig(_ context.Context) *config.TokenServiceConfig {
	if s.tokenServiceConfig == nil {
		s.tokenServiceConfig = &config.TokenServiceConfig{
			RefreshTokenSecretKey:  s.SecretRefreshTokenConfig().SecretKey(),
			AccessTokenSecretKey:   s.SecretAccessTokenConfig().SecretKey(),
			RefreshTokenExpiration: config.RefreshTokenExpiration,
			AccessTokenExpiration:  config.AccessTokenExpiration,
		}
	}

	return s.tokenServiceConfig
}

func (s *serviceProvider) UserService(ctx context.Context) service.User {
	if s.userService == nil {
		s.userService = userService.NewService(
			s.UserStorage(ctx),
			s.PasswordHasher(ctx),
		)
	}

	return s.userService
}

func (s *serviceProvider) AuthService(ctx context.Context) service.Auth {
	if s.authService == nil {
		s.authService = authService.NewService(
			s.UserStorage(ctx),
			s.PasswordHasher(ctx),
			s.TokenProvider(ctx),
			s.TokenServiceConfig(ctx),
		)
	}

	return s.authService
}

func (s *serviceProvider) AccessService(ctx context.Context) service.Access {
	if s.accessService == nil {
		s.accessService = accessService.NewService(
			s.TokenProvider(ctx),
			s.AccessStorage(ctx),
			s.SecretAccessTokenConfig().SecretKey(),
			s.CacheStorage(ctx),
		)
	}

	return s.accessService
}

func (s *serviceProvider) UserAPI(ctx context.Context) *userAPI.User {
	if s.userAPI == nil {
		s.userAPI = userAPI.New(s.UserService(ctx))
	}

	return s.userAPI
}

func (s *serviceProvider) AuthAPI(ctx context.Context) *authAPI.Auth {
	if s.authAPI == nil {
		s.authAPI = authAPI.New(s.AuthService(ctx))
	}

	return s.authAPI
}

func (s *serviceProvider) AccessAPI(ctx context.Context) *accessAPI.Access {
	if s.accessAPI == nil {
		s.accessAPI = accessAPI.New(s.AccessService(ctx))
	}

	return s.accessAPI
}
