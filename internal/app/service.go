package app

import (
	pb "github.com/KolesnikDmitriy/item/pkg/api"
)

// ItemService ...
type ItemService struct {
	pb.ItemServer

	// items и nextID намеренно используются без синхронизации — учебное упрощение:
	// параллельные запросы к PostItem/GetItem дают data race.
	// В реальном коде нужен sync.RWMutex (или sync.Map).
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
