package main

import (
	"app/config"

	core "github.com/andreyAKor/core-app-linux-sys"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Init service
	service := core.NewService("app", new(config.Configuration), new(App))

	// Start service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
