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

func Read() Config {
	configPath := getConfigFilePath()

	data, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Printf("\n Could not read config file. Err: %v", err)
		return Config{}
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Printf("\n Could not convert config to JSON. Err: %v", err)
		return Config{}
	}

	return config
}

func (c *Config) SetUser(user string) {

	c.CurrentUserName = user

	bytes, err := json.Marshal(c)
	if err != nil {
		fmt.Printf("\n Could not convert config to bytes. Err: %v", err)
		return
	}
	configPath := getConfigFilePath()
	err = os.WriteFile(configPath, bytes, os.ModeDevice)
	if err != nil {
		fmt.Printf("\n Could not write to config file. Err: %v", err)
		return
	}
}

func getConfigFilePath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("\n Could not get user's home directory. Err: %v", err)
		return ""
	}

	configPath := homeDir + "/.gatorconfig.json"

	return configPath
}
