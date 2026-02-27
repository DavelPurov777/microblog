package main

import (
	"fmt"
	"os"

	"net/http"
	_ "net/http/pprof"

	handler "github.com/DavelPurov777/microblog/internal/handlers"
	mylogger "github.com/DavelPurov777/microblog/internal/logger"
	"github.com/DavelPurov777/microblog/internal/queue"
	"github.com/DavelPurov777/microblog/internal/repository"
	"github.com/DavelPurov777/microblog/internal/server"
	"github.com/DavelPurov777/microblog/internal/service"
	"github.com/DavelPurov777/microblog/internal/storage"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	os.Exit(run())
}

func run() int {
	logger := mylogger.NewLogger(100)
	defer logger.Close()

	if err := godotenv.Load(); err != nil {
		logger.Error(fmt.Sprintf("error loading env variables: %s", err.Error()))
		return 1
	}

	if err := initConfig(); err != nil {
		logger.Error(fmt.Sprintf("error initializing configs: %s", err.Error()))
		return 1
	}

	db, err := storage.NewPostgresDB(storage.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logger.Error(fmt.Sprintf("failed to initialize DB: %s", err.Error()))
		return 1
	}
	defer db.Close()

	likeQueue := queue.NewLikeQueue(100)
	salt := viper.GetString("salt")

	repos := repository.NewRepository(db)
	services := service.NewService(repos, likeQueue, salt)
	services.PostsList.StartLikeWorker(logger)
	httpHandler := handler.NewHandler(services, logger)

	if os.Getenv("PPROF_ENABLED") == "true" {
		go func() {
			logger.Info("pprof started on :6060")
			if err := http.ListenAndServe("localhost:6060", nil); err != nil {
				logger.Error(fmt.Sprintf("pprof server error: %v", err))
			}
		}()
	}

	srv := new(server.Server)
	if err := srv.Run(viper.GetString("port"), httpHandler.InitRoutes()); err != nil {
		logger.Error(fmt.Sprintf("error occured while running HTTP server %s", err.Error()))
		return 1
	}

	return 0
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
