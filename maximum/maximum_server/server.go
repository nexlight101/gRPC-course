package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/nexlight101/grpc-go-course/maximum/maximum_server/maximumpb"
	"google.golang.org/grpc"
)

// Create server struct
type server struct{}

func (*server) MaxEveryone(stream maximumpb.MaxService_MaxEveryoneServer) error {
	fmt.Printf("MaxEveryone function was invoked with a streaming request\n")
	maxHolder := make([]int32, 0, 10)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}
		number := req.GetMaximum().GetNumber()
		maxHolder = append(maxHolder, number)
		result := findMax(maxHolder)
		sendErr := stream.Send(&maximumpb.MaxEveryoneResponse{
			Result: result,
		})
		if sendErr != nil {
			log.Fatalf("Error while sending data to client: %v", err)
			return err
		}

	}

}

// findMax finds the maximum of a []int32.
func findMax(maxHolder []int32) int32 {
	var max int32 = maxHolder[0]
	fmt.Println(maxHolder)
	for _, value := range maxHolder {
		if max < value {
			max = value
		}
	}
	return max
}
func main() {
	fmt.Println("Hello Maximum gRPC")

	//  Create a listener at port 50051.
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v ", err)
	}
	// Create a new grpc server.
	s := grpc.NewServer()
	maximumpb.RegisterMaxServiceServer(s, &server{})
	// Serve the listener.
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
