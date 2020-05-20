package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"google.golang.org/grpc"
	pb "github.com/pedrooct/ubiwhereChallenge/gRPC/pb"
)

const (
	address     = "localhost:50051"
)

func main() {
	var arr[4]int64
	for i:=0 ; i< 4;i++{
		arr[i]= rand.Int63n(99)
		fmt.Printf("%d|",arr[i])
	}
	fmt.Printf("\n")
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSimulateServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()
	// Send DATA
	r, err := c.Simulate(ctx, &pb.SimulateDataRequest{D1: int64(arr[0]),D2: int64(arr[1]),D3: int64(arr[2]),D4: int64(arr[3])})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetOk())
}