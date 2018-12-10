package core

import (
	"errors"
	"flag"

	"github.com/andreyAKor/core-app-linux-sys/config"

	"github.com/kardianos/osext"
	"github.com/kardianos/service"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	// Инициализация системного лога
	logger service.Logger
)

// Структура службы
type Service struct {
	configName        string                // Имя файла конфига
	app               App                   // Экземпляр приложения
	configuration     interface{}           // Структура конфига приложения
	coreConfiguration *config.Configuration // Структура конфига ядра
}

// Конструктор
func NewService(configName string, configuration interface{}, app App) *Service {
	log.Debug("Construct NewService")

	return &Service{
		configName:    configName,
		app:           app,
		configuration: configuration,
	}
}

// Запуск службы
func (s *Service) Run() error {
	// Инициализация аргументов приложения
	svcFlag, configFlag := s.initFlags()

	// Инициализация конфига приложения
	if err := s.initConfiguration(configFlag); err != nil {
		return err
	}

	// Инициализация логгера
	s.initLogger()

	// Смотрим наличие менеджеров сервисов в ОС
	// Если хоть что-то есть, то запускаем приложение через системный менеджер сервисов
	// иначе приложение будет работать как обычная программа
	if service.ChosenSystem() != nil {
		log.Info("Service system is available: ", service.AvailableSystems())

		// Инициализация инстанса сервиса
		if err := s.initService(svcFlag); err != nil {
			return err
		}
	} else {
		log.Info("Service system is not found")

		// Инициализация приложения
		s.app.Init(s.configuration)

		// Запускаем обработчик приложения
		s.app.Run()
	}

	return nil
}

// Инициализация конфига приложения
func (s *Service) initConfiguration(configFlag *string) error {
	// Имя файла yml-конфига
	viper.SetConfigName(s.configName)

	// По умолчанию конфиг лежит тамже, где и приложение
	viper.AddConfigPath(*configFlag)

	// by https://stackoverflow.com/questions/18537257/how-to-get-the-directory-of-the-currently-running-file
	folderPath, err := osext.ExecutableFolder()
	if err != nil {
		return err
	}
	viper.AddConfigPath(folderPath)

	// Читаем конфиг-файл
	if err := viper.ReadInConfig(); err != nil {
		return errors.New("Error reading config file, " + err.Error())
	}

	// Парсим конфиг для приложения
	if err := viper.Unmarshal(&s.configuration); err != nil {
		return errors.New("Unable to decode into struct, " + err.Error())
	}

	// Парсим конфиг для ядра
	if err := viper.Unmarshal(&s.coreConfiguration); err != nil {
		return errors.New("Unable to decode into struct, " + err.Error())
	}

	return nil
}

// Инициализация инстанса сервиса
func (s *Service) initService(svcFlag *string) error {
	// Структура программы
	prg := &Program{
		exit:              make(chan struct{}),
		app:               s.app,
		coreConfiguration: s.coreConfiguration,
	}

	// Конфиг сервиса
	svcConfig := &service.Config{
		Name:        s.coreConfiguration.App.Name,
		DisplayName: s.coreConfiguration.App.DisplayName,
		Description: s.coreConfiguration.App.Description,
	}

	// Создание экземпляра сервиса
	srv, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}

	prg.service = srv

	// Инициализация системного логгера
	errs := make(chan error, 5)
	logger, err = srv.Logger(errs)
	if err != nil {
		log.Fatal(err)
	}

	// Вывод лога ошибок в консоль терминала
	go func() {
		for {
			err := <-errs
			if err != nil {
				log.Error(err)
			}
		}
	}()

	// Управление сервисом
	if len(*svcFlag) != 0 {
		err := service.Control(srv, *svcFlag)
		if err != nil {
			log.Info("Valid actions: ", service.ControlAction)
			log.Fatal(err)
		}

		return nil
	}

	// Инициализация приложения
	s.app.Init(s.configuration)

	// Запуск обработчика сервиса
	err = s.Run()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

// Инициализация аргументов приложения
func (s *Service) initFlags() (*string, *string) {
	// Управление системные сервис-менеджером (install/uninstall/start/stop)
	svcFlag := flag.String("service", "", "Control the system service.")

	// Путь к конфиг файлу
	// По умолчанию находится тамже, где и само приложение
	configFlag := flag.String("config", ".", "Path to the config file (default \".\").")

	flag.Parse()

	return svcFlag, configFlag
}

// Инициализация логгера
func (s *Service) initLogger() {
	switch s.coreConfiguration.Logging.Level {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warning":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	case "panic":
		log.SetLevel(log.PanicLevel)
	}
}
