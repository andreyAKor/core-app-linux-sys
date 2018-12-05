package main

import (
	"config"

	core "github.com/andreyAKor/core-app-linux-sys"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Init service
	service := core.NewService("app.conf", new(config.Configuration), new(App))

	// Start service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
