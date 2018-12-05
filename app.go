package core

// Интерфейс приложения
type App interface {
	Init(configuration interface{})
	Run()
}
