package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/hossainemruz/hello-grpc/calculator"
	"google.golang.org/grpc"
)

const (
	ServerPort = ":8090"
)

type server struct {
	calculator.UnimplementedCalculatorServer
}

// implement Add service
func (s *server) Add(ctx context.Context, in *calculator.Operands) (*calculator.Summation, error) {
	fmt.Printf("Adding.... A: %d B: %d\n", in.A, in.B)
	return &calculator.Summation{Sum: in.A + in.B}, nil
}

// implement Sum service
func (s *server) Sum(ctx context.Context, in *calculator.Numbers) (*calculator.Summation, error) {
	fmt.Println("Summing a list of numbers.......")
	var sum int32
	for i := range in.Num {
		sum += in.Num[i]
	}
	return &calculator.Summation{Sum: sum}, nil
}

func main() {
	listener, err := net.Listen("tcp", ServerPort)
	if err != nil {
		log.Fatalln(err)
	}

	// create a new gRPC server
	s := grpc.NewServer()
	// register calculator server in this gRPC server
	calculator.RegisterCalculatorServer(s, &server{})
	// start gRPC server
	if err := s.Serve(listener); err != nil {
		log.Fatal("failed to start gRPC server")
	}
}
