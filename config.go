package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

// объект конфига содержит все что в ямл файле

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-default:"local" env-required:"true"`
	StoragePath string `yaml:"storagePath" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeOut time.Duration `yaml:"idleTimeOut" env-default:"60s"`
}

// функция которая прочитает файл с конфигом и заполнит объект конфиг

func MustLoad() *Config {
	configPath := "/home/white/GolandProjects/url-short/config/local.yaml"
	err := os.Setenv("CONFIG_PATH", configPath) // считывание файла с конфигом (из переменной окружения)
	// если не найдется файл
	if err != nil {
		log.Fatal("CONFIG_PATH is not set")
	}

	envValue := os.Getenv("CONFIG_PATH")
	fmt.Println("CONFIG_PATH value:", envValue)
	// проверяем существует ли такой файл
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}
	// объявляем объект конфига
	var cfg Config
	// считываем файл по пути, который указан
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("can not read config: %s", err)
	}
	return &cfg
}
