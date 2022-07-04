package server

import (
	"context"
	"crypto/md5"
	"errors"
	"log"

	"github.com/midepeter/grpc-servie/db"
	"github.com/midepeter/grpc-service/proto/userpb"
)

var store map[string]string

type Server struct {
	db db.Db
	userpb.UnimplementedUserServer
}

func (s *Server) Register(ctx context.Context, in *userpb.UserRequest) (*userpb.UserResponse, error) {
	//log.Println("=== Register User ===")
	var err error
	if in.Email == "" || in.Password == " " {
		return nil, errors.New("Invalid Email or Password")
	}

	hashPassword := md5.Sum([]byte(in.Password))
	
	stmt := fmt.Sprintln(
		`INSERT INTO users (name, email, password) VALUES (%s, %s, %s)`,
		in.Name, in.Email, hashPassword
	)

	err = s.db.Insert(ctx, stmt )
	if err != nil {
		return nil, errors.New("Unable to insert user details into database") 
	}

	//log.Printf("The username %s password %s email %s", in.Name, in.Password, in.Email)
	return &userpb.UserResponse{}, nil
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
