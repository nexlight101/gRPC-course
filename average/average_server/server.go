package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/nexlight101/grpc-go-course/greet/greet_server/greetpb"
	"google.golang.org/grpc"
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

// LongGreet receives a stream of request from client and returns a  response once stream completes.
func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	fmt.Printf("GreetManyTimes function was invoked with a streaming request\n")
	result := "Hello "
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// we have finished reading the client stream
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("Cannot read from client stream: %v", err)
			return err
		}
		firstName := req.GetGreeting().GetFirstName()
		result += firstName + "! "
	}

}
func main() {
	fmt.Println("Hello gRPC")

	//  Create a listener at port 50051.
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v ", err)
	}
	// Create a new grpc server.
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})
	// Serve the listener.
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
