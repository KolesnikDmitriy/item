package app

import (
	pb "github.com/KolesnikDmitriy/item/pkg/api"
	"google.golang.org/grpc"
)

// ItemService ...
type ItemService struct {
	pb.ItemServer
}

// NewItemService ...
func NewItemService() *ItemService {
	return &ItemService{}
}

// RegisterNewItemService ...
func RegisterNewItemService(s *grpc.Server) {
	pb.RegisterItemServer(s, NewItemService())
}
