package main

import (
	"github.com/Mohsen20031203/learn-gochain-core/config"
	api "github.com/Mohsen20031203/learn-gochain-core/internal/api/http"
	usecase "github.com/Mohsen20031203/learn-gochain-core/internal/domain/blockchain"
	"github.com/Mohsen20031203/learn-gochain-core/internal/usecase/blockchain"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	chain := usecase.New()
	service := blockchain.NewService(chain)
	handler := api.NewHandler(service)
	server := api.NewServer(cfg, handler)

	if err := server.Start(); err != nil {
		panic(err)
	}
}
