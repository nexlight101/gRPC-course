package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

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
	numb1, numb2, opperator := readArgs()
	doUnary(c, numb1, numb2, opperator)

}

// Populate the protocol buffer request
func doUnary(c calcpb.SumServiceClient, num1 int64, num2 int64, opperator string) {
	fmt.Println("Starting to do Unary RPC...")
	sreq := &calcpb.SumRequest{
		Opperands: &calcpb.Opperands{
			Num1: num1,
			Num2: num2,
		},
		Opperator: &calcpb.Opperator{
			Opperator: opperator,
		},
	}
	res, err := c.Sum(context.Background(), sreq)
	if err != nil {
		log.Fatalf("error while calling calc RPC: %v\n", err)
	}
	log.Printf("%d %s %d = %v", sreq.GetOpperands().GetNum1(), sreq.GetOpperator().GetOpperator(), sreq.GetOpperands().GetNum2(), res.Sum)
}

func readArgs() (num1 int64, num2 int64, opperator string) {
	fmt.Println("Please enter 2 numbers to (add, sub, mult, div)!")
	if len(os.Args) != 4 {
		log.Fatalln("Usage: number1 number2 opperation(add, sub, mult, div) enter")
		return
	}
	num1c, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalln("Cannot convert first number given")
		return
	}
	num2c, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalln("Cannot convert second number given")
		return
	}

	opperator = os.Args[3]
	if num2c == 0 && opperator == "div" {
		log.Fatalln("Cannot devide by zero")
		return
	}
	return int64(num1c), int64(num2c), opperator
}
