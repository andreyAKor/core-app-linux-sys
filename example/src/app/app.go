package main

import (
	"os"
	"time"

	"app/config"

	// Указываем свои каналы связи
	_ "github.com/micro/go-plugins/broker/nats"
	_ "github.com/micro/go-plugins/registry/nats"
	_ "github.com/micro/go-plugins/transport/nats"

	core "github.com/andreyAKor/core-app-linux-sys"
	"github.com/micro/go-micro"
	log "github.com/sirupsen/logrus"
)

// Структура приложения
type App struct {
	core.App

	configuration *config.Configuration
	service       micro.Service
}

// Инициализация micro
func (this *App) InitMicro() {
	// Готовим переменное окружение для micro
	os.Setenv("MICRO_REGISTRY", this.configuration.Micro.Registry)
	os.Setenv("MICRO_REGISTRY_ADDRESS", this.configuration.Micro.RegistryAddress)
	os.Setenv("MICRO_BROKER", this.configuration.Micro.Broker)
	os.Setenv("MICRO_BROKER_ADDRESS", this.configuration.Micro.BrokerAddress)
	os.Setenv("MICRO_TRANSPORT", this.configuration.Micro.Transport)
	os.Setenv("MICRO_TRANSPORT_ADDRESS", this.configuration.Micro.TransportAddress)

	// Экземпляр сервиса
	this.service = micro.NewService(
		micro.Name(this.configuration.App.Name),
		micro.Version(this.configuration.App.Version),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// optionally setup command line usage
	this.service.Init()
}

// Инициализация приложения
func (this *App) Init(configuration interface{}) {
	this.configuration = configuration.(*config.Configuration)

	// Инициализация micro
	this.InitMicro()
}

// Обработчик приложения
func (this *App) Run() {
	// Run server
	if err := this.service.Run(); err != nil {
		log.Fatal(err)
	}
}
