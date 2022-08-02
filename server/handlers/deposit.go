package handlers

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/midepeter/thrift/domain/model"
	"github.com/midepeter/thrift/proto/transactionpb"
	"github.com/rs/zerolog"
)

type Transaction struct {
	log          *zerolog.Logger
	AccountLimit int
	transactionpb.UnimplementedTransactionsServer
}

func (t *Transaction) Deposit(ctx context.Context, req *transactionpb.DepositRequest) (*transactionpb.DepositResponse, error) {
	var transactionId string = uuid.NewString()

	t.log.Info().Msgf("The user %v making deposit transaction %s at %s", " ", transactionId, "")
	if req == nil {
		return nil, fmt.Errorf("Unable to process invalid request")
	}

	if req.Amount > float32(t.AccountLimit) {
		return nil, fmt.Errorf("Amount exceeds limit")
	}

	//intialBalance of the user needs to be fetched
	balance := float32(700.00)

	_ = model.Deposit{
		Id:      transactionId,
		UserID:  "",
		Amount:  float32(req.Amount),
		Balance: balance,
	}

	return nil, nil
}
