package app

import (
	"google.golang.org/grpc"

	pb "github.com/KolesnikDmitriy/item/pkg/api"
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
