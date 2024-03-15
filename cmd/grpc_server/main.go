package main

import (
	"context"
	"flag"
	"log"
	"net"

	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/Shemistan/grpc_user_api/internal/api"
	"github.com/Shemistan/grpc_user_api/internal/config"
	"github.com/Shemistan/grpc_user_api/internal/config/env"
	user2 "github.com/Shemistan/grpc_user_api/internal/service/user"
	"github.com/Shemistan/grpc_user_api/internal/storage/user"
	"github.com/Shemistan/grpc_user_api/internal/utils/hasher"
	pb "github.com/Shemistan/grpc_user_api/pkg/user_api_v1"
)

var configPath string

func init() {
	configPath = ".env"
	//flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

func main() {
	flag.Parse()
	ctx := context.Background()

	err := config.Load(configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	grpcConfig, err := env.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to get grpc config: %v", err)
	}

	pgConfig, err := env.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to get pg config: %v", err)
	}

	hashSecret, err := env.NewSecretHashConfig()
	if err != nil {
		log.Fatalf("failed to get hash secret config: %v", err)
	}

	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatalln("failed to listen: ", err.Error())
	}

	pool, err := pgxpool.Connect(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	if errPing := pool.Ping(ctx); errPing != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	storage := user.NewStorage(pool)
	service := user2.NewService(storage, hasher.New(hashSecret.PasswordHashKey()))

	pb.RegisterUserV1Server(s, &api.User{
		Service: service,
	})

	log.Println("server listening at:", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalln("failed to server:", err.Error())
	}
}
