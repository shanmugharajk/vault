package utils

import (
	"encoding/json"
	"log"
	"os"

	"github.com/shanmugharajk/vault/internal/file"
)

type Config struct {
	DbPath     string
	BackUpPath string
}

func GetConfig() Config {
	configPath := ConfigFilePath()
	defaultConfig := Config{
		DbPath:     GetExecutablePath(),
		BackUpPath: GetExecutablePath(),
	}

	if file.CreateIfNotExist(configPath) {
		file.WriteFile(configPath, defaultConfig)
		return defaultConfig
	}

	content, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal("config file is missing")
	}

	var config Config

	err = json.Unmarshal(content, &config)
	if err != nil {
		log.Fatal("invalid config file")
	}

	return config
}
