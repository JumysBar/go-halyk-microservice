package main

import (
	"github.com/spf13/viper"
	"log"
)

func getRemoteConf() {
	cfg := viper.New()

	viper.SetConfigName("configWithoutCreds")
	viper.SetConfigType("json")

	cfg.AllowEmptyEnv(true)
	cfg.AutomaticEnv()

	err := cfg.ReadInConfig()
	if err != nil {
		log.Fatalf(err.Error())
	}

	// подключение к консулу для чтения конфигурационного файла, если APP_USE_CONSUL = true
	if cfg.GetBool("use_consul") {
		if err := cfg.AddRemoteProvider(
			"consul",
			cfg.GetString("consul_addr"),
			cfg.GetString("my_conf_prefix")); err != nil {
			log.Fatalf(err.Error())
		}

		cfg.SetConfigType("json")
		err = cfg.ReadRemoteConfig()
		if err != nil {
			log.Fatalf("Error when Fetching Configuration from Consul - %s", err)
		}
	}
}
