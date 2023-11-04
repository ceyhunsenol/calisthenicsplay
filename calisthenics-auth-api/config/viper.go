package config

import (
	"github.com/spf13/viper"
	"log"
	"strings"
)

func InitConfig(env string) {
	viper.SetConfigName("application")
	viper.AddConfigPath("./resource")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading the general configuration file: %s", err)
	}

	if env != "" {
		viper.SetConfigName("application-" + env)
		err := viper.MergeInConfig()
		if err != nil {
			log.Fatalf("Error while merging %s configuration file: %s", env, err)
		}
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}
