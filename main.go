package main

import (
	"os"

	"github.com/sleepy-day/ER-BE/er_api"
	"gopkg.in/yaml.v3"
)

var AppConfig Config
var ApiManager er_api.ERApiManager

func main() {
	setConfig()
	ApiManager = er_api.BuildERApiManager(AppConfig.ApiKey.Secret)

	ApiManager.GetI10nFiles()
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
