package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/KolesnikDmitriy/item/internal/app"
	pb "github.com/KolesnikDmitriy/item/pkg/api"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
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

	httpMux := http.NewServeMux()
	httpMux.Handle("GET /docs/", serveBytes("text/html; charset=utf-8", docsHTML))
	httpMux.Handle("GET /swagger.json", serveBytes("application/json", pb.SwaggerJSON))
	httpMux.Handle("/", mux)

	if err := http.ListenAndServe(":50052", httpMux); err != nil {
		log.Fatalf("failed to ListenAndServe: %v", err)
	}
}

func serveBytes(contentType string, body []byte) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", contentType)
		_, _ = w.Write(body)
	})
}
