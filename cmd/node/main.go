package main

import (
	"github.com/Mohsen20031203/learn-gochain-core/config"
	api "github.com/Mohsen20031203/learn-gochain-core/internal/api/http"
	"github.com/Mohsen20031203/learn-gochain-core/internal/usecase/blockchain"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	service := blockchain.NewService(cfg)
	handler := api.NewHandler(service)
	server := api.NewServer(cfg, handler)

	if err := server.Start(); err != nil {
		panic(err)
	}
}
