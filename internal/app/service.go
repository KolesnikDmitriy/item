package app

import (
	pb "github.com/KolesnikDmitriy/item/pkg/api"
)

// ItemService ...
type ItemService struct {
	pb.ItemServer

	items  map[int64]*pb.GetItemResponse
	nextID int64
}

// NewItemService ...
func NewItemService() *ItemService {
	return &ItemService{
		items:  make(map[int64]*pb.GetItemResponse, 8),
		nextID: 1,
	}
}
