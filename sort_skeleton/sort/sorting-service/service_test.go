package main

import (
	"context"
	"testing"

	"github.com/Stamenov96/GoLang101/sort/gen"
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
