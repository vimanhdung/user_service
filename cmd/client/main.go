package main

import (
	"awesomeProject/client"
	"awesomeProject/pb"
	"awesomeProject/random"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	//serverAddress := flag.String("address", "", "the server address")
	//log.Printf("dial serer %s", *serverAddress)

	cc, err := grpc.Dial("0.0.0.0:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("cant dial server: ", err)
	}

	userClient := client.NewUserClient(cc)
	//use service
	fmt.Println("Action what you want to do?")
	fmt.Println("1: Test Create User")
	fmt.Println("2: Search User")
	fmt.Println("--------------------")
	var action int
	_, err = fmt.Scanf("%d", &action)
	if err != nil {
		inputNotValid(err)
	}
	switch action {
	case 1:
		testCreateUser(userClient)
	case 2:
		number := scanInputInt("Input number user generate random")
		generateRandomUser(number, userClient)
		min_age := scanInputInt("Input min age")
		max_age := scanInputInt("Input max age")
		gender := scanInputInt("Input gender")
		role := scanInputInt("Input role")
		rank := scanInputInt("Input rank")
		filter := pb.Filter{
			MinAge:   int32(min_age),
			MaxAge:   int32(max_age),
			Gender:   int32(gender),
			Role:     int32(role),
			Rank:     float32(rank),
		}
		userClient.Search(&filter)
	default:
		return
	}
}

func testCreateUser(userClient *client.UserClient) {
	//create random data
	newUser := random.User()

	userClient.CreateUser(newUser)
	return
}

func generateRandomUser(number int, userClient *client.UserClient) error {
	for i := 0; i < number; i++ {
		testCreateUser(userClient)
	}
	return nil
}

func inputNotValid(err error)  {
	log.Fatalf("Input not valid")
}

func scanInputInt(message string) int {
	var i int
	fmt.Println(message)
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		log.Fatalf("cant scan input: %w", err)
	}
	return i
}