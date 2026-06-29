package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	configPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Printf("\n Could not read config file. Err: %v", err)
		return Config{}, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Printf("\n Could not convert config to JSON. Err: %v", err)
		return Config{}, err
	}

	return config, nil
}

func (c *Config) SetUser(user string) error {

	c.CurrentUserName = user

	bytes, err := json.Marshal(c)
	if err != nil {
		return fmt.Errorf("\n Could not convert config to bytes. Err: %w", err)
	}
	configPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	err = os.WriteFile(configPath, bytes, os.ModeDevice)
	if err != nil {
		return fmt.Errorf("\n Could not write to config file. Err: %w", err)
	}

	return nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("\n Could not get user's home directory. Err: %w", err)
	}

	configPath := homeDir + "/.gatorconfig.json"

	return configPath, nil
}
