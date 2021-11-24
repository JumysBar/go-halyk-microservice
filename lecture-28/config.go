package main

import (
	"github.com/go-playground/validator"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	InstanceName string `json:"service_name" env:"SERVICE_NAME" `
	MaxThreads   int64  `json:"max_threads"`

	DatabaseConfig Database `json:"postgres"`
	KafkaConfig    Kafka    `json:"kafka"`
	RedisConfig    Redis    `json:"redis"`

	SomePort string `json:"some_port" validate:"numeric"`
}

type Database struct {
	User   string `json:"user" env:"POSTGRES_USER"`
	Passwd string `json:"password" env:"POSTGRES_PASSWORD"`
	Host   string `json:"host"`
	Port   int    `json:"port" env-default:"5432"`
	DBName string `json:"db_name"`
}

type Kafka struct {
	KafkaBrokers            []string `json:"brokers"`
	KafkaUser               string   `json:"user" env:"KAFKA_USER"`
	KafkaPass               string   `json:"password" env:"KAFKA_PASSWORD"`
	KafkaTopicConsume       string   `json:"topic_consume"`
	KafkaDataRowTypeProduce string   `json:"row_type_produce" env-default:"json"`
	MaxFetchSize            int      `json:"max_fetch_size" env-default:"10000"`
	MaxWait                 int64    `json:"max_wait" env-default:"10"`
}

type Redis struct {
	RedisAddress string `json:"address"`
	RedisUser    string `json:"user" env:"REDIS_USER"`
	RedisPass    string `json:"password" env:"REDIS_PASSWORD"`
	RedisDB      int    `json:"db" env-default:"0"`
}

// GetConfig - read, valid and set data from config file
func GetConfig(path string) (Config, error) {
	var config Config

	if err := cleanenv.ReadConfig(path, &config); err != nil {
		return config, err
	}

	v := validator.New()
	if err := v.Struct(config); err != nil {
		return Config{}, err
	}

	return config, nil
}
