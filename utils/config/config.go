package config

import (
	"github.com/spf13/viper"
	"log"
	"path/filepath"
)

func InitConfig(path string) *viper.Viper {
	config := viper.New()
	config.SetConfigFile(filepath.ToSlash(path))
	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("errno is %+v", err)
		return nil
	}
	return config
}
