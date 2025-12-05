// Package config provides configuration loading and parsing capabilities.
package config

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

type Configuration struct {
	Port            string `json:"port"`
	DefaultLanguage string `json:"default_language"`
	LegacyEndpoint  string `json:"legacy_endpoint"`
	DatabaseType    string `json:"database_type"`
	DatabaseURL     string `json:"database_url"`
}

var defaultConfiguration = Configuration{
	Port:            ":8080",
	DefaultLanguage: "english",
}

// LoadFromEnv will load configuration solely from the environment.
func (c *Configuration) LoadFromEnv() {
	if lang := os.Getenv("DEFAULT_LANGUAGE"); lang != "" {
		c.DefaultLanguage = lang
	}
	if port := os.Getenv("PORT"); port != "" {
		c.Port = port
	}
}

// ParsePort will check to see if the port is in the proper format and a number.
func (c *Configuration) ParsePort() {
	if c.Port[0] != ':' {
		c.Port = ":" + c.Port
	}
	if _, err := strconv.Atoi(string(c.Port[1:])); err != nil {
		fmt.Printf("invalid port %s", c.Port)
		c.Port = defaultConfiguration.Port
	}
}

// LoadFromJSON will read a JSON file and update the configuration based on the file.
func (c *Configuration) LoadFromJSON(filePath string) error {
	log.Printf("loading configuration from file: %s\n", filePath)
	b, err := os.ReadFile(filepath.Clean(filePath))
	if err != nil {
		log.Printf("error load file: %v\n", err)
		return errors.New("unable to load configuration file")
	}
	if err := json.Unmarshal(b, c); err != nil {
		log.Printf("error parsing configuration file: %v\n", err)
		return errors.New("unable to parse configuration file")
	}
	// Verify required fields
	if c.Port == "" {
		log.Printf("empty port, reverting to default: %s\n", defaultConfiguration.Port)
		c.Port = defaultConfiguration.Port
	}
	if c.DefaultLanguage == "" {
		log.Printf("empty default language, reverting to default: %s\n", defaultConfiguration.DefaultLanguage)
		c.DefaultLanguage = defaultConfiguration.DefaultLanguage
	}
	return nil
}

// LoadConfiguration will provide cycle through flags, file, and finally env variables to load configuration.
func LoadConfiguration() Configuration {
	cfgfileFlag := flag.String("config_file", "", "load configuration from file")
	portFlag := flag.String("port", "", "port to run the server on")

	flag.Parse()
	cfg := defaultConfiguration

	if cfgfileFlag != nil && *cfgfileFlag != "" {
		if err := cfg.LoadFromJSON(*cfgfileFlag); err != nil {
			log.Printf("error loading configuration from file: %v\n", err)
			log.Println("using default values")
		}
	}

	cfg.LoadFromEnv()

	if portFlag != nil && *portFlag != "" {
		cfg.Port = *portFlag
	}

	cfg.ParsePort()
	return cfg
}
