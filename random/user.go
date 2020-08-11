package random

import (
	"awesomeProject/pb"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"math/rand"
)

func User() *pb.User {
	birthday, _ := ptypes.TimestampProto(randomBirthday())
	user := &pb.User{
		UserName:   fmt.Sprint("user_", rand.Int()),
		FirstName:  "Vi",
		MiddleName: "Manh",
		LastName:   "Dung",
		Age:        int32(randomInt(18, 80)),
		Gender:     randomGender(),
		Birthday:   birthday,
		Role:       int32(randomRole()),
		Weight:     &pb.User_WeightKg{WeightKg: randomFloat(30, 120)},
		High:       float32(randomFloat(1.1, 2.2)),
		Position:   "developer",
		Rank:       randomRank(),
	}

	return user
}
