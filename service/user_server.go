package service

import (
	"awesomeProject/pb"
	"errors"
	"fmt"
	"sync"

	"github.com/jinzhu/copier"
)

var ErrAlreadyExists = errors.New("record already exists")

type UserStore interface {
	Save(user *pb.User) error
	Find(id string) (*pb.User, error)
}

type InMemoryUserStore struct {
	mutex sync.RWMutex
	data map[string]*pb.User
}

func NewUserStore() *InMemoryUserStore {
	return &InMemoryUserStore{
		data: map[string]*pb.User{},
	}
}

func (store *InMemoryUserStore) Save(user *pb.User) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.data[user.Id] != nil {
		return ErrAlreadyExists
	}

	other, err := deepCopy(user)
	if err != nil {
		return err
	}

	store.data[other.Id] = other
	return nil
}

func deepCopy(user *pb.User) (*pb.User, error) {
	other := &pb.User{}

	err := copier.Copy(other, user)
	if err != nil {
		fmt.Errorf("cant copy user data: %w", err)
	}
	return other, nil
}