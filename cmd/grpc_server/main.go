package main

import (
	"context"
	"flag"

	"log"
	"net"

	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/Shemistan/grpc_user_api/internal/api"
	"github.com/Shemistan/grpc_user_api/internal/config"
	"github.com/Shemistan/grpc_user_api/internal/config/env"
	pb "github.com/Shemistan/grpc_user_api/pkg/user_api_v1"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

func main() {
	flag.Parse()
	ctx := context.Background()

	// Считываем переменные окружения
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

	secrConf, err := env.NewTesConfig()
	if err != nil {
		log.Fatalf("failed to get pg config: %v", err)
	}

	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatalln("failed to listen: ", err.Error())
	}

	// Создаем пул соединений с базой данных
	pool, err := pgxpool.Connect(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterUserV1Server(s, &api.User{})

	log.Println("server listening at:", lis.Addr(), "secret: ", secrConf)

	if err = s.Serve(lis); err != nil {
		log.Fatalln("failed to server:", err.Error())
	}
}
