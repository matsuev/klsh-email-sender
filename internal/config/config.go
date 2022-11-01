package config

import (
	"flag"

	"github.com/ilyakaznacheev/cleanenv"
)

const DEFAULT_PATH string = "config.yml"
const EMPTY_STRING string = ""

// AppConfig struct
type AppConfig struct {
	Host         string `yaml:"host" env-default:"localhost"`
	Port         string `yaml:"port" env-default:"25"`
	DataFile     string `yaml:"data" env-default:"list.csv"`
	TemplatePath string `yaml:"template" env-default:"template.tpl"`
}

// Create function
func Create() (*AppConfig, error) {
	// Read -c flag from command line
	cfgPath := flag.String("c", "", "Path to config file. Config must be in YAML format")
	csvPath := flag.String("t", "", "Path to CSV file that contains a email's and data")
	flag.Parse()

	if *cfgPath == EMPTY_STRING {
		*cfgPath = DEFAULT_PATH
	}

	cfg := &AppConfig{}

	if err := cleanenv.ReadConfig(*cfgPath, cfg); err != nil {
		return nil, err
	}

	if *csvPath != EMPTY_STRING {
		cfg.DataFile = *csvPath
	}

	return cfg, nil
}
