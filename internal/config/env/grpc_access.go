package env

import (
	"errors"
	"net"
	"os"

	"github.com/Shemistan/grpc_user_api/internal/config"
)

const (
	grpcAccessHostEnvName = "GRPC_ACCESS_HOST"
	grpcAccessPortEnvName = "GRPC_ACCESS_PORT"
)

type grpcAccessConfig struct {
	host string
	port string
}

// NewGRPCAccessConfig - получить
func NewGRPCAccessConfig() (config.GRPCAccessConfig, error) {
	host := os.Getenv(grpcAccessHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("grpc access host not found")
	}

	port := os.Getenv(grpcAccessPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("grpc access port not found")
	}

	return &grpcAccessConfig{
		host: host,
		port: port,
	}, nil
}

func (cfg *grpcAccessConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
