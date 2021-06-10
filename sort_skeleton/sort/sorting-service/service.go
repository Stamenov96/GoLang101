package main

import (
	"context"
	"errors"
	"math/rand"
	"reflect"

	"github.com/Stamenov96/GoLang101/sort/gen"
)

func newSortingService() gen.SortingRobotServer {
	return &sortingService{}
}

var currentItems []gen.Item
var selectedItem gen.Item

type sortingService struct{}

func (s *sortingService) LoadItems(ctx context.Context, in *gen.LoadItemsRequest) (*gen.LoadItemsResponse, error) {
	for _, item := range in.GetItems() {
		currentItems = append(currentItems, *item)
	}
	return &gen.LoadItemsResponse{}, nil
}

func (s *sortingService) SelectItem(context.Context, *gen.SelectItemRequest) (*gen.SelectItemResponse, error) {

	if len(currentItems) <= 0 {
		return nil, errors.New("no items in the input bin")
	}

	randomIndex := rand.Intn(len(currentItems))
	selectedItem = currentItems[randomIndex]
	currentItems = append(currentItems[:randomIndex], currentItems[randomIndex+1:]...)

	return &gen.SelectItemResponse{Item: &selectedItem}, nil
}

func (s *sortingService) MoveItem(ctx context.Context, in *gen.MoveItemRequest) (*gen.MoveItemResponse, error) {

	if reflect.ValueOf(selectedItem).IsZero() {
		return nil, errors.New("there is no selected item")
	}

	selectedItem = gen.Item{}
	return &gen.MoveItemResponse{}, nil
}
