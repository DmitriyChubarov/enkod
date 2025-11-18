package main

import (
	"os"

	"github.com/gocraft/dbr/v2"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"

	"github.com/DmitriyChubarov/enkod/internal/config"
	"github.com/DmitriyChubarov/enkod/internal/http"
	"github.com/DmitriyChubarov/enkod/internal/logic"
	repositorypostgres "github.com/DmitriyChubarov/enkod/internal/repository_postgres"
)

func main() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	log.Info("Запуск сервиса")

	cfg := config.LoadConfig()
	log.WithField("dsn", cfg.PostgresDSN).Debug("Подключение к базе данных")

	conn, err := dbr.Open("pgx", cfg.PostgresDSN, nil)
	if err != nil {
		log.Fatal("Ошибка при открытии соединения с базой данных: ", err)
	}
	defer conn.Close()
	log.Info("Соединение с базой данных успешно установлено")

	if err := conn.Ping(); err != nil {
		log.Fatal("Ошибка при проверке соединения с базой данных: ", err)
	}
	log.Info("Соединение с базой данных успешно проверено")

	session := conn.NewSession(nil)

	repo := repositorypostgres.NewPersonRepository(session)
	service := logic.NewPersonService(repo)
	handler := http.NewPersonHandler(service)

	e := echo.New()
	handler.Register(e)

	log.Infof("Сервис запущен на порту %s", cfg.HTTPPort)
	e.Logger.Fatal(e.Start(":" + cfg.HTTPPort))
	log.Info("Сервис остановлен")
}
