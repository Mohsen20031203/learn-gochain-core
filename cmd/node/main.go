package main

import (
	"context"

	"github.com/Mohsen20031203/learn-gochain-core/config"
	api "github.com/Mohsen20031203/learn-gochain-core/internal/api/http"
	"github.com/Mohsen20031203/learn-gochain-core/internal/usecase/blockchain"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	node := blockchain.NewService(cfg)
	handler := api.NewHandler(node)
	server := api.NewServer(cfg, handler)

	node.StartMiner(ctx)

	if err := server.Start(); err != nil {
		panic(err)
	}
}
