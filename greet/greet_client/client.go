package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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
	doBiDiStreaming(c)

}

func doBiDiStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Bidi streaming  RPC...")
	// we create a stream by invoking the client
	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v", err)
		return
	}
	requests := []*greetpb.GreetEveryoneRequest{
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Hendrik",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Jhonny",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Cobus",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Danniele",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Michele",
			},
		},
	}
	waitc := make(chan struct{})
	// we send a bunch of messages to the client (go routine)

	go func() {
		//function to send a bunch of messages
		for _, req := range requests {
			fmt.Printf("Sending message: %v \n", req)
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()
	// we receive a bunch of messages from the client (go routine)
	go func() {
		// function to receive a bunch of messages
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving: %v", err)
				break
			}
			fmt.Printf("Received: %v\n", res.GetResult())
		}
		close(waitc)
	}()

	// block until everything is done
	<-waitc
}
