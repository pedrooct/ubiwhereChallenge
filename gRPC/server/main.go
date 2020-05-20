package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "github.com/pedrooct/ubiwhereChallenge/tree/master/gRPC/pb"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedSimulateServiceServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) simulate(ctx context.Context, in *pb.SimulateDataRequest) (*pb.SimulateDateResponse, error) {
	log.Printf("Received: %v", in.GetD1())
	return &pb.SimulateDataResponse{Message: "Ok"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSimulateServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}