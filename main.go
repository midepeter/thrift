package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	log "github.com/rs/zerolog/log"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	postgres "github.com/midepeter/thrift/db"
	transaction "github.com/midepeter/thrift/db/transactions"
	db "github.com/midepeter/thrift/db/userstore"
	"github.com/midepeter/thrift/gen/proto/transaction/transactionpbconnect"
	"github.com/midepeter/thrift/gen/proto/user/userpbconnect"
	"github.com/midepeter/thrift/server/handlers/transactions"
	"github.com/midepeter/thrift/server/middleware"
	"github.com/midepeter/thrift/server/user"
)

func main() {
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		panic(fmt.Errorf("Unable to load environment variables %v", err))
	}

	var (
		DBUser     = os.Getenv("DBUSER")
		DBHost     = os.Getenv("DBHOST")
		DBPassword = os.Getenv("DBPASSWORD")
		DBName     = os.Getenv("DBNAME")
		DBPort     = os.Getenv("DBPORT")

		certFile = os.Getenv("CERTFILE")
		keyFile  = os.Getenv("KEYFILE")
		logLevel = os.Getenv("LOGLEVEL")
	)

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	if logLevel != "" {
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", DBUser, DBPassword, DBHost, DBPort, DBName)
	log.Info().Msgf("The connection string %s", dsn)
	pg := postgres.NewPostgres()
	conn, err := pg.Open(ctx, dsn)
	if err != nil {
		panic(fmt.Errorf("Unable to connect to db: %v", err))
	}

	queries := db.New(conn)
	transactionQueries := transaction.New(conn)

	mux := http.NewServeMux()

	userPath, userHandler := userpbconnect.NewUserHandler(&user.Server{
		Db: queries,
	})

	transactionPath, transactionHandler := transactionpbconnect.NewTransactionsHandler(&transactions.Transaction{
		Db:           transactionQueries,
		AccountLimit: 10000,
	})

	log.Info().Msg(fmt.Sprintf("The userPath %v userHandler %v\n", userPath, userHandler))
	mux.Handle(userPath, userHandler)
	mux.Handle(transactionPath, transactionHandler)

	h2s := &http2.Server{}

	srv := &http.Server{
		Addr:         os.Getenv("PORT"),
		ReadTimeout:  5 * time.Minute,
		WriteTimeout: 10 * time.Second,
		Handler:      h2c.NewHandler(middleware.AuthMiddleware(mux), h2s),
		TLSConfig:    &tls.Config{ServerName: "localhost"},
	}

	go func(certFile, keyFile string) {
		if certFile != "" && keyFile != "" {
			srv.ListenAndServeTLS(certFile, keyFile)
		} else {
			srv.ListenAndServe()
		}
	}(certFile, keyFile)

	var cancel context.CancelFunc
	ctx, cancel = signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	select {
	case <-ctx.Done():
		log.Info().Msg("Shutting down server")
		cancel()
	}
}
