package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/nexlight101/grpc-go-course/greet/greet_server/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

// Create server struct
type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	fmt.Println("Hello gRPC")

	//  Create a listener at port 50051.
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v ", err)
	}

	tls := false
	opts := []grpc.ServerOption{}
	if tls {
		certFile := "../../ssl/server.crt"
		keyFile := "../../ssl/server.pem"
		creds, sslErr := credentials.NewServerTLSFromFile(certFile, keyFile)
		if sslErr != nil {
			log.Fatalf("Failed loading certificates: %v", sslErr)
			return
		}
		opts = append(opts, grpc.Creds(creds))
	}
	// Create a new grpc server.
	s := grpc.NewServer(opts...)
	greetpb.RegisterGreetServiceServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	// Serve the listener.
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
