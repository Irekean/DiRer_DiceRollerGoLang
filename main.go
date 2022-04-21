package main

import (
	"fmt"
	"irekean-discord-direr/bot"
	"irekean-discord-direr/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}
