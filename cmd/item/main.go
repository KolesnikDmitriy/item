package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/KolesnikDmitriy/item/internal/app"
)

const port = ":50051"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	app.RegisterNewItemService(server)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
