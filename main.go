package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/dimzrio/grpc-courses/example/nginx-grpc/model"
)

const (
	protocol = "tcp"
	port     = ":50051"
)

type server struct {
}

// SayHello RPC
func (*server) SayHello(ctx context.Context, req *model.HalloReq) (*model.HalloResp, error) {
	log.Println(req)

	// Set timeout
	if ctx.Err() == context.Canceled {
		log.Println("Request was canceled by timeout")
		return nil, status.Error(codes.Canceled, "Request timeout")
	}

	firstname := req.GetFirstName()
	lastname := req.GetLastName()

	resp := &model.HalloResp{
		Fullname: fmt.Sprintf("Hi! %s %s\n", firstname, lastname),
	}

	return resp, nil
}

// DetailsInfo
func (*server) DetailsInfo(ctx context.Context, req *model.InfoReq) (*model.InfoResp, error) {
	log.Println(req)

	// Set timeout
	if ctx.Err() == context.Canceled {
		log.Println("Request was canceled by timeout")
		return nil, status.Error(codes.Canceled, "Request timeout")
	}

	name := req.GetName()

	resp := &model.InfoResp{
		Details: fmt.Sprintf(`*** Information ***
Name: %s
This is an example.!!!`, name),
	}

	return resp, nil
}

func main() {
	log.Println("*** Nginx gRPC Server *** ")

	listen, err := net.Listen(protocol, port)

	if err != nil {
		log.Fatalf("Failed open port: %v", err)
	}

	if err != nil {
		log.Fatalf("Failed read credentials: %v", err)
	}

	serverRPC := grpc.NewServer()
	model.RegisterPersonalServiceServer(serverRPC, &server{})
	reflection.Register(serverRPC)

	err = serverRPC.Serve(listen)
	if err != nil {
		log.Fatalf("Failed server grpc server: %v", err)
	}
}
