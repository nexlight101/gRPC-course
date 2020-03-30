package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math"
	"net"

	"github.com/nexlight101/grpc-go-course/calculator/calc_server/calcpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (*server) SquareRoot(ctx context.Context, req *calcpb.SquareRootRequest) (*calcpb.SquareRootReponse, error) {
	fmt.Printf("Squareroot RPC invoked...")
	number := req.GetNumber()
	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number: %v", number),
		)
	}
	return &calcpb.SquareRootReponse{
		NumberRoot: math.Sqrt(float64(number)),
	}, nil
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
	calcpb.RegisterCalculatorServiceServer(s, &server{})
	// Serve the listener.
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
