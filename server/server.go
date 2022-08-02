package server

import (
	"context"
	"fmt"
	"strconv"
	"time"

	db "github.com/midepeter/thrift/db/userstore"
	"github.com/midepeter/thrift/proto/userpb"
	"github.com/midepeter/thrift/utils"
	"github.com/rs/zerolog"
	"golang.org/x/crypto/bcrypt"
)

type Server struct {
	Db  *db.Queries
	log *zerolog.Logger
	userpb.UnimplementedUserServer
}

func (s *Server) Register(ctx context.Context, in *userpb.RegisterUser) (*userpb.UserResponse, error) {
	var err error
	if in.Email == "" || in.Password == " " {
		return nil, fmt.Errorf("Empty Email or password")
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), 10)
	if err != nil {
		return nil, err
	}

	user, err := s.Db.CreateUser(ctx, db.CreateUserParams{
		ID:          1,
		Email:       in.Email,
		FirstName:   in.FirstName,
		LastName:    in.LastName,
		PhoneNumber: in.PhoneNumber,
		Password:    string(hashPassword),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})

	if err != nil {
		return nil, fmt.Errorf("Error while creating the user %v", err)
	}

	s.log.Info().Msgf("The user %s registered successfully", user.Email)

	return &userpb.UserResponse{
		UserID:     strconv.Itoa(int(user.ID)),
		StatusCode: true,
	}, nil
}

func (s *Server) SignIn(ctx context.Context, in *userpb.UserRequest) (*userpb.UserResponse, error) {
	if in.Email == "" && in.Password == "" {
		return nil, fmt.Errorf("Empty Email or Password")
	}

	// switch in {
	// case in.Email:
	// 	user, err := s.Db.GetUser(ctx, strconv.Itoa(in.UserID))
	// 	if err != nil {
	// 		return nil, errors.New("Invalid credentials")
	// 	}
	// default:
	// 	return nil, errors.New("Unable to retrieve user details")
	// }

	token, err := utils.GenerateJwtToken(in.Email)
	if err != nil {
		return nil, fmt.Errorf("Unable to generate jwt token %v", err)
	}

	return &userpb.UserResponse{
		UserID:     token,
		StatusCode: true,
	}, nil
}

func (s *Server) SignOut(ctx context.Context, in *userpb.UserRequest) (*userpb.UserResponse, error) {
	//Since jwt token are stateless-- we cannot forcefully expire token so the best form of
	//of signing out is for the token to be deleted in the client side
	s.log.Info().Msgf("%s has loggged out successfully", in.Email)
	return nil, nil
}
