package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/nexlight101/grpc-go-course/prime/prime_server/primepb"
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
	c := primepb.NewPrimeServiceClient(conn)

	doServerStreaming(c)

}

// Create a server request for a stream
func doServerStreaming(c primepb.PrimeServiceClient) {
	fmt.Println("Starting to do Streaming server RPC...")

	req := &primepb.PrimeManyTimesRequest{
		Prime: &primepb.Prime{
			Number: 210,
		},
	}
	resStream, err := c.PrimeManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling PrimeManyTime RPC: %v ", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// we've reached the end of the stream
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}
		log.Printf("Response from GreeManyTimes: %d", msg.GetPrimeFactor())
	}
}
