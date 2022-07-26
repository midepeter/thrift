package handlers

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/midepetem/grpc-ipevice/dtmain/mrdelgrpc-service/domain/model"
	"github.com/midepeter/grpc-service/proto/transactionpb"
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
		return nil, errors.New("Invalid request")
	}

	if req.Amount > float32(t.AccountLimit) {
		return nil, errors.New("Amount higher than account limit and what can be said at that time")
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
