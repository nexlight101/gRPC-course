package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/nexlight101/grpc-go-course/prime/prime_server/primepb"
	"google.golang.org/grpc"
)

// Create server struct
type server struct{}

func (*server) PrimeManyTimes(req *primepb.PrimeManyTimesRequest, stream primepb.PrimeService_PrimeManyTimesServer) error {
	fmt.Printf("PrimeManyTimes function was invoked with %v\n", req)
	number := req.GetPrime().GetNumber()
	var k int32 = 2
	for {
		if number%k == 0 { // Find a factor
			primeFactor := k
			res := &primepb.PrimeManyTimesResponse{
				PrimeFactor: primeFactor,
			}
			stream.Send(res)
			time.Sleep(1000 * time.Millisecond)
			number = number / k
			if number <= 1 {
				break
			}
		} else {
			k++
		}
	}
	return nil
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
	primepb.RegisterPrimeServiceServer(s, &server{})
	// Serve the listener.
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
