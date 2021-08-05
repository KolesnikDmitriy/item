package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	pb "github.com/KolesnikDmitriy/item/pkg/api"
)

func TestItem(t *testing.T) {
	t.Parallel()

	req := pb.ItemRequest{}
	conn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
	defer conn.Close()
	client := pb.NewItemClient(conn)

	_, err := client.Item(context.Background(), &req)

	require.NoError(t, err)
}
