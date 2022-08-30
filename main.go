package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"

	"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	log "github.com/rs/zerolog/log"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"net/http"
	_ "net/http/pprof"

	"github.com/midepeter/thrift/db/transactions"
	db "github.com/midepeter/thrift/db/userstore"
	"github.com/midepeter/thrift/gen/proto/transaction/transactionpbconnect"
	"github.com/midepeter/thrift/gen/proto/user/userpbconnect"
	"github.com/midepeter/thrift/server"
	"github.com/midepeter/thrift/server/handlers"
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

	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatal().Msg("Failed to set up database connection")
	}

	//srv := grpc.NewServer()

	queries := db.New(conn)
	transactionQueries := transactions.New(conn)

	mux := http.NewServeMux()
	///transactionpb.RegisterTransactionsServer(srv, &handlers.Transaction{
	//	Db:           transactionQueries,
	//	AccountLimit: 10000,
	//})
	//
	userPath, userHandler := userpbconnect.NewUserHandler(&server.Server{
		Db: queries,
	})

	transactionPath, transactionHandler := transactionpbconnect.NewTransactionsHandler(&handlers.Transaction{
		Db:           transactionQueries,
		AccountLimit: 10000,
	})

	log.Info().Msg(fmt.Sprintf("The userPath %v userHandler %v\n", userPath, userHandler))
	log.Info().Msg(fmt.Sprintf("The transactionPath %v transactionHandler %v\n", transactionPath, transactionHandler))
	mux.Handle(userPath, userHandler)
	mux.Handle(transactionPath, transactionHandler)

	http.ListenAndServe("localhost:5000", h2c.NewHandler(mux, &http2.Server{}))
	//lis, err := net.Listen("tcp", port)
	//if err != nil {
	//	panic(err)
	//}

	//log.Info().Msgf("Server running on port %s", port)
	//if err := srv.Serve(lis); err != nil {
	//	log.Debug().Msg("Server failed abruptly")
	//	panic(err)
	//}

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
