package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/wrapperspb"

	pb "github.com/Shemistan/grpc_user_api/pkg/user_api_v1"
)

const (
	serviceAddress = "localhost:50051"
	userID         = 1
)

func main() {
	conn, err := grpc.Dial(serviceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("failed to connect to service:", err.Error())
	}
	defer conn.Close() //nolint:errcheck

	c := pb.NewUserV1Client(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	runClientMethods(ctx, c)
}

func runClientMethods(ctx context.Context, c pb.UserV1Client) {
	resGet, err := c.Get(ctx, &pb.GetRequest{Id: userID})
	if err != nil {
		log.Println("failed to get user_v1: ", err.Error())
	}
	log.Printf("get response: %+v\n", resGet)

	_, err = c.Delete(ctx, &pb.DeleteRequest{Id: userID})
	if err != nil {
		log.Println("failed to delete user_v1: ", err.Error())
	}

	resCreate, err := c.Create(ctx, &pb.CreateRequest{
		Name:            "test_name",
		Email:           "test_email",
		Password:        "test_password",
		PasswordConfirm: "test_password",
		Role:            1,
	})
	if err != nil {
		log.Println("failed to Create user_v1: ", err.Error())
	}
	log.Printf("create response: %+v\n", resCreate)

	_, err = c.Update(ctx, &pb.UpdateRequest{
		Id:    userID,
		Name:  nil,
		Email: &wrapperspb.StringValue{Value: "test_email"},
		Role:  0,
	})
	if err != nil {
		log.Println("failed to Update user_v1: ", err.Error())
	}
}
