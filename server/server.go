package main

import (
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/midepeter/grpc-streaming/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":5000"
)

type server struct{}

func (s *server) MakeTransaction(in *pb.MoneyRequest, stream pb.MoneyTransaction_MakeTransactionServe) error {
	log.Println("Got transaction Request....")
	log.Printf("Amount: %v, From:%v, to: %s", in.Amount, in.From, in.To)

	for i := 0; i < 3; i++ {
		time.Sleep(10 * time.Second)
		stream.Send(&pb.TransactionResponse{Confirmation: true, status: "completed",
			description: fmt.Sprintf("Description of step %d", int32(i))})
	}
	log.Printf("Successfully transferred $%v from %v to %v", in.Amount, in.From, in.To)
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Listening error: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMoneyTransactionServer(s, server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Serving error: %v", err)
	}
}
