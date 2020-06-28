package main

import (
	"Golang/courses/grpc/greet/greetpb"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Hello from client!")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer conn.Close()

	client := greetpb.NewGreetServiceClient(conn)
	// fmt.Println("Created client: %f", client)

	doUnary(client)
}

func doUnary(c greetpb.GreetServiceClient) {

	fmt.Println("Starting to do a unary RPC")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Kishan",
			LastName:  "Nigam",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}
