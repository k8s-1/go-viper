package main

import (
	"fmt"
	"log"
  "github.com/spf13/viper"
)

type Config struct {
    Host        string
    Port        int
    LogLevel    string
    // Add other config fields as needed
}

func LoadConfig() (*Config, error) {
    viper.SetConfigName("config")    // Name of the config file (without extension)
    viper.SetConfigType("yaml")      // OR viper.SetConfigType("json")
    viper.AddConfigPath(".")         // Path to look for the config file

    // Set defaults in case some values are missing in the config file
    viper.SetDefault("Host", "localhost")
    viper.SetDefault("Port", 8080)

    if err := viper.ReadInConfig(); err != nil {
        return nil, err
    }

    config := &Config{
        Host:     viper.GetString("Host"),
        Port:     viper.GetInt("Port"),
        LogLevel: viper.GetString("LogLevel"),
    }

    return config, nil
}

var config *Config

func main() {
    var err error
    config, err = LoadConfig()
    if err != nil {
        log.Fatalf("Error loading configuration: %v", err)
    }

    fmt.Print(config)
}
