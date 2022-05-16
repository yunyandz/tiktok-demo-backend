package config

import (
	"os"
	"sync"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Http struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"http"`
	Redis struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Password string `yaml:"password"`
		Database int    `yaml:"database"`
	} `yaml:"redis"`
	Mysql struct {
		User     string `yaml:"user"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	} `yaml:"mysql"`
	S3 struct {
		AccessKey string `yaml:"accesskey"`
		SecretKey string `yaml:"secretkey"`
		Region    string `yaml:"region"`
		Bucket    string `yaml:"bucket"`
	} `yaml:"s3"`
	Kafka struct {
		Brokers []string `yaml:"brokers"`
	} `yaml:"kafka"`
}

var (
	cfg  *Config
	once sync.Once
)

func Phase() (*Config, error) {
	once.Do(func() {
		cfg = &Config{}
		configfile, err := os.ReadFile("config.yaml")
		if err != nil {
			panic(err)
		}
		if len(configfile) == 0 {
			panic("config file is empty")
		}
		err = yaml.Unmarshal([]byte(configfile), cfg)
		if err != nil {
			panic(err)
		}
	})
	return cfg, nil
}
