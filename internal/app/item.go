package app

import (
	"context"

	pb "github.com/KolesnikDmitriy/item/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ItemService) GetItem(ctx context.Context, in *pb.GetItemRequest) (*pb.GetItemResponse, error) {
	if in.Id <= 0 {
		return nil, status.Error(codes.InvalidArgument, "wrong id")
	}

	item, ok := s.items[in.Id]

	if !ok {
		return nil, status.Error(codes.NotFound, "item not found")
	}
	return &pb.GetItemResponse{Title: item.Title, Description: item.Description}, nil
}

func (s *ItemService) PostItem(ctx context.Context, in *pb.PostItemRequest) (*pb.PostItemResponse, error) {
	if in.Title == "" {
		return nil, status.Error(codes.InvalidArgument, "wrong title")
	}
	if in.Description == "" {
		return nil, status.Error(codes.InvalidArgument, "wrong description")
	}

	id := s.nextID
	s.nextID++
	s.items[id] = &pb.GetItemResponse{Title: in.Title, Description: in.Description}

	return &pb.PostItemResponse{Id: id}, nil
}
