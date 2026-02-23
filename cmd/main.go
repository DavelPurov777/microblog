package main

import (
	"fmt"
	"os"

	handler "github.com/DavelPurov777/microblog/internal/handlers"
	mylogger "github.com/DavelPurov777/microblog/internal/logger"
	"github.com/DavelPurov777/microblog/internal/queue"
	"github.com/DavelPurov777/microblog/internal/repository"
	"github.com/DavelPurov777/microblog/internal/server"
	"github.com/DavelPurov777/microblog/internal/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	logger := mylogger.NewLogger(100)
	defer logger.Close()

	if err := initConfig(); err != nil {
		logger.Error(fmt.Sprintf("error initializing configs: %s", err.Error()))
		os.Exit(1)
	}

	if err := godotenv.Load(); err != nil {
		logger.Error(fmt.Sprintf("error loading env variables: %s", err.Error()))
		os.Exit(1)
	}

	db, err := service.NewPostgresDB(service.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logger.Error(fmt.Sprintf("failed to initialize DB: %s", err.Error()))
		os.Exit(1)
	}

	likeQueue := queue.NewLikeQueue(100)
	repos := repository.NewRepository(db)
	services := service.NewService(repos, likeQueue)
	handlers := handler.NewHandler(services, logger)

	likeQueue.Start(func(id int) {
		err := services.PostsList.ProcessLike(id)
		if err != nil {
			logger.Error(err.Error())
		}
	})

	srv := new(server.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logger.Error(fmt.Sprintf("error occured while running HTTP server %s", err.Error()))
		os.Exit(1)
	}
	logger.Info("Todo App started")
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
