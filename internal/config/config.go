package internal

import (
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
	Log      Log      `yaml:"log"`
}

type Server struct {
	Port string `yaml:"port"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type Log struct {
	Level string `yaml:"level"`
}

func LoadConfig() *Config {
	godotenv.Load()
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH environment variable not set")
	}
	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
	}
	if user := os.Getenv("DB_USER"); user != "" {
		config.Database.Username = user
	}
	if pass := os.Getenv("DB_PASSWORD"); pass != "" {
		config.Database.Password = pass
	}
	return &config
}
