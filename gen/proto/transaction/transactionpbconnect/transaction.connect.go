// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: proto/transaction/transaction.proto

package transactionpbconnect

import (
	context "context"
	errors "errors"
	http "net/http"
	strings "strings"

	connect_go "github.com/bufbuild/connect-go"
	transactionpb "github.com/midepeter/thrift/gen/proto/transaction"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// TransactionsName is the fully-qualified name of the Transactions service.
	TransactionsName = "transaction.Transactions"
)

// TransactionsClient is a client for the transaction.Transactions service.
type TransactionsClient interface {
	Deposit(context.Context, *connect_go.Request[transactionpb.DepositRequest]) (*connect_go.Response[transactionpb.DepositResponse], error)
	Lock(context.Context, *connect_go.Request[transactionpb.LockRequest]) (*connect_go.Response[transactionpb.LockResponse], error)
	Balance(context.Context, *connect_go.Request[transactionpb.BalanceRequest]) (*connect_go.Response[transactionpb.BalanceResponse], error)
	Withdraw(context.Context, *connect_go.Request[transactionpb.WithdrawalRequest]) (*connect_go.Response[transactionpb.WithdrawalResponse], error)
}

// NewTransactionsClient constructs a client for the transaction.Transactions service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTransactionsClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) TransactionsClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &transactionsClient{
		deposit: connect_go.NewClient[transactionpb.DepositRequest, transactionpb.DepositResponse](
			httpClient,
			baseURL+"/transaction.Transactions/Deposit",
			opts...,
		),
		lock: connect_go.NewClient[transactionpb.LockRequest, transactionpb.LockResponse](
			httpClient,
			baseURL+"/transaction.Transactions/Lock",
			opts...,
		),
		balance: connect_go.NewClient[transactionpb.BalanceRequest, transactionpb.BalanceResponse](
			httpClient,
			baseURL+"/transaction.Transactions/Balance",
			opts...,
		),
		withdraw: connect_go.NewClient[transactionpb.WithdrawalRequest, transactionpb.WithdrawalResponse](
			httpClient,
			baseURL+"/transaction.Transactions/Withdraw",
			opts...,
		),
	}
}

// transactionsClient implements TransactionsClient.
type transactionsClient struct {
	deposit  *connect_go.Client[transactionpb.DepositRequest, transactionpb.DepositResponse]
	lock     *connect_go.Client[transactionpb.LockRequest, transactionpb.LockResponse]
	balance  *connect_go.Client[transactionpb.BalanceRequest, transactionpb.BalanceResponse]
	withdraw *connect_go.Client[transactionpb.WithdrawalRequest, transactionpb.WithdrawalResponse]
}

// Deposit calls transaction.Transactions.Deposit.
func (c *transactionsClient) Deposit(ctx context.Context, req *connect_go.Request[transactionpb.DepositRequest]) (*connect_go.Response[transactionpb.DepositResponse], error) {
	return c.deposit.CallUnary(ctx, req)
}

// Lock calls transaction.Transactions.Lock.
func (c *transactionsClient) Lock(ctx context.Context, req *connect_go.Request[transactionpb.LockRequest]) (*connect_go.Response[transactionpb.LockResponse], error) {
	return c.lock.CallUnary(ctx, req)
}

// Balance calls transaction.Transactions.Balance.
func (c *transactionsClient) Balance(ctx context.Context, req *connect_go.Request[transactionpb.BalanceRequest]) (*connect_go.Response[transactionpb.BalanceResponse], error) {
	return c.balance.CallUnary(ctx, req)
}

// Withdraw calls transaction.Transactions.Withdraw.
func (c *transactionsClient) Withdraw(ctx context.Context, req *connect_go.Request[transactionpb.WithdrawalRequest]) (*connect_go.Response[transactionpb.WithdrawalResponse], error) {
	return c.withdraw.CallUnary(ctx, req)
}

// TransactionsHandler is an implementation of the transaction.Transactions service.
type TransactionsHandler interface {
	Deposit(context.Context, *connect_go.Request[transactionpb.DepositRequest]) (*connect_go.Response[transactionpb.DepositResponse], error)
	Lock(context.Context, *connect_go.Request[transactionpb.LockRequest]) (*connect_go.Response[transactionpb.LockResponse], error)
	Balance(context.Context, *connect_go.Request[transactionpb.BalanceRequest]) (*connect_go.Response[transactionpb.BalanceResponse], error)
	Withdraw(context.Context, *connect_go.Request[transactionpb.WithdrawalRequest]) (*connect_go.Response[transactionpb.WithdrawalResponse], error)
}

// NewTransactionsHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTransactionsHandler(svc TransactionsHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/transaction.Transactions/Deposit", connect_go.NewUnaryHandler(
		"/transaction.Transactions/Deposit",
		svc.Deposit,
		opts...,
	))
	mux.Handle("/transaction.Transactions/Lock", connect_go.NewUnaryHandler(
		"/transaction.Transactions/Lock",
		svc.Lock,
		opts...,
	))
	mux.Handle("/transaction.Transactions/Balance", connect_go.NewUnaryHandler(
		"/transaction.Transactions/Balance",
		svc.Balance,
		opts...,
	))
	mux.Handle("/transaction.Transactions/Withdraw", connect_go.NewUnaryHandler(
		"/transaction.Transactions/Withdraw",
		svc.Withdraw,
		opts...,
	))
	return "/transaction.Transactions/", mux
}

// UnimplementedTransactionsHandler returns CodeUnimplemented from all methods.
type UnimplementedTransactionsHandler struct{}

func (UnimplementedTransactionsHandler) Deposit(context.Context, *connect_go.Request[transactionpb.DepositRequest]) (*connect_go.Response[transactionpb.DepositResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("transaction.Transactions.Deposit is not implemented"))
}

func (UnimplementedTransactionsHandler) Lock(context.Context, *connect_go.Request[transactionpb.LockRequest]) (*connect_go.Response[transactionpb.LockResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("transaction.Transactions.Lock is not implemented"))
}

func (UnimplementedTransactionsHandler) Balance(context.Context, *connect_go.Request[transactionpb.BalanceRequest]) (*connect_go.Response[transactionpb.BalanceResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("transaction.Transactions.Balance is not implemented"))
}

func (UnimplementedTransactionsHandler) Withdraw(context.Context, *connect_go.Request[transactionpb.WithdrawalRequest]) (*connect_go.Response[transactionpb.WithdrawalResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("transaction.Transactions.Withdraw is not implemented"))
}
