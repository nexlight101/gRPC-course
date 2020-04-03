package main

import (
	"context"
	"fmt"
	"log"

	"github.com/nexlight101/grpc-go-course/greet/greet_server/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {

	fmt.Println("Hello I'm a client")
	// to use insecure client
	tls := false
	opts := grpc.WithInsecure()
	if tls {
		certFile := "../../ssl/ca.crt" //Certificate Authority Trust certificate
		creds, sslErr := credentials.NewClientTLSFromFile(certFile, "")
		if sslErr != nil {
			log.Fatalf("Error while loading CA trust certificate %v", sslErr)
			return
		}
		opts = grpc.WithTransportCredentials(creds)
	}

	// Create a connection
	conn, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	// Close connection when done
	defer conn.Close()

	// Create a new client
	c := greetpb.NewGreetServiceClient(conn)
	// fmt.Printf("Created client %f", c)
	doUnary(c)

}

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
