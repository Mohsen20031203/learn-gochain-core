package config

import (
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Port            int      `mapstructure:"API_PORT"`
	Difficulty      int      `mapstructure:"BLOCKCHAIN_DIFFICULTY"`
	FileStoragePath string   `mapstructure:"FILE_STORAGE_PATH"`
	NodeID          string   `mapstructure:"NODE_ID"`
	BatchSize       int      `mapstructure:"BATCH_SIZE"`
	Peers           []string `mapstructure:"PEERS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.SetConfigFile(path + "/app.env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}

	portStr := strconv.Itoa(config.Port)
	for i := 0; i < len(config.Peers); i++ {
		if strings.Contains(config.Peers[i], portStr) {
			config.Peers = append(config.Peers[:i], config.Peers[i+1:]...)
			break
		}
	}
	if config.Port == 0 {
		viper.SetDefault("API_PORT", 9090)
	}

	if config.Difficulty == 0 {
		viper.SetDefault("BLOCKCHAIN_DIFFICULTY", 3)
	}

	if config.FileStoragePath == "" {
		viper.SetDefault("FILE_STORAGE_PATH", "chainDB")
	}

	return
}
