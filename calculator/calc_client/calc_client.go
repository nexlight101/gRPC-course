package main

import (
	"context"
	"fmt"
	"log"

	"github.com/nexlight101/grpc-go-course/calculator/calc_server/calcpb"
	"google.golang.org/grpc"
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
	c := calcpb.NewSumServiceClient(conn)
	// fmt.Printf("Created client %f", c)
	doUnary(c)

}

func doUnary(c calcpb.SumServiceClient) {
	fmt.Println("Starting to do Unary RPC...")
	sreq := &calcpb.SumRequest{
		Opperands: &calcpb.Opperands{
			Num1: 20,
			Num2: 22,
		},
	}
	res, err := c.Sum(context.Background(), sreq)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	log.Printf("Addind %d too %d = %v", sreq.GetOpperands().GetNum1(), sreq.GetOpperands().GetNum2(), res.Sum)
}
