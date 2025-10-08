package main

import (
	"fmt"

	"github.com/Mohsen20031203/learn-gochain-core/config"
)

func main() {

	config, err := config.LoadConfig(".")
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	_ = config
}
