package app

import (
	"context"
	"log"

	"github.com/Shemistan/platform_common/pkg/closer"
	"github.com/Shemistan/platform_common/pkg/db"
	"github.com/Shemistan/platform_common/pkg/db/pg"
	"github.com/Shemistan/platform_common/pkg/db/transaction"

	userAPI "github.com/Shemistan/grpc_user_api/internal/api/user_v1"
	"github.com/Shemistan/grpc_user_api/internal/config"
	"github.com/Shemistan/grpc_user_api/internal/config/env"
	"github.com/Shemistan/grpc_user_api/internal/service"
	userService "github.com/Shemistan/grpc_user_api/internal/service/user"
	"github.com/Shemistan/grpc_user_api/internal/storage"
	userStorage "github.com/Shemistan/grpc_user_api/internal/storage/user"
	"github.com/Shemistan/grpc_user_api/internal/utils"
	"github.com/Shemistan/grpc_user_api/internal/utils/hasher"
)

type serviceProvider struct {
	pgConfig         config.PGConfig
	grpcConfig       config.GRPCConfig
	httpConfig       config.HTTP
	swaggerConfig    config.Swagger
	secretHashConfig config.SecretHashConfig

	dbClient  db.Client
	txManager db.TxManager

	userStorage storage.User
	userService service.User
	userAPI     *userAPI.User

	passwordHasher utils.Hasher
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
	if s.grpcConfig == nil {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
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

func (s *serviceProvider) PasswordHasher(_ context.Context) utils.Hasher {
	if s.passwordHasher == nil {
		s.passwordHasher = hasher.New(s.SecretHashConfig().PasswordHashKey())
	}

	return s.passwordHasher
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

func (s *serviceProvider) UserAPI(ctx context.Context) *userAPI.User {
	if s.userAPI == nil {
		s.userAPI = userAPI.New(s.UserService(ctx))
	}

	return s.userAPI
}
