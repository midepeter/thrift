package server

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/bufbuild/connect-go"
	db "github.com/midepeter/thrift/db/userstore"
	userpb "github.com/midepeter/thrift/gen/proto/user"
	"github.com/midepeter/thrift/pkg/jwt"
	"github.com/rs/zerolog"
	"golang.org/x/crypto/bcrypt"
)

type Server struct {
	Db  *db.Queries
	log *zerolog.Logger
}

func (s *Server) Register(ctx context.Context, in *connect.Request[userpb.RegisterUser]) (*connect.Response[userpb.UserResponse], error) {
	var err error
	if in.Msg.GetEmail() == "" || in.Msg.GetPassword() == " " {
		return nil, fmt.Errorf("Empty Email or password")
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(in.Msg.Password), 10)
	if err != nil {
		return nil, err
	}

	user, err := s.Db.CreateUser(ctx, db.CreateUserParams{
		Email:       in.Msg.Email,
		FirstName:   in.Msg.FirstName,
		LastName:    in.Msg.LastName,
		PhoneNumber: in.Msg.GetPhoneNumber(),
		Password:    string(hashPassword),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})

	if err != nil {
		return nil, fmt.Errorf("Error while creating the user %v", err)
	}

	//s.log.Info().Msgf("The user %s registered successfully", user.Email)

	res := connect.NewResponse(&userpb.UserResponse{
		UserID:     strconv.Itoa(int(user.ID)),
		StatusCode: true,
	})

	return res, nil
}

func (s *Server) SignIn(ctx context.Context, in *connect.Request[userpb.UserRequest]) (*connect.Response[userpb.SignInResponse], error) {
	if in.Msg.Email == "" && in.Msg.Password == "" {
		return nil, fmt.Errorf("Empty Email or Password")
	}

	token, err := jwt.GenerateJwtToken(in.Msg.Email)
	if err != nil {
		return nil, fmt.Errorf("Unable to generate jwt token %v", err)
	}

	res := connect.NewResponse(&userpb.SignInResponse{
		Token:      token,
		StatusCode: true,
	})
	return res, nil
}

func (s *Server) SignOut(ctx context.Context, in *connect.Request[userpb.UserRequest]) (*connect.Response[userpb.UserResponse], error) {
	//Since jwt token are stateless-- we cannot forcefully expire token so the best form of
	//of signing out is for the token to be deleted in the client side
	//s.log.Info().Msgf("%s has loggged out successfully", in.Email)
	log.Println("The user %v has logged out successfully", in.Msg.Email)
	res := connect.NewResponse(&userpb.UserResponse{
		UserID:     in.Msg.Email,
		StatusCode: true,
	})
	return res, nil
}
