package main

import (
	"context"
	"fmt"
	"log"

	"github.com/dimzrio/grpc-example-server/model"

	"google.golang.org/grpc"
)

func main() {
	log.Println("*** gRCP Client ***")

	conn, _ := grpc.Dial("localhost:8080", grpc.WithInsecure())

	grpcClient := model.NewPersonalServiceClient(conn)

	req := &model.HalloReq{
		FirstName: "Dimas",
		LastName:  "Rio",
	}
	resp, _ := grpcClient.SayHello(context.Background(), req)

	fmt.Println(resp)

}
