package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/Shemistan/grpc_user_api/internal/api"
	pb "github.com/Shemistan/grpc_user_api/pkg/user_api_v1"
)

const grpcPort = 50051

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalln("failed to listen: ", err.Error())
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterUserV1Server(s, &api.User{})

	log.Println("server listening at:", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalln("failed to server:", err.Error())
	}
}
