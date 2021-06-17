package main

import (
	"context"
	"testing"

	"github.com/Stamenov96/GoLang101/sort_skeleton/sort/gen"
	"github.com/stretchr/testify/assert"
)

func TestLoadItems(t *testing.T) {

	sorting_service := newSortingService()

	req := &gen.LoadItemsRequest{
		Items: []*gen.Item{&gen.Item{}, &gen.Item{}},
	}

	res, err := sorting_service.LoadItems(context.Background(), req)

	assert.NotEqual(t, res, nil, "result should be empty LoadItemsResult")
	assert.Equal(t, err, nil, "there should be no error")

}

func TestSelectItem(t *testing.T) {

	sorting_service := newSortingService()

	load_request := &gen.LoadItemsRequest{
		Items: []*gen.Item{&gen.Item{}, &gen.Item{}},
	}
	sorting_service.LoadItems(context.Background(), load_request)

	res, err := sorting_service.SelectItem(context.Background(), &gen.SelectItemRequest{})

	assert.NotEqual(t, res, nil, "result should be empty SelectItemsResult")
	assert.Equal(t, err, nil, "there should be no error")

}

func TestSelectItemWithEmptyBin(t *testing.T) {

	assert := assert.New(t)
	sorting_service := newSortingService()

	res, err := sorting_service.SelectItem(context.Background(), &gen.SelectItemRequest{})
	assert.Nil(res, "Result should be empty")
	assert.NotNil(err, "There should be no error")
}

func TestMoveItem(t *testing.T) {

	sorting_service := newSortingService()

	//  Load
	load_request := &gen.LoadItemsRequest{
		Items: []*gen.Item{&gen.Item{}, &gen.Item{}},
	}
	sorting_service.LoadItems(context.Background(), load_request)

	// Select
	sorting_service.SelectItem(context.Background(), &gen.SelectItemRequest{})

	res, err := sorting_service.MoveItem(context.Background(), &gen.MoveItemRequest{})

	assert.NotEqual(t, res, nil, "result should be empty SelectItemsResult")
	assert.Equal(t, err, nil, "there should be no error")

}
