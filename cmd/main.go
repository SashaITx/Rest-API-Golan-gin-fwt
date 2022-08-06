package main

import (
	rest "Rest_API_Golan-gin-fwt"
	"Rest_API_Golan-gin-fwt/internal/handler"
	"Rest_API_Golan-gin-fwt/internal/repository"
	"Rest_API_Golan-gin-fwt/internal/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "qwerty",
		DBName:   "postgres",
		SSLMode:  "disable",
	})
	if err != nil {
		logrus.Fatal("failed creating DB: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(rest.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		logrus.Fatal("error while running http server: %s", err.Error())
	}
}
