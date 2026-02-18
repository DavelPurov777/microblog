package main

import (
	"os"

	handler "github.com/DavelPurov777/microblog/internal/handlers"
	"github.com/DavelPurov777/microblog/internal/queue"
	"github.com/DavelPurov777/microblog/internal/repository"
	"github.com/DavelPurov777/microblog/internal/server"
	"github.com/DavelPurov777/microblog/internal/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables %s", err.Error())
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
		logrus.Fatalf("failed to initialize DB: %s", err.Error())
	}

	likeQueue := queue.NewLikeQueue(100)
	repos := repository.NewRepository(db)
	services := service.NewService(repos, likeQueue)
	handlers := handler.NewHandler(services)

	likeQueue.Start(func(id int) {
		err := repos.PostsList.LikePost(id)
		if err != nil {
			logrus.Errorf("like error: %v", err)
		}
	})

	srv := new(server.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running HTTP server %s", err.Error())
	}
	logrus.Print("Todo App started")
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
