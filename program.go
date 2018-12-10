package core

import (
	"github.com/andreyAKor/core-app-linux-sys/config"

	"github.com/kardianos/service"
	log "github.com/sirupsen/logrus"
)

// Структура программы
type Program struct {
	exit              chan struct{}
	service           service.Service
	app               App
	coreConfiguration *config.Configuration
}

// Обработчик старта сервиса
func (p *Program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	log.Info("Start")

	go p.run()
	return nil
}

// Обработчик программы сервиса
func (p *Program) run() {
	log.Info("Runnig ", p.coreConfiguration.App.DisplayName)

	defer func() {
		// Смотрим наличие менеджеров сервисов в ОС
		if service.ChosenSystem() != nil {
			if service.Interactive() {
				p.Stop(p.service)
			} else {
				p.service.Stop()
			}
		}
	}()

	// Запускаем обработчик приложения
	p.app.Run()
}

// Обработчик остановки сервиса
func (p *Program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	log.Info("Stop")

	return nil
}
