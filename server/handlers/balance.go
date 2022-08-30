package handlers

import (
	"context"
	"fmt"

	"github.com/bufbuild/connect-go"
	transactionpb "github.com/midepeter/thrift/gen/proto/transaction"
)

func (t *Transaction) Balance(ctx context.Context, req *connect.Request[transactionpb.BalanceRequest]) (*connect.Response[transactionpb.BalanceResponse], error) {
	if req == nil {
		return nil, fmt.Errorf("Invalid balance request")
	}

	balance, err := t.Db.GetBalance(ctx, req.Msg.UserId)
	if err != nil {
		return nil, fmt.Errorf("Unable to return balance from db: %v", err)
	}

	res := connect.NewResponse(&transactionpb.BalanceResponse{
		UserId:        req.Msg.UserId,
		BalanceAmount: float32(balance.BalanceAmount),
	})
	return res, nil
}
