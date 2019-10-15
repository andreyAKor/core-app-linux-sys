package config

// Настройки логгирования
type Logging struct {
	Level          string // Уровень логгирования, варианты: debug, info, warning, error, fatal и panic
	IsReportCaller bool   // Показывать информацию о в каком месте был вызван лог
}
