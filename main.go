package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"

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
	"github.com/midepeter/thrift/server/middleware"
)

const (
	port = ":5000"
)

func main() {
	var (
		//certFile string = "./out/localhsot.crt"
		//	keyFile  string = "./out/localhost.key"
		certFile string = ""
		keyFile  string = ""
	)

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	fatal := flag.Bool("fatal", false, "It is used to set the log level")

	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	if *fatal {
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	}

	connString := fmt.Sprintf("postgres://midepeter:password@localhost:5432/userdb?sslmode=disable")

	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatal().Msg("Failed to set up database connection")
	}

	queries := db.New(conn)
	transactionQueries := transactions.New(conn)

	mux := http.NewServeMux()

	userPath, userHandler := userpbconnect.NewUserHandler(&server.Server{
		Db: queries,
	})

	transactionPath, transactionHandler := transactionpbconnect.NewTransactionsHandler(&handlers.Transaction{
		Db:           transactionQueries,
		AccountLimit: 10000,
	})

	log.Info().Msg(fmt.Sprintf("The userPath %v userHandler %v\n", userPath, userHandler))
	//	log.Info().Msg(fmt.Sprintf("The transactionPath %v transactionHandler %v\n", transactionPath, transactionHandler))
	mux.Handle(userPath, userHandler)
	mux.Handle(transactionPath, transactionHandler)

	h2s := &http2.Server{}

	srv := &http.Server{
		Addr:         port,
		ReadTimeout:  5 * time.Minute,
		WriteTimeout: 10 * time.Second,
		Handler:      h2c.NewHandler(middleware.AuthMiddleware(mux), h2s),
		TLSConfig:    &tls.Config{ServerName: "localhost"},
	}

	if certFile != "" && keyFile != "" {
		srv.ListenAndServeTLS(certFile, keyFile)
	} else {
		srv.ListenAndServe()
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
