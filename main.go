package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var AppConfig Config

func main() {
	setConfig()
	fmt.Println(AppConfig)
}

func setConfig() {
	f, err := os.Open("config.yaml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		panic(err)
	}

	AppConfig = cfg
}

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Username string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"database"`
	ApiKey struct {
		Secret string `yaml:"secret"`
	} `yaml:"apikey"`
}
