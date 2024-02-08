package config

import (
	"io/ioutil"
	"log"

	"github.com/TNXG/ProcessReporterWingo/tools"
	"gopkg.in/yaml.v2"
)

type ServerConfig struct {
	Endpoint   string `yaml:"Endpoint"`
	Token      string `yaml:"Token"`
	ReportTime int    `yaml:"ReportTime"`
}

type Rule struct {
	MatchApplication string  `yaml:"MatchApplication"`
	Replace          Replace `yaml:"Replace"`
}

type Replace struct {
	Application string `yaml:"Application"`
	Description string `yaml:"Description"`
}

type MainConfig struct {
	ServerConfig ServerConfig `yaml:"ServerConfig"`
	Rules        []Rule       `yaml:"Rules"`
}

func LoadConfig() MainConfig {
	workdir, _ := tools.Getwd()
	data, err := ioutil.ReadFile(workdir + "/config.yml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var config MainConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return config
}
