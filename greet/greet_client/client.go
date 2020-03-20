package main

import (
	"fmt"
	"log"

	"github.com/nexlight101/grpc-go-course/greet/greet_server/greetpb"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Hello I'm a client")
	// Create a connection
	conn, err := grpc.Dial("lovalhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	// Close connection when done
	defer conn.Close()

	// Create a new client
	c := greetpb.NewGreetServiceClient(conn)
	fmt.Printf("Created client %f", c)

}
