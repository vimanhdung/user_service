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

func (server *UserServer) SearchUser(req *pb.SearchUserRequest, stream pb.UserService_SearchUserServer) error {
	filter := req.GetFilter()
	log.Printf("receive and request search User with filter: %v", filter)
	err := server.userStore.Search(
			*filter,
		func(user *pb.User) error {
			err := stream.Send(&pb.SearchUserResponse{
				User: user,
			})
			if err != nil {
				log.Fatalf("has error when send stream response: %v", err)
				return err
			}
			return nil
		},
	)

	if err != nil {
		return status.Errorf(codes.Internal, "has error when stream data user when search: %v", err)
	}

	return nil
}