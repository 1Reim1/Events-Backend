package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	BindAddr string
	LogLevel string
	DataFile string
}

func NewConfig() (*Config, error) {
	config := Config{}
	content, err := ioutil.ReadFile("config\\config.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(content, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
