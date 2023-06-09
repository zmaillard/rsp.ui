package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

type Config struct {
	DBUser     string `mapstructure:"DB_USER"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	DBPort     string `mapstructure:"DB_PORT"`
	HugoPath   string
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)

	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignoring - will try to use environment variables
		} else {
			return //
		}
	}

	err = viper.Unmarshal(&config)

	p, _ := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	config.HugoPath = p
	return
}
