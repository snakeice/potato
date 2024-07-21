package configuration

import (
	"log"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/snakeice/potato/internal/pkg/definitions"
	"gopkg.in/yaml.v2"
)

func filePath() string {
	homeDir := os.Getenv("POTATO_CONFIG")
	if homeDir == "" {
		systemHome, err := homedir.Dir()
		homeDir = systemHome
		if err != nil {
			log.Fatal("Error to get cache paths", err)
		}

		homeDir = filepath.FromSlash(homeDir)
	}

	return filepath.Join(homeDir, ".potato")
}

func LoadConfiguration() *definitions.PotatoConfig {
	configPath := filepath.Join(filePath(), "potato.yaml")
	c := &definitions.PotatoConfig{
		AlwaysSudo: false,
		Commands:   map[string]definitions.Command{"hello": {Description: "Hello world", Template: []string{"echo Hello world"}}},
		Shell:      []string{"/usr/bin/sh", "-c"},
		User:       "",
		Path:       configPath,
	}

	buff, err := os.ReadFile(configPath)
	if err != nil {
		log.Printf("WARN: Fail to load cfg %s", err)
		saveConfiguration(c, filePath())
		return c
	}

	err = yaml.Unmarshal(buff, c)
	if err != nil {
		log.Fatalf("Fail to load cfg %s", err)
		return c
	}

	return c
}

func saveConfiguration(c *definitions.PotatoConfig, path string) {
	str, err := yaml.Marshal(c)
	if err != nil {
		log.Fatal("Error on parse config file", err.Error())
	}

	err = os.MkdirAll(path, 0755)
	if err != nil {
		log.Fatal("Failed on create path", path, err.Error())
	}

	configPath := filepath.Join(path, "potato.yaml")
	f, err := os.Create(configPath)
	if err != nil {
		log.Fatal("Failed on create file ", configPath, err.Error())
	}

	defer f.Close()

	_, err = f.Write(str)
	if err != nil {
		log.Printf("Failed on write cache file: %s", err.Error())
	}
}
