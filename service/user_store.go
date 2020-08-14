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
	Search(filter pb.Filter, found func(user *pb.User) error) error
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
	//mutex lock and unlock
	store.mutex.Lock()
	defer store.mutex.Unlock()

	//check if data is already in memory
	if store.data[user.Id] != nil {
		return ErrAlreadyExists
	}

	//copy data
	other, err := deepCopy(user)
	if err != nil {
		return err
	}

	//store user
	store.data[other.Id] = other
	return nil
}

func (store *InMemoryUserStore) Find(id string) (*pb.User, error) {
	//mutex
	store.mutex.Lock()
	defer store.mutex.Unlock()

	user := store.data[id]
	if user == nil {
		return nil, nil
	}
	return deepCopy(user)
}

func deepCopy(user *pb.User) (*pb.User, error) {
	other := &pb.User{}

	err := copier.Copy(other, user)
	if err != nil {
		return nil, fmt.Errorf("cant copy user data: %w", err)
	}
	return other, nil
}

func (store *InMemoryUserStore) Search(filter pb.Filter, found func(user *pb.User) error) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	//process search
	for _, user := range store.data{
		ok := isQualify(filter, *user)
		if ok {
			//deep copy value
			other, err := deepCopy(user)
			if err != nil {
				return err
			}
			found(other)
		}
	}

	return nil
}

func isQualify(filter pb.Filter, user pb.User) bool {
	//if filter.Rank < user.Rank.Rank { return false}
	//if filter.Gender != int32(user.Gender.Number()) { return false }
	if filter.MinAge > user.Age || filter.MaxAge < user.Age { return false }
	return true
}