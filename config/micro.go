package config

// Данные для инициализации сервиса
type Micro struct {
	// Данные для работы с реестром
	Registry        string
	RegistryAddress string

	// Данные для работы с брокером
	Broker        string
	BrokerAddress string

	// Данные для работы с траyспортом
	Transport        string
	TransportAddress string
}
