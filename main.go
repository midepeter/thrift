package main

import (
	"context"
	"crypto/tls"
	"flag"
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
	"github.com/midepeter/thrift/db/transactions"
	db "github.com/midepeter/thrift/db/userstore"
	"github.com/midepeter/thrift/gen/proto/transaction/transactionpbconnect"
	"github.com/midepeter/thrift/gen/proto/user/userpbconnect"
	"github.com/midepeter/thrift/server"
	"github.com/midepeter/thrift/server/handlers"
	"github.com/midepeter/thrift/server/middleware"
)

func main() {
	ctx := context.Background()
	var (
		//certFile string = "./out/localhsot.crt"
		//keyFile  string = "./out/localhost.key"
		certFile string = ""
		keyFile  string = ""
	)

	err := godotenv.Load()
	if err != nil {
		panic(fmt.Errorf("Unable to load environment variables %v", err))
	}

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	fatal := flag.Bool("fatal", false, "It is used to set the log level")

	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	if *fatal {
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	}

	DBUser := os.Getenv("DBUSER")
	DBHost := os.Getenv("DBHOST")
	DBPassword := os.Getenv("DBPASSWORD")
	DBName := os.Getenv("DBNAME")
	DBPort := os.Getenv("DBPORT")

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", DBUser, DBPassword, DBHost, DBPort, DBName)
	log.Info().Msgf("The connection string %s", connString)
	pg := postgres.NewPostgres()
	conn, err := pg.Open(ctx, connString)
	if err != nil {
		panic(fmt.Errorf("Unable to connect to db: %v", err))
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

	if certFile != "" && keyFile != "" {
		srv.ListenAndServeTLS(certFile, keyFile)
	} else {
		srv.ListenAndServe()
	}

	var cancel context.CancelFunc
	ctx, cancel = signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	select {
	case <-ctx.Done():
		log.Info().Msg("Shutting down server")
		cancel()
	}
}
