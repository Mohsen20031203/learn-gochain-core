package main

import (
	"fmt"

	"github.com/Mohsen20031203/learn-gochain-core/config"
	"github.com/Mohsen20031203/learn-gochain-core/internal/api"
)

func main() {

	config, err := config.LoadConfig(".")
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	server := api.NewServer(config)
	server.Start(":" + config.ApiPort)

}
