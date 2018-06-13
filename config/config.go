package config

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Dir 				string		`yaml:"dir"`
	SaveDir				string		`yaml:"save_dir"`
	ExcludeList			[]string	`yaml:"exclude_list"`
}

var config *Config

func Load(path string) error {
	result, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(result, &config)
}

func Get() *Config {
	return config
}