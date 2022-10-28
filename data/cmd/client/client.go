package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/jucsantana/fc2-grp/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect gRPC server %v", err)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection)

	//AddUser(client)
	// AddUserVerbose(client)
	// AddUsers(client)
	AddUserStreamBoth(client)

}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Julio",
		Email: "jucsantana@gmail.com",
	}

	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request %v", err)
	}

	fmt.Println(res)
}

func AddUserVerbose(client pb.UserServiceClient) {

	req := &pb.User{
		Id:    "0",
		Name:  "Julio",
		Email: "jucsantana@gmail.com",
	}

	responseStream, err := client.AddUserVerbose(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request %v", err)
	}

	for {
		stream, err := responseStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Could not receiver stream of mensage %v", err)
		}

		fmt.Println("Status:", stream.Status)
		if stream.Status == "Init" || stream.Status == "Completed" {
			fmt.Println(stream.User)
		}

	}
}

func AddUsers(client pb.UserServiceClient) {

	reqs := []*pb.User{

		{
			Id:    "w1",
			Name:  "Wesley 1",
			Email: "wes1@wes.com",
		},
		{
			Id:    "w2",
			Name:  "Wesley 2",
			Email: "wes2@wes.com",
		},
		{
			Id:    "w3",
			Name:  "Wesley 3",
			Email: "wes3@wes.com",
		},
		{
			Id:    "w4",
			Name:  "Wesley 4",
			Email: "wes4@wes.com",
		},
		{
			Id:    "w5",
			Name:  "Wesley 5",
			Email: "wes5@wes.com",
		},
	}

	stream, err := client.AddUsers(context.Background())

	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	fmt.Println(res)
}

func AddUserStreamBoth(client pb.UserServiceClient) {

	stream, err := client.AddUserStreamBoth(context.Background())

	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	reqs := []*pb.User{

		{
			Id:    "w1",
			Name:  "Wesley 1",
			Email: "wes1@wes.com",
		},
		{
			Id:    "w2",
			Name:  "Wesley 2",
			Email: "wes2@wes.com",
		},
		{
			Id:    "w3",
			Name:  "Wesley 3",
			Email: "wes3@wes.com",
		},
		{
			Id:    "w4",
			Name:  "Wesley 4",
			Email: "wes4@wes.com",
		},
		{
			Id:    "w5",
			Name:  "Wesley 5",
			Email: "wes5@wes.com",
		},
	}

	wait := make(chan int)

	go func() {
		for _, req := range reqs {
			fmt.Println("Sending user: ", req.Name)
			stream.Send(req)
			time.Sleep(time.Second * 2)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error receiving stream from the server: %v", err)
				break
			}

			fmt.Printf("Receving user %v with status: %v \n", res.GetUser().GetName(), res.GetStatus())
		}
		close(wait)
	}()

	<-wait

}
