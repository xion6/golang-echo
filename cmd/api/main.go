package main

import (
	"github.com/xion6/golang-echo/internal/api"
	"github.com/xion6/golang-echo/pkg/config"
	"github.com/xion6/golang-echo/pkg/data"
)

func main() {
	cfg := config.New()
	db := data.NewMongoConnection(cfg)
	defer db.Disconnect()
	application := api.New(cfg, db.Client)
	application.Start()
}
