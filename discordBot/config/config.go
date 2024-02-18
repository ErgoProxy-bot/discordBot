package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Token     string `json:"token"`
	BotPrefix string `json:"botPrefix"`
}

func ReadConfig() (Config, error) {
	file, err := ioutil.ReadFile("../config.json") // assuming the configuration is stored in a file named config.json
	if err != nil {
		return Config{}, fmt.Errorf("failed to read config file: %v", err)
	}

	var cfg Config
	err = json.Unmarshal(file, &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("failed to unmarshal config: %v", err)
	}

	return cfg, nil
}
