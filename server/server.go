package server

import (
	"context"
	"errors"
	"strconv"
	"time"

	db "github.com/midepeter/grpc-service/db/userStore"
	"github.com/midepeter/grpc-service/proto/userpb"
	"github.com/midepeter/grpc-service/utils"
	"github.com/rs/zerolog"
	"golang.org/x/crypto/bcrypt"
)

type Server struct {
	Db  *db.Queries
	log *zerolog.Logger
	userpb.UnimplementedUserServer
}

func (s *Server) Register(ctx context.Context, in *userpb.UserRequest) (*userpb.UserResponse, error) {
	var err error
	if in.Email == "" || in.Password == " " {
		return nil, errors.New("Invalid Email or Password")
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), 10)
	if err != nil {
		return nil, err
	}

	dateCreated := time.Now().String()

	user, err := s.Db.CreateUser(ctx, db.CreateUserParams{
		ID:          1,
		Email:       in.Email,
		Password:    string(hashPassword),
		DateCreated: dateCreated,
	})

	if err != nil {
		return nil, errors.New("Unable to insert user details into database")
	}

	s.log.Info().Msgf("The user %s registered successfully", user.Email)

	return &userpb.UserResponse{
		UserID:     strconv.Itoa(int(user.ID)),
		StatusCode: true,
	}, nil
}

func (s *Server) SignIn(ctx context.Context, in *userpb.UserRequest) (*userpb.UserResponse, error) {
	if in.Email == "" && in.Password == "" {
		return nil, errors.New("Invalid Email or Password")
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
		return nil, errors.New("Unable to generate jwt token")
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
