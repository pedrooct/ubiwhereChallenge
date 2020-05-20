package main

import (
	"context"
	"log"
	"net"

	pb "github.com/pedrooct/ubiwhereChallenge/gRPC/pb"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedSimulateServiceServer
}


func (s *server) Simulate(ctx context.Context, in *pb.SimulateDataRequest) (*pb.SimulateDataResponse, error) {
	log.Printf("Received: %v,%v,%v,%v", in.GetD1(),in.GetD2(),in.GetD3(),in.GetD4())
	return &pb.SimulateDataResponse{Ok: "Ok"}, nil
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