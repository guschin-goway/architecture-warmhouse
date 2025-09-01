package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	App AppConfig `yaml:"app" env-prefix:"TEMPERATURE_API_"`
	Log LogConfig `yaml:"log" env-prefix:"TEMPERATURE_API_LOG_"`
}

type AppConfig struct {
	Name      string `yaml:"name" env:"NAME" env-default:"gp"`
	Version   string `yaml:"version" env:"VERSION" env-default:"dev"`
	Env       string `yaml:"env" env:"ENV" env-default:"development"`
	Address   string `yaml:"address" env:"ADDRESS" env-default:"localhost"`
	Port      int    `yaml:"port" env:"PORT" env-default:"80"`
	SSL       bool   `yaml:"ssl" env:"SSL" env-default:"false"`
	SSLEnable bool   `yaml:"ssl_enable" env:"SSL_ENABLE"`
	SSLConfig struct {
		CertPath string `yaml:"cert_path" env:"CERT_PATH"`
		KeyPath  string `yaml:"key_path" env:"KEY_PATH"`
	} `yaml:"ssl_config"`
}

type LogConfig struct {
	Level        string `yaml:"level" env:"LEVEL"`
	Format       string `yaml:"format" env:"FORMAT"`
	EnableCaller bool   `yaml:"enable_caller" env:"ENABLE_CALLER"`
}

func Load() (*Config, error) {
	var cfg Config

	// Читаем конфиг из файла и переменных окружения
	err := cleanenv.ReadConfig("./config/config.yaml", &cfg)
	if err != nil {
		return nil, fmt.Errorf("error reading config: %w", err)
	}

	return &cfg, err
}

func (c *AppConfig) AppAddress() string {
	prefix := "http://"
	address := "localhost"
	if c.SSLEnable {
		prefix = "https://"
		address = c.Address
	}
	return fmt.Sprintf("%s%s:%d", prefix, address, c.Port)
}

func (c *AppConfig) AppPort() string {
	return fmt.Sprintf(":%d", c.Port)
}
