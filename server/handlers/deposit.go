package handlers

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/midepeter/thrift/db/transactions"
	"github.com/midepeter/thrift/domain/model"
	"github.com/midepeter/thrift/proto/transactionpb"
	"github.com/rs/zerolog"
)

type Transaction struct {
	Db           *transactions.Queries
	log          *zerolog.Logger
	AccountLimit int
	transactionpb.UnimplementedTransactionsServer
}

func (t *Transaction) Deposit(ctx context.Context, req *transactionpb.DepositRequest) (*transactionpb.DepositResponse, error) {
	var depositId string = uuid.NewString()

	t.log.Info().Msgf("The user %v making deposit transaction %s at %s", " ", transactionId, "")
	if req == nil {
		return nil, fmt.Errorf("Unable to process invalid request")
	}

	if req.Amount > float32(t.AccountLimit) {
		return nil, fmt.Errorf("Amount exceeds limit")
	}

	depositRequest = model.Deposit{
		Id:      transactionId,
		UserID:  req.UserId),
		Amount:  float32(req.Amount),
		Balance: balance,
	}

	transaction, err := t.Db.CreateTransaction(ctx, transactions.Transaction{
		TransactionID: depositID,
		UserID: req.UserID,
		CurrencyID: 1,
		TransactionAmount: req.Amount,
		TransactionDate: time.Now(),
	})

	if err != nil {
		return nil, fmt.Errorf("Unable to make deposit transaction", err)
	}

	return transactionpb.DepositResponse{
		Id: transaction,
		DateCreated: time.Now(),
	}, nil
}
