package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	"google.golang.org/grpc"
	pb "github.com/pedrooct/ubiwhereChallenge/tree/master/gRPC/pb"
)

const (
	address     = "localhost:50051"
	defaultName = "ubiwhere"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewsimulateServiceClient(conn)

	// Contact the server and print out its response.
	var arr[4]int
	for i:=0 ; i< 4;i++{
		arr[i]= rand.Intn(99)
		//fmt.Printf("%d|",arr[i])
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.simulate(ctx, &pb.simulateDataRequest{d1: arr[0],d2: arr[1],d3: arr[2],d4: arr[3]})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}