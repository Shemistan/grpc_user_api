package env

import (
	"errors"
	"net"
	"os"

	"github.com/Shemistan/grpc_user_api/internal/config"
)

const (
	grpcAuthHostEnvName = "GRPC_AUTH_HOST"
	grpcAuthPortEnvName = "GRPC_AUTH_PORT"
)

type grpcAuthConfig struct {
	host string
	port string
}

// NewGRPCAuthConfig - получить
func NewGRPCAuthConfig() (config.GRPCAuthConfig, error) {
	host := os.Getenv(grpcAuthHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("grpc auth host not found")
	}

	port := os.Getenv(grpcAuthPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("grpc auth port not found")
	}

	return &grpcAuthConfig{
		host: host,
		port: port,
	}, nil
}

func (cfg *grpcAuthConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
