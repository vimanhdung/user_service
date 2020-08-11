package service

import (
	"awesomeProject/pb"
	"context"
	"errors"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type UserServer struct {
	userStore UserStore
}

func NewUserServer(userStore UserStore) *UserServer {
	return &UserServer{
		userStore,
	}
}

func (server *UserServer) CreateUser(
	ctx context.Context,
	req *pb.CreateUserRequest,
) (*pb.CreateUserResponse, error) {
	user := req.GetUser()
	log.Printf("receive a create-user request with id: %s", user.Id)

	if len(user.Id) > 0 {
		//check if id is a UUID
		if _, err := uuid.Parse(user.Id); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "user ID is not a valid UUID: %v", err)
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "cant generate new user ID: %v", err)
		}
		user.Id = id.String()
	}

	err := server.userStore.Save(user)
	if err != nil {
		code := codes.Internal
		if errors.Is(err, ErrAlreadyExists) {
			code = codes.AlreadyExists
		}
		return nil, status.Errorf(code, "cant save user to the store: %v", err)
	}

	log.Printf("save user with id: %v", user.Id)

	res := &pb.CreateUserResponse{
		Id: user.Id,
	}
	return res, nil
}
