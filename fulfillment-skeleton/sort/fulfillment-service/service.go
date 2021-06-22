package main

import (
	"context"
	"errors"
	"log"

	"github.com/Stamenov96/GoLang101/fulfillment-service/sort/gen"
	"github.com/preslavmihaylov/ordertocubby"
)

func newFulfillmentService() gen.FulfillmentServer {
	return &fulfillmentService{}
}

type fulfillmentService struct{}

func (fs *fulfillmentService) LoadOrders(ctx context.Context, in *gen.LoadOrdersRequest) (*gen.CompleteResponse, error) {
	mapOrderCubby := mappingOrderToCubby(in.Orders)
	log.Print(mapOrderCubby)
	return nil, errors.New("not implemented")
}

func mappingOrderToCubby(orders []*gen.Order) map[string]string {

	mapOrderToCubby := make(map[string]string)
	mapUniqueValues := make(map[string]struct{})
	seed := uint32(0)
	for _, order := range orders {
		cubby_id := ordertocubby.Map(order.Id, 0, 10)
		_, ok := mapUniqueValues[cubby_id]

		for ok {
			seed++
			cubby_id = ordertocubby.Map(order.Id, seed, 10)
			_, ok = mapUniqueValues[cubby_id]
		}

		mapOrderToCubby[order.Id] = cubby_id
		mapUniqueValues[cubby_id] = struct{}{}
		log.Print()
	}

	return mapOrderToCubby
}
