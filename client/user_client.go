package client

import (
	"awesomeProject/pb"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type UserClient struct {
	service pb.UserServiceClient
}

func NewUserClient(cc *grpc.ClientConn) *UserClient {
	service := pb.NewUserServiceClient(cc)
	return &UserClient{
		service,
	}
}

func (userClient *UserClient) CreateUser(user *pb.User) {
	req := &pb.CreateUserRequest{
		User: user,
	}

	res, err := userClient.service.CreateUser(context.Background(), req)
	if err != nil {
		stt, ok := status.FromError(err)
		if ok && stt.Code() == codes.AlreadyExists {
			log.Printf("user already exists")
		} else {
			log.Printf("cant create user: %v", err)
		}
		return
	}
	log.Printf("created user with ID: %s", res.Id)
}