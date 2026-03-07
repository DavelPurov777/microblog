package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"net/http"
	_ "net/http/pprof"

	"github.com/DavelPurov777/microblog/configs/config"
	handler "github.com/DavelPurov777/microblog/internal/handlers"
	mylogger "github.com/DavelPurov777/microblog/internal/logger"
	"github.com/DavelPurov777/microblog/internal/queue"
	"github.com/DavelPurov777/microblog/internal/repository"
	"github.com/DavelPurov777/microblog/internal/server"
	"github.com/DavelPurov777/microblog/internal/service"
	"github.com/DavelPurov777/microblog/internal/storage"
	"github.com/joho/godotenv"
)

func main() {
	os.Exit(run())
}

func run() int {
	logger := mylogger.NewLogger(100)
	defer logger.Close()

	if err := godotenv.Load(); err != nil {
		logger.Info(fmt.Sprintf("error loading env variables: %s", err.Error()))
	}

	cfg, err := config.Load()
	if err != nil {
		logger.Error(fmt.Sprintf("error initializing configs: %s", err.Error()))
		return 2
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := storage.NewPgxPool(ctx, cfg.DB, cfg.DBPool)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to initialize NewPxgPool: %s", err.Error()))
		return 3
	}
	defer pool.Close()

	db := storage.NewSQLXFromPgxPool(pool)
	defer db.Close()

	likeQueue := queue.NewLikeQueue(cfg.LikeQueueBuffer)
	salt := cfg.Salt

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
	if err := srv.Run(cfg.Port, httpHandler.InitRoutes()); err != nil {
		logger.Error(fmt.Sprintf("error occured while running HTTP server %s", err.Error()))
		return 4
	}

	return 0
}
