package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/DavelPurov777/microblog/configs/config"
	"github.com/DavelPurov777/microblog/services/engagement/internal/consumer"
	"github.com/DavelPurov777/microblog/services/engagement/internal/handlers"
	"github.com/DavelPurov777/microblog/services/engagement/internal/logger"
	"github.com/DavelPurov777/microblog/services/engagement/internal/repository"
	"github.com/DavelPurov777/microblog/services/engagement/internal/service"
	"github.com/DavelPurov777/microblog/services/engagement/internal/storage"
)

func main() {
	os.Exit(run())
}

func run() int {
	log := logger.NewLogger()

	cfg, err := config.Load()
	if err != nil {
		log.Error(fmt.Sprintf("error initializing configs: %s", err.Error()))
		return 2
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// используем те же настройки подключения к БД, что и api
	pool, err := storage.NewPgxPool(ctx, cfg.DB, cfg.DBPool)
	if err != nil {
		log.Error(fmt.Sprintf("failed to initialize NewPxgPool: %s", err.Error()))
		return 3
	}
	defer pool.Close()

	db := storage.NewSQLXFromPgxPool(pool)
	defer db.Close()

	repo := repository.NewPostLikesRepo(db)
	statsService := service.NewStatsService(repo)

	likeConsumer := consumer.NewLikeConsumer(
		[]string{"kafka:9092"},
		"likes",
		"engagement-service",
		repo,
	)

	// запускаем консьюмера
	go func() {
		if err := likeConsumer.Run(context.Background()); err != nil {
			log.Error(fmt.Sprintf("consumer error: %v", err))
		}
	}()

	h := handlers.NewHandler(statsService, log)
	router := h.InitRoutes()

	port := "8082"

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	log.Info("engagement service started on :" + port)

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Error(fmt.Sprintf("HTTP server error: %v", err))
		return 4
	}

	return 0
}
