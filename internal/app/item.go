package app

import (
	"context"
	"time"

	pb "github.com/KolesnikDmitriy/item/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ItemService) GetItem(ctx context.Context, in *pb.GetItemRequest) (*pb.GetItemResponse, error) {
	start := time.Now()
	defer func() { s.analytics.Log("GET", time.Since(start)) }()

	if in.Id <= 0 {
		return nil, status.Error(codes.InvalidArgument, "wrong id")
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	item, ok := s.items[in.Id]
	if !ok {
		return nil, status.Error(codes.NotFound, "item not found")
	}
	return &pb.GetItemResponse{Title: item.Title, Description: item.Description}, nil
}
func (s *ItemService) PostItem(ctx context.Context, in *pb.PostItemRequest) (*pb.PostItemResponse, error) {
	start := time.Now()
	defer func() { s.analytics.Log("POST", time.Since(start)) }()

	if in.Title == "" {
		return nil, status.Error(codes.InvalidArgument, "wrong title")
	}
	if in.Description == "" {
		return nil, status.Error(codes.InvalidArgument, "wrong description")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	id := s.nextID.Add(1)
	s.items[id] = &pb.GetItemResponse{Title: in.Title, Description: in.Description}

	return &pb.PostItemResponse{Id: id}, nil
}
