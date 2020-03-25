package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/nexlight101/grpc-go-course/average/average_server/averagepb"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Hello I'm a client")
	// Create a connection
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	// Close connection when done
	defer conn.Close()

	// Create a new client
	c := averagepb.NewAverageServiceClient(conn)
	// fmt.Printf("Created client %f", c)
	// doUnary(c)
	doClientStreaming(c)

}

func doClientStreaming(c averagepb.AverageServiceClient) {
	fmt.Println("Starting to do Client Streaming RPC...")

	requests := []*averagepb.LongAverageRequest{
		&averagepb.LongAverageRequest{
			Average: &averagepb.Average{
				Number: 1,
			},
		},
		&averagepb.LongAverageRequest{
			Average: &averagepb.Average{
				Number: 2,
			},
		},
		&averagepb.LongAverageRequest{
			Average: &averagepb.Average{
				Number: 3,
			},
		},
		&averagepb.LongAverageRequest{
			Average: &averagepb.Average{
				Number: 4,
			},
		},
	}
	stream, err := c.LongAverage(context.Background())
	if err != nil {
		log.Fatalf("error while calling LongAverage: %v", err)
	}
	// we uterate over our slice and send each message individually
	for _, req := range requests {
		fmt.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(100 * time.Millisecond)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from LongAverage: %v", err)
	}
	fmt.Printf("LongAverage Response: %.1f\n", res.Result)
}
