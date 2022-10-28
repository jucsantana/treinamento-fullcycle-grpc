package main

import (
	"net"
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"github.com/jucsantana/fc2-grp/pb"
	"github.com/jucsantana/fc2-grp/services"
)

func main(){

	lis, err := net.Listen("tcp", "localhost:50051")

	if err != nil{
		log.Fatalf("Could not connect: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, services.NewUserService())
	reflection.Register(grpcServer)
	

	if err := grpcServer.Serve(lis); err != nil{
		log.Fatalf("Could not serve: %v", err)
	}

}