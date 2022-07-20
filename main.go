package main

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"os"
	"os/signal"
	"flag"

	"github.com/midepeter/grpc-service/db/userStore"
	"github.com/midepeter/grpc-service/proto/userpb"
	"github.com/midepeter/grpc-service/server"
	"google.golang.org/grpc"
    _ "github.com/lib/pq"
	"github.com/rs/zerolog"
	log "github.com/rs/zerolog/log"
)

const (
	port = ":5000"
)


func main() {
	// var (
	// 	certFile string = "server.crt"
	// 	keyFile  string = "server.key"
	// )

	// srv, err := setUpTLS(certFile, keyFile)
	// if err != nil {
	// 	panic(fmt.Errorf("Failed while setting up tls %v\n", err))
	// }

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	fatal := flag.Bool("fatal", false, "It is used to set the debug level to either fatal or not")
	
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	if *fatal {
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	}

	connString := fmt.Sprintf("postgres://midepeter:password@localhost:5432/userdb?sslmode=disable")

	dbConn, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal().Msg("Failed to set up database connection")
		panic(err)
	}

	srv := grpc.NewServer()

	queries := db.New(dbConn)

	userpb.RegisterUserServer(srv, &server.Server{
		Db: queries,
	})

	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	
	log.Info().Msgf("Server running on port %s", port)
	if err := srv.Serve(lis); err != nil {
		log.Debug().Msg("Server failed abruptly")
		panic(err)
	}

	var cancel context.CancelFunc
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	select {
	case <-ctx.Done():
		log.Info().Msg("Shutting down server")
		cancel()
	}
}

// func setUpTLS(certFile, keyFile string) (*grpc.Server, error) {
// 	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
// 	if err != nil {
// 		return nil, fmt.Errorf("Unable to parse certificates: %v\n", err)

// 	}

// 	options := []grpc.ServerOption{
// 		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
// 	}

// 	srv := grpc.NewServer(options...)
// 	return srv, nil
// }
