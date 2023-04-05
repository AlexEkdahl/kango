package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	Host   string
	Port   int
	DBPath string
	DBName string
}

func New() *Config {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yml")

	// Set default values
	v.SetDefault("dbPath", "~/.config/kango/")
	v.SetDefault("dbName", "kango.db")

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Error finding user home directory: %v", err)
	}

	configPath := filepath.Join(homeDir, ".config", "kango")
	v.AddConfigPath(configPath)

	_ = v.ReadInConfig()

	return &Config{
		Host:   v.GetString("host"),
		Port:   v.GetInt("port"),
		DBPath: v.GetString("dbPath"),
		DBName: v.GetString("dbName"),
	}
}
