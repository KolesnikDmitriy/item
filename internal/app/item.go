package app

import (
	"context"
	"errors"

	pb "github.com/KolesnikDmitriy/item/pkg/api"
)

func (s *ItemService) GetItem(ctx context.Context, in *pb.GetItemRequest) (*pb.GetItemResponse, error) {
	if in.Id <= 0 {
		return nil, errors.New("wrong id")
	}

	item, ok := s.items[in.Id]

	if !ok {
		return nil, errors.New("item not found")
	}
	return &pb.GetItemResponse{Title: item.Title, Description: item.Description}, nil
}

func (s *ItemService) PostItem(ctx context.Context, in *pb.PostItemRequest) (*pb.PostItemResponse, error) {
	if in.Title == "" {
		return nil, errors.New("wrong title")
	}
	if in.Description == "" {
		return nil, errors.New("wrong description")
	}

	id := s.nextID
	s.nextID++
	s.items[id] = &pb.GetItemResponse{Title: in.Title, Description: in.Description}

	return &pb.PostItemResponse{Id: id}, nil
}
