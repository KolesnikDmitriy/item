package app

import (
	"context"
	"errors"

	pb "github.com/KolesnikDmitriy/item/pkg/api"
)

func (s *ItemService) GetItem(ctx context.Context, in *pb.GetItemRequest) (*pb.GetItemResponce, error) {
	if in.Id <= 0 {
		return nil, errors.New("wrong id")
	}
	return &pb.GetItemResponce{Title: "Book", Description: "Very Great Book"}, nil
}
