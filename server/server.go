package server

import (
	"context"
	"log"
	"reflect"

	"github.com/midepeter/grpc-service/proto/userpb"
)

var store map[string]string

type Server struct {
	userpb.UnimplementedUserServer
}

func (s *Server) Register(ctx context.Context, in *userpb.UserRequest) (*userpb.UserResponse, error) {
	log.Println("=== Register User ===")
	log.Printf("The value of the map %v, ", reflect.ValueOf(store))

	log.Printf("The username %s password %s email %s", in.Name, in.Password, in.Email)
	return nil, nil
}

func (s *Server) SignIn(ctx context.Context, in *userpb.UserRequest) (*userpb.UserResponse, error) {
	log.Println("=== LogIn User ===")
	log.Println("the user is being signed in")

	return nil, nil
}

func (s *Server) SignOut(ctx context.Context, in *userpb.UserRequest) (*userpb.UserResponse, error) {
	log.Println("=== SignOut User===")
	log.Println("The user is being signed out")
	return nil, nil
}
