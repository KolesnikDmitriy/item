package app

import (
	"context"

	pb "github.com/KolesnikDmitriy/item/pkg/api"
)

func (s *ItemService) Item(ctx context.Context, in *pb.ItemRequest) (*pb.ItemResponce, error) {
	return &pb.ItemResponce{}, nil
}
