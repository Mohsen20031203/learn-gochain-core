package config

import "github.com/spf13/viper"

type Config struct {
	Port            string `mapstructure:"API_PORT"`
	Difficulty      int    `mapstructure:"BLOCKCHAIN_DIFFICULTY"`
	FileStoragePath string `mapstructure:"FILE_STORAGE_PATH"`
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

	if config.Port == "" {
		viper.SetDefault("API_PORT", "9090")
	}

	if config.Difficulty == 0 {
		viper.SetDefault("BLOCKCHAIN_DIFFICULTY", 3)
	}

	if config.FileStoragePath == "" {
		viper.SetDefault("FILE_STORAGE_PATH", "chainDB")
	}
	err = viper.Unmarshal(&config)

	return
}
