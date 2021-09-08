package tests

import (
	"testing"
	"time"

	"github.com/KolesnikDmitriy/item/test/internal/client"
	"github.com/KolesnikDmitriy/item/test/internal/config"
)

var (
	api *client.ItemClient
	cfg config.Config
)

func TestMain(m *testing.M) {
	cfg = config.New()
	client, conn := client.New(cfg.GrpcUrl, time.Second)
	defer conn.Close()
	api = client

	m.Run()
}
