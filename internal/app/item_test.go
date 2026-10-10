package app

import (
	"context"
	"testing"

	pb "github.com/KolesnikDmitriy/item/pkg/api"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGetItem(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		id   int64
		code codes.Code
	}{
		{
			name: "wrong id",
			id:   0,
			code: codes.InvalidArgument,
		},
		{
			name: "negative id",
			id:   -5,
			code: codes.InvalidArgument,
		},
		{
			name: "item not found",
			id:   999,
			code: codes.NotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			svc := NewItemService()

			got, err := svc.GetItem(context.Background(), &pb.GetItemRequest{Id: tt.id})

			require.Error(t, err)
			require.Nil(t, got)
			require.Equal(t, tt.code, status.Code(err))
		})
	}
}

func TestGetItemSuccess(t *testing.T) {
	t.Parallel()

	svc := NewItemService()
	svc.items[1] = &pb.GetItemResponse{Title: "Book", Description: "Very Great Book"}

	got, err := svc.GetItem(context.Background(), &pb.GetItemRequest{Id: 1})

	require.NoError(t, err)
	require.Equal(t, &pb.GetItemResponse{Title: "Book", Description: "Very Great Book"}, got)
}

func TestPostItem(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		title       string
		description string
		code        codes.Code
	}{
		{
			name:  "empty title",
			title: "", description: "Very Great Book",
			code: codes.InvalidArgument,
		},
		{
			name:  "empty description",
			title: "Book", description: "",
			code: codes.InvalidArgument,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			svc := NewItemService()

			got, err := svc.PostItem(context.Background(), &pb.PostItemRequest{
				Title:       tt.title,
				Description: tt.description,
			})

			require.Error(t, err)
			require.Nil(t, got)
			require.Equal(t, tt.code, status.Code(err))
		})
	}
}

func TestPostItemSuccess(t *testing.T) {
	t.Parallel()

	svc := NewItemService()

	got, err := svc.PostItem(context.Background(), &pb.PostItemRequest{
		Title:       "Book",
		Description: "Very Great Book",
	})

	require.NoError(t, err)
	require.Equal(t, int64(1), got.Id)
}

func TestPostItemIncrementID(t *testing.T) {
	t.Parallel()

	svc := NewItemService()
	ctx := context.Background()

	first, err := svc.PostItem(ctx, &pb.PostItemRequest{Title: "a", Description: "b"})
	require.NoError(t, err)
	second, err := svc.PostItem(ctx, &pb.PostItemRequest{Title: "c", Description: "d"})
	require.NoError(t, err)

	require.Equal(t, first.Id+1, second.Id)
}

func TestPostGetFlow(t *testing.T) {
	t.Parallel()

	svc := NewItemService()
	ctx := context.Background()

	post, err := svc.PostItem(ctx, &pb.PostItemRequest{Title: "Book", Description: "Very Great Book"})
	require.NoError(t, err)

	got, err := svc.GetItem(ctx, &pb.GetItemRequest{Id: post.Id})
	require.NoError(t, err)
	require.Equal(t, &pb.GetItemResponse{Title: "Book", Description: "Very Great Book"}, got)
}

func BenchmarkPostItem(b *testing.B) {
	svc := NewItemService()
	ctx := context.Background()
	req := &pb.PostItemRequest{Title: "Book", Description: "Very Great Book"}

	b.ResetTimer()
	for b.Loop() {
		_, err := svc.PostItem(ctx, req)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGetItem(b *testing.B) {
	svc := NewItemService()
	ctx := context.Background()
	post, err := svc.PostItem(ctx, &pb.PostItemRequest{Title: "Book", Description: "Very Great Book"})
	if err != nil {
		b.Fatal(err)
	}

	req := &pb.GetItemRequest{Id: post.Id}

	b.ResetTimer()
	for b.Loop() {
		_, err := svc.GetItem(ctx, req)
		if err != nil {
			b.Fatal(err)
		}
	}
}
