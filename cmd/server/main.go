package main

import (
	"awesomeProject/pb"
	"awesomeProject/service"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	port := flag.Int("port", 0, "the user server port")
	flag.Parse()

	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cant start server: ", err)
	}

	userStore := service.NewUserStore()
	userServer := service.NewUserServer(userStore)
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, userServer)

	log.Printf("Start gRPC user server at %s", address)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cant start server: ", err)
	}
	return
}
