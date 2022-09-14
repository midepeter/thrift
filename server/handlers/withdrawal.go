package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"math/big"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"

	"github.com/bufbuild/connect-go"
	"github.com/midepeter/thrift/db/transactions"
	transactionpb "github.com/midepeter/thrift/gen/proto/transaction"
)

func (t *Transaction) Withdraw(ctx context.Context, req *connect.Request[transactionpb.WithdrawalRequest]) (*connect.Response[transactionpb.WithdrawalResponse], error) {
	var withdrawalId string = uuid.NewString()

	if req == nil {
		return nil, fmt.Errorf("Error: withdraw request is empty")
	}

	balance, err := t.Db.GetBalance(ctx, req.Msg.UserId)
	if err != nil {
		return nil, fmt.Errorf("Unable to fetch user's balance %v", err)
	}

	if req.Msg.Amount > float32(balance.BalanceAmount) {
		return nil, fmt.Errorf("Unable to complete withdrawal request: Amount execeed withdrawable balance")
	}

	transaction, err := t.Db.CreateTransaction(ctx, transactions.CreateTransactionParams{
		TransactionID: withdrawalId,
		UserID:        req.Msg.UserId,
		CurrencyID:    1,
		TransactionAmount: pgtype.Numeric{
			Int:    big.NewInt(int64(-(req.Msg.Amount))),
			Status: pgtype.Present,
		},
		TransactionDate: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	})

	if err != nil {
		return nil, fmt.Errorf("Unable to make withdrawal transaction %v", err)
	}

	res := connect.NewResponse(&transactionpb.WithdrawalResponse{
		WithdrawalId: transaction.TransactionID,
	})
	return res, nil
}
