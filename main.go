package main

import (
	"IroBot/config"
	"IroBot/controller"
	"fmt"

	"github.com/sirupsen/logrus"
)

func Init() {
	// Initialize the config files.
	config.Init("./config.yaml")
}

func main() {
	Init()
	if err := controller.PowerBot(); err != nil {
		logrus.Fatal(fmt.Sprintf("Fail to start the bot with error message: [ %s ]", err.Error()))
	}
}
