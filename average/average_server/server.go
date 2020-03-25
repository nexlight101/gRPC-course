package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/nexlight101/grpc-go-course/average/average_server/averagepb"
	"google.golang.org/grpc"
)

// Create server struct
type server struct{}

// LongGreet receives a stream of request from client and returns a  response once stream completes.
func (*server) LongGreet(stream averagepb.AverageService_LongGreetServer) error {
	fmt.Printf("GreetManyTimes function was invoked with a streaming request\n")
	count := 0
	var avg int
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// we have finished reading the client stream
			return stream.SendAndClose(&averagepb.LongAverageResponse{
				Result: float64(avg / count), // Average to return to client
			})
		}
		if err != nil {
			log.Fatalf("Cannot read from client stream: %v", err)
			return err
		}
		number := req.GetAverage().GetNumber()
		avg += int(number)
		count++
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
	averagepb.RegisterAverageServiceServer(s, &server{})
	// Serve the listener.
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
