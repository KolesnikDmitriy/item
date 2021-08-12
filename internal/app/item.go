package app

import (
	"context"
	"errors"

	pb "github.com/KolesnikDmitriy/item/pkg/api"
)

func (s *ItemService) Item(ctx context.Context, in *pb.ItemRequest) (*pb.ItemResponce, error) {
	if in.Id <= 0 {
		return nil, errors.New("wrong id")
	}
	return &pb.ItemResponce{Title: "Book", Description: "Very Great Book"}, nil
}
