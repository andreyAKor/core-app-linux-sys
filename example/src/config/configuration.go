package config

import (
	config "github.com/andreyAKor/core-app-linux-sys/config"
)

type Configuration struct {
	App     config.App     // Информация о приложении
	Logging config.Logging // Настройки логгирования
	Micro   config.Micro   // Данные для инициализации сервиса

	Subscrption Subscrption // Настройки подписки
}
