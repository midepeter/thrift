package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/midepeter/thrift/db/transactions"
	"github.com/midepeter/thrift/proto/transactionpb"
)

func (t *Transaction) Withdraw(ctx context.Context, req *transactionpb.WithdrawalRequest) (*transactionpb.WithdrawalResponse, error) {
	var withdrawalId string = uuid.NewString()

	if req == nil {
		return nil, fmt.Errorf("Error: withdraw request is empty")
	}

	balance, err := t.Db.GetBalance(ctx, req.UserId)
	if err != nil {
		return nil, fmt.Errorf("Unable to fetch user's balance", err)
	}

	if req.Amount > balance {
		return nil, fmt.Errorf("Unable to complete withdrawal request: Amount execeed withdrawable balance")
	}

	transaction, err := t.Db.CreateTransaction(ctx, transactions.Transaction{
		TransactionID:     withdrawalId,
		UserID:            req.UserId,
		CurrencyID:        1,
		TransactionAmount: -(req.Amount),
		TransactionDate:   time.Now(),
	})

	if err != nil {
		return nil, fmt.Errorf("Unable to make withdrawal transaction ", err)
	}

	return transactionpb.WithdrawalResponse{
		WithdrawalId:      withdrawalId,
		WithrawalResponse: transaction,
	}, nil
}
