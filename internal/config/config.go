package config

import (
	"os"
	"sync"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Debug bool  `yaml:"debug"` // 在DEBUG模式下，会打印更多日志
	Http  Http  `yaml:"http"`
	Redis Redis `yaml:"redis"`
	Mysql Mysql `yaml:"mysql"`
	S3    S3    `yaml:"s3"`
	Kafka Kafka `yaml:"kafka"`
}

type Http struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Redis struct {
	Vaild    bool   `yaml:"vaild"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	Database int    `yaml:"database"`
}

type Mysql struct {
	User     string `yaml:"user"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type S3 struct {
	Vaild     bool   `yaml:"vaild"` //是否启用s3,如果关闭的话，则不上传视频到s3。
	Endpoint  string `yaml:"endpoint"`
	AccessKey string `yaml:"accesskey"`
	SecretKey string `yaml:"secretkey"`
	Region    string `yaml:"region"`
	Bucket    string `yaml:"bucket"`
}

type Kafka struct {
	Vaild   bool     `yaml:"vaild"` //是否启用kafka,仅供测试使用
	Brokers []string `yaml:"brokers"`
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
