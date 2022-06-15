package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Database   Database   `yaml:"database"`
	MarketData MarketData `yaml:"market_data"`
}

type Database struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
}

type MarketData struct {
	ApiKey  string `yaml:"api_key"`
	BaseUrl string `yaml:"base_url"`
}

func LoadConfig(path string) (Config, error) {
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	var config Config

	err = yaml.Unmarshal(fileContent, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
