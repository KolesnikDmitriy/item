package app

import (
	"sync"
	"sync/atomic"

	"github.com/KolesnikDmitriy/item/internal/analytics"
	pb "github.com/KolesnikDmitriy/item/pkg/api"
)

type ItemService struct {
	pb.ItemServer

	mu sync.RWMutex

	items  map[int64]*pb.GetItemResponse
	nextID atomic.Int64

	analytics *analytics.Analytics
}

func NewItemService(a *analytics.Analytics) *ItemService {
	return &ItemService{
		items:     make(map[int64]*pb.GetItemResponse, 8),
		analytics: a,
	}
}
