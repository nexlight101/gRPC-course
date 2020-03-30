package main

import (
	"context"
	"fmt"
	"log"

	"github.com/nexlight101/grpc-go-course/calculator/calc_server/calcpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {

	fmt.Println("Hello I'm a client calculator")
	// Create a connection
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	// Close connection when done
	defer conn.Close()

	// Create a new client
	c := calcpb.NewCalculatorServiceClient(conn)
	// fmt.Printf("Created client %f", c)
	// numb1, numb2, opperator := readArgs()
	// doUnary(c, numb1, numb2, opperator)
	doErrorUnary(c)

}

// Populate the protocol buffer request for Squareroot
func doErrorUnary(c calcpb.CalculatorServiceClient) {
	fmt.Println("Starting to do Squareroot Unary RPC...")
	number := int32(10)
	//correct call
	doErrorCall(c, number)

	number = -10
	doErrorCall(c, number)
	//error call
}

func doErrorCall(c calcpb.CalculatorServiceClient, n int32) {
	res, err := c.SquareRoot(context.Background(), &calcpb.SquareRootRequest{Number: n})
	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			// actual error from gRPC (user error)
			fmt.Printf("Error message from server: %v\n", respErr.Message())
			fmt.Println(respErr.Code())
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("We probably sent a negative number!")
				return
			}
		} else {
			log.Fatalf("Big Error calling SquareRoot: %v", err)
			return
		}
	}
	fmt.Printf("Result of squareroot of %d: %.1f\n", n, res.GetNumberRoot())
}
