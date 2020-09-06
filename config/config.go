package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type ClientConfig struct {
	Host string `yaml:"host"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

type DatabaseConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
}

type EmailConfig struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
}

type Config struct {
	Server             ServerConfig   `yaml:"server"`
	Database           DatabaseConfig `yaml:"database"`
	Email              EmailConfig    `yaml:"email"`
	Client             ClientConfig   `yaml:"client"`
	PasswordSecret     string         `yaml:"passwordSecret"`
	AccessTokenSecret  string         `yaml:"accessTokenSecret"`
	RefreshTokenSecret string         `yaml:"refreshTokenSecret"`
}

func GetConfig() *Config {
	var cfg Config

	f, err := os.Open("config/config.yml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)

	return &cfg
}
