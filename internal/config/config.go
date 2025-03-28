package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
    DBUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	configPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}, err
	}

	var cfg Config
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func (c Config) SetUser() error {
	c.CurrentUserName = "postgres"
	err := write(c)
	if err != nil {
		return err
	}

	return nil
}

func getConfigFilePath() (string, error) {
	dir, dirErr := os.UserConfigDir()

	if dirErr != nil {
		log.Fatal(dirErr)
		return "", dirErr
	}

	configPath := filepath.Join(dir, configFileName)
	return configPath, nil
}

func write(cfg Config) error {
	configPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(cfg, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(configPath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
