package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Database Database `yaml:"database"`
}

type Database struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
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
