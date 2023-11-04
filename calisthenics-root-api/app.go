package main

import (
	"calisthenics-root-api/config"
	"github.com/spf13/viper"
	"os"
)

func Run() {
	env := "dev"
	if len(os.Args) > 1 {
		env = os.Args[1]
	}
	config.InitConfig(env)
	e := InitializeApp()
	serverPort := viper.GetString("server.port")
	e.Logger.Fatal(e.Start(serverPort))
}
