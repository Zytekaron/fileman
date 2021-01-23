package main

import (
	"errors"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Domain string `yaml:"domain"`
	Path   string `yaml:"path"`
	Port   int    `yaml:"port"`
}

var configDirs = []string{
	"/etc/opt/fileman.yml",
	"config.yml",
}

func loadConfig() (cfg *Config, err error) {
	var file *os.File
	for i := 0; i < len(configDirs); i++ {
		file, err = os.Open(configDirs[i])
		if err != nil {
			continue
		}

		err = yaml.NewDecoder(file).Decode(&cfg)
		return
	}

	return nil, errors.New("could not find any config files")
}
