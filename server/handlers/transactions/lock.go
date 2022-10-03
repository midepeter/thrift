package transactions

import (
	"context"

	"github.com/bufbuild/connect-go"
	transactionpb "github.com/midepeter/thrift/gen/proto/transaction"
)

func (t *Transaction) Lock(ctx context.Context, req *connect.Request[transactionpb.LockRequest]) (*connect.Response[transactionpb.LockResponse], error) {
	return nil, nil
}
