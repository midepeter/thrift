package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/midepeter/thrift/db/transactions"
	transactionpb "github.com/midepeter/thrift/gen/proto/transaction"
	"github.com/rs/zerolog"
)

type Transaction struct {
	Db           *transactions.Queries
	log          *zerolog.Logger
	AccountLimit int
}

func (t *Transaction) Deposit(ctx context.Context, req *connect.Request[transactionpb.DepositRequest]) (*connect.Response[transactionpb.DepositResponse], error) {
	var depositId string = uuid.NewString()

	log.Printf("The user %v making deposit transaction %s at %v", req.Msg.UserId, depositId, time.Now())
	if req == nil {
		return nil, fmt.Errorf("Unable to process invalid request")
	}

	if req.Msg.Amount > float32(t.AccountLimit) {
		return nil, fmt.Errorf("Amount exceeds limit")
	}

	_, err := t.Db.CreateTransaction(ctx, transactions.CreateTransactionParams{
		TransactionID: depositId,
		UserID:        req.Msg.UserId,
		CurrencyID:    1,
		TransactionAmount: pgtype.Numeric{
			Int:    big.NewInt(int64(req.Msg.Amount)),
			Status: pgtype.Present,
		},
		TransactionDate: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	})

	if err != nil {
		return nil, fmt.Errorf("Unable to make deposit transaction %#v", err)
	}

	res := connect.NewResponse(&transactionpb.DepositResponse{
		Id:            depositId,
		DepositStatus: 2,
	})
	return res, nil
}
