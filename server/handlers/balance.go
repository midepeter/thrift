package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/midepeter/thrift/domain/model"
	"github.com/midepeter/thrift/proto/transactionpb"
)

func (t *Transaction) Balance(ctx context.Context, req *transactionpb.BalanceRequest) (*transactionpb.BalanceRespone, error) {
	if req == nil {
		return nil, fmt.Errorf("Invalid balance request")
	}

	//This is an unncessary code
	requestBalance := model.Balance{
		UserId:     req.UserId,
		UpdateTime: time.Now(),
	}

	balance, err := t.Db.GetBalance(ctx, requestBalance.UserID)
	if err != nil {
		return nil, fmt.Errorf("Unable to return balance from db: %v", err)
	}

	return transactionpb.BalanceResponse{
		UserId:        requestBalance.UserID,
		BalanceAmount: balance,
	}, nil
}
