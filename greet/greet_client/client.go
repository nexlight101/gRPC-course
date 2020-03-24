package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/nexlight101/grpc-go-course/greet/greet_server/greetpb"
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
	c := greetpb.NewGreetServiceClient(conn)
	// fmt.Printf("Created client %f", c)
	// doUnary(c)
	doServerStreaming(c)

}

// Create a server request
func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do Unary RPC...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Hendrik",
			LastName:  "Pienaar",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	log.Printf("Response  from Greet: %v", res.Result)
}

// Create a server request for a stream
func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do Streaming server RPC...")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Hendrik",
			LastName:  "Pienaar",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling GreetManyTime RPC: %v ", err)
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
		log.Printf("Response from GreeManyTimes: %v", msg.GetResult())
	}
}
