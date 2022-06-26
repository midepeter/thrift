package server

import (
	"context"
	"errors"
	"log"

	"github.com/midepeter/grpc-service/db"
	"github.com/midepeter/grpc-service/proto/userpb"
)



type Server struct {
	db db.Db
}

func (s *Server) Register(ctx context.Context, in *userpb.UserRequest) (*userpb.UserResponse, error) {
	log.Println("=== Register User ===")
	stmt := `INSERT INTO users `
	err := s.db.Insert(ctx, stmt)
	if err != nil {
		return nil, errors.New("Failed to insert user detailed to db")
	}

	return nil, nil
}

func (s *Server) SignIn(ctx context.Context, in *userpb.UserRequest) (*userpb.UserResponse, error) {
	log.Println("=== LogIn User ===")
	stmt := `SELECT * FROM users`

	err := s.db.Select(ctx, stmt)
	if err != nil {
		return nil, err
	}

	return nil,nil
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
