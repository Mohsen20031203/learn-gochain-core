package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Port            string   `mapstructure:"API_PORT"`
	Difficulty      int      `mapstructure:"BLOCKCHAIN_DIFFICULTY"`
	FileStoragePath string   `mapstructure:"FILE_STORAGE_PATH"`
	NodeID          string   `mapstructure:"NODE_ID"`
	BatchSize       int      `mapstructure:"BATCH_SIZE"`
	Peers           []string `mapstructure:"PEERS"`
	TCPAddress      string   `mapstructure:"TCP_PORT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	fmt.Println(config.NodeID)

	for i, p := range config.Peers {
		if strings.Contains(p, config.TCPAddress) {
			config.Peers = append(config.Peers[:i], config.Peers[i+1:]...)
		}
	}

	if config.Port == "" {
		viper.SetDefault("API_PORT", "9090")
	}

	if config.Difficulty == 0 {
		viper.SetDefault("BLOCKCHAIN_DIFFICULTY", 3)
	}

	if config.FileStoragePath == "" {
		viper.SetDefault("FILE_STORAGE_PATH", "chainDB")
	}

	for _, p := range config.Peers {
		fmt.Println("port : ", p)
	}
	return
}
