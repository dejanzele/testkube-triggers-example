package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"net/http"
	"os"
)

var cfg *Config

func main() {
	logger := log.Default()
	logger.Println("Loading application config")
	config, err := loadConfig()
	if err != nil {
		logger.Fatalf("error loading config: %v", err)
	}
	cfg = config
	logger.Println("Starting service on port 8080")
	http.HandleFunc("/health", healthCheckHandler)
	http.ListenAndServe(":8080", nil)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	log.Default().Printf("Crash variable is set to %t", cfg.Crash)
	if cfg.Crash == true {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Ooops, application stopped working :(")
		return
	}
	fmt.Fprint(w, "Everything is running smoothly :)")
}

func loadConfig() (*Config, error) {
	configPath := "config.yaml"
	if path := os.Getenv("CONFIG_PATH"); path != "" {
		configPath = path
	}
	f, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

type Config struct {
	Crash bool `yaml:"crash"`
}
