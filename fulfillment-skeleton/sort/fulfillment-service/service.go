package main

import (
	"context"
	"errors"

	"github.com/Stamenov96/GoLang101/fulfillment-service/sort/gen"
)

func newFulfillmentService() gen.FulfillmentServer {
	return &fulfillmentService{}
}

type fulfillmentService struct{}

func (fs *fulfillmentService) LoadOrders(ctx context.Context, in *gen.LoadOrdersRequest) (*gen.CompleteResponse, error) {
	return nil, errors.New("not implemented")
}
