package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	"github.com/KolesnikDmitriy/item/internal/app"
	pb "github.com/KolesnikDmitriy/item/pkg/api"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to Listen: %v", err)
	}
	server := grpc.NewServer()
	app.RegisterNewItemService(server)
	go func() {
		if err := server.Serve(lis); err != nil {
			log.Fatalf("failed to Serve: %v", err)
		}
	}()

	mux := runtime.NewServeMux()
	err = pb.RegisterItemHandlerServer(ctx, mux, app.NewItemService())
	if err != nil {
		log.Fatalf("failed to RegisterItemHandlerServer: %v", err)
	}
	if err := http.ListenAndServe(":50052", mux); err != nil {
		log.Fatalf("failed to ListenAndServe: %v", err)
	}
}
