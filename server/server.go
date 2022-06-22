package server

import (
	"context"
	"log"

	"github.com/midepeter/grpc-service/proto/userpb"
)

const (
	port = ":5000"
)

type Server struct{}

func (s *Server) Register(ctx context.Context, in *userpb.UserRequest) (*userpb.UserResponse, error) {
	log.Println("==Register User ===")
	return nil, nil
}


func startUserServer() error {
	return nil
}

// type server struct{}

// func (s *server) MakeTransaction(in *pb.MoneyRequest, stream pb.MoneyTransaction_MakeTransactionServe) error {
// 	log.Println("Got transaction Request....")
// 	log.Printf("Amount: %v, From:%v, to: %s", in.Amount, in.From, in.To)

// 	for i := 0; i < 3; i++ {
// 		time.Sleep(10 * time.Second)
// 		stream.Send(&pb.TransactionResponse{Confirmation: true, status: "completed",
// 			description: fmt.Sprintf("Description of step %d", int32(i))})
// 	}
// 	log.Printf("Successfully transferred $%v from %v to %v", in.Amount, in.From, in.To)
// 	return nil
// }

// func main() {
// 	lis, err := net.Listen("tcp", port)
// 	if err != nil {
// 		log.Fatalf("Listening error: %v", err)
// 	}
// 	s := grpc.NewServer()
// 	pb.RegisterMoneyTransactionServer(s, server{})
// 	reflection.Register(s)
// 	if err := s.Serve(lis); err != nil {
// 		log.Fatalf("Serving error: %v", err)
// 	}
// }
