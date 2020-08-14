package client

import (
	"awesomeProject/pb"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
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

func (userClient *UserClient) Search(filter *pb.Filter) {
	req := &pb.SearchUserRequest{
		Filter: filter,
	}

	stream, err := userClient.service.SearchUser(context.Background(), req)
	if err != nil {
		log.Fatal("cant search user: ", err)
	}
	for {
		log.Print("co 1 thang")
		res, err := stream.Recv()
		if err == io.EOF {
			return
		}

		if err != nil {
			log.Fatal("cant receive response: ", err)
		}

		user := res.GetUser()
		log.Println("- found:", user.GetId())
		log.Println("   +age:", user.GetAge())
		//log.Println("   +rank:", user.GetRank())
		log.Println("   +role:", user.GetRole())
	}
}