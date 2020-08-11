package main

import (
	"awesomeProject/client"
	"awesomeProject/random"
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
	testCreateUser(userClient)
}

func testCreateUser(userClient *client.UserClient) {
	//create random data
	newUser := random.User()

	userClient.CreateUser(newUser)
	return
}