package main

import (
	"github.com/joho/godotenv"
	"gitub.com/Vinicamilotti/homelab-lain-utils-servicecli/cli"
	confighandler "gitub.com/Vinicamilotti/homelab-lain-utils-servicecli/configHandler"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	config := confighandler.CreateServicesConfig()

	cliApp := cli.NewCli(config)
	cliApp.RunApp()

}
