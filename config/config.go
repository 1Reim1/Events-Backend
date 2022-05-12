package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	BindAddr         string
	DatabaseHost     string
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string
}

func NewConfig() (*Config, error) {
	config := Config{}
	content, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(content, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
