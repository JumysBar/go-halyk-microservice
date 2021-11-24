package main

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	InstanceName string `mapstructure:"service_name"`
	MaxThreads   int64  `mapstructure:"max_threads"`

	DatabaseConfig Database `mapstructure:"postgres"`

	SomePort string `mapstructure:"some_port" validate:"numeric"`
}

type Database struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	DBName   string `mapstructure:"db_name"`
}

// GetConfig - read, valid and set data from config file
func GetConfig() (Config, error) {
	c := Config{}

	viper.SetDefault("some_port", "444")

	viper.SetConfigName("configWithoutCreds")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetEnvPrefix("APP")

	if err := viper.ReadInConfig(); err != nil {
		return Config{}, fmt.Errorf("Fatal error config file: %w \n", err)
	}

	if err := viper.Unmarshal(&c); err != nil {
		return Config{}, err
	}

	return c, nil
}
