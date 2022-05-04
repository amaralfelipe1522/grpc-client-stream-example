package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/amaralfelipe1522/grpc-client-stream-example/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)
	}
	defer connection.Close()

	client := pb.NewUserServiceClient(connection)
	//AddUser(client)
	//AddUserVerbose(client)
	AddUsers(client)
}

// Request

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Felipe Amaral de Souza",
		Email: "amaral.felipe@protonmail.com",
	}

	resStream, err := client.AddUserVerbose(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	for {
		stream, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Could not receive the message: %v", err)
		}
		fmt.Println("Status: ", stream.Status)
	}
}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		&pb.User{
			Id:    "0",
			Name:  "Felipe Amaral de Souza",
			Email: "amaral.felipe@protonmail.com",
		},
		&pb.User{
			Id:    "1",
			Name:  "Gustavo Henrique Freitas",
			Email: "gustavo.souza@protonmail.com",
		},
		&pb.User{
			Id:    "2",
			Name:  "Leandro Machado Siqueira",
			Email: "leandro.siqueira@protonmail.com",
		},
		&pb.User{
			Id:    "3",
			Name:  "Kaique Spagnol Tofoli",
			Email: "kaique.tofoli@protonmail.com",
		},
	}

	stream, err := client.AddUsers(context.Background()) // Garante que se a mensagem não for chegar, a requisição será encerrada
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiiving response: %v", err)
	}
	fmt.Println(res)
}
