package config

import (
	"errors"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	DBUser     string `mapstructure:"DB_USER"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	DBPort     string `mapstructure:"DB_PORT"`
	HugoPath   string
}

func (c Config) IsValid() bool {
	if len(c.DBUser) > 0 && len(c.DBHost) > 0 && len(c.DBPassword) > 0 && len(c.DBName) > 0 && len(c.DBPort) > 0 {
		return true
	} else {
		return false
	}
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.AutomaticEnv()
	_, pathErr := os.Stat(filepath.Join(path, ".env"))
	if !errors.Is(pathErr, os.ErrNotExist) {
		viper.SetConfigFile(".env")
		err = viper.ReadInConfig()
	} else {
		for _, k := range viper.AllKeys() {
			v := viper.GetString(k)
			viper.Set(k, os.ExpandEnv(v))
		}
	}

	err = viper.Unmarshal(&config)

	p, _ := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	config.HugoPath = p

	if !config.IsValid() {
		err = errors.New("config is missing required fields")
	}

	return
}
