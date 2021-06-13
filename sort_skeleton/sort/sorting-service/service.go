package main

import (
	"context"
	"errors"
	"math/rand"

	"github.com/Stamenov96/GoLang101/sort_skeleton/sort/gen"
)

func newSortingService() gen.SortingRobotServer {
	return &sortingService{}
}

type sortingService struct {
	currentItems []gen.Item
	selectedItem *gen.Item
}

func (s *sortingService) LoadItems(ctx context.Context, in *gen.LoadItemsRequest) (*gen.LoadItemsResponse, error) {
	for _, item := range in.GetItems() {
		s.currentItems = append(s.currentItems, *item)
	}
	return &gen.LoadItemsResponse{}, nil
}

func (s *sortingService) SelectItem(context.Context, *gen.SelectItemRequest) (*gen.SelectItemResponse, error) {

	if len(s.currentItems) <= 0 {
		return nil, errors.New("no items in the input bin")
	}

	randomIndex := rand.Intn(len(s.currentItems))
	s.selectedItem = &s.currentItems[randomIndex]
	s.currentItems = append(s.currentItems[:randomIndex], s.currentItems[randomIndex+1:]...)
	return &gen.SelectItemResponse{Item: s.selectedItem}, nil
}

func (s *sortingService) MoveItem(ctx context.Context, in *gen.MoveItemRequest) (*gen.MoveItemResponse, error) {

	if s.selectedItem == nil {
		return nil, errors.New("there is no selected item")
	}

	*s.selectedItem = gen.Item{}
	return &gen.MoveItemResponse{}, nil
}
