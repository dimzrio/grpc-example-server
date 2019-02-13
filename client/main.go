package main

import (
	"context"
	"fmt"
	"log"

	"github.com/dimzrio/grpc-example-server/model"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	log.Println("*** gRCP Client ***")

	cred, err := credentials.NewClientTLSFromFile("../ssl/ca.crt", "")

	if err != nil {
		log.Fatalf("Found Credentials Error : %v\n", err)
	}

	conn, _ := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(cred))

	grpcClient := model.NewPersonalServiceClient(conn)

	req := &model.HalloReq{
		FirstName: "Dimas",
		LastName:  "Rio",
	}
	resp, _ := grpcClient.SayHello(context.Background(), req)

	fmt.Println(resp)

}
