package config

type Configuration struct {
	App     App     // Информация о приложении
	Logging Logging // Настройки логгирования
	Micro   Micro   // Данные для инициализации сервиса
}
