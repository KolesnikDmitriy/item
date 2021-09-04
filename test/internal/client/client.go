package client

import (
	"context"
	"log"
	"time"

	item "github.com/KolesnikDmitriy/item/pkg/api"
	"google.golang.org/grpc"
)

type ItemClient struct {
	item.ItemClient
	timeout time.Duration
}

func New(host string, timeout time.Duration) (*ItemClient, *grpc.ClientConn) {
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("func grpc.Dial; host %q;", host)
	}
	client := ItemClient{
		ItemClient: item.NewItemClient(conn),
		timeout:    timeout,
	}

	return &client, conn
}

func (i *ItemClient) Item(in *item.GetItemRequest) (*item.GetItemResponce, error) {
	ctx, cancel := context.WithTimeout(context.Background(), i.timeout)
	defer cancel()
	return i.ItemClient.GetItem(ctx, in)
}
