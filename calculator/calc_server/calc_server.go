package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/nexlight101/grpc-go-course/calculator/calc_server/calcpb"
	"google.golang.org/grpc"
)

// Create server struct
type server struct{}

func (*server) Sum(ctx context.Context, sreq *calcpb.SumRequest) (*calcpb.SumResponse, error) {
	fmt.Printf("Sum function was invoked... %v\n", sreq)
	num1 := sreq.GetOpperands().GetNum1()
	num2 := sreq.GetOpperands().GetNum2()
	opperation := sreq.GetOpperator().GetOpperator()
	var sum float64
	switch opperation {
	case "add":
		sum = float64(num2 + num1)
	case "sub":
		sum = float64(num1 - num2)
	case "mult":
		sum = float64(num1 * num2)
	case "div":
		sum = float64(num1 / num2)
	default:
		fmt.Println("Supply the correct opperation(add, sub, mult, div)")
		return nil, errors.New("Opperation error")
	}
	res := &calcpb.SumResponse{
		Sum: sum,
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
	// Create a new grpc server.
	s := grpc.NewServer()
	calcpb.RegisterSumServiceServer(s, &server{})
	// Serve the listener.
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
