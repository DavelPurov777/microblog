package handler

import (
	"github.com/DavelPurov777/microblog/internal/logger"
	"github.com/DavelPurov777/microblog/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
	logger   logger.Logger
}

func NewHandler(services *service.Service, log logger.Logger) *Handler {
	return &Handler{
		services: services,
		logger:   log,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	register := router.Group("/register")
	{
		register.POST("/", h.register)
	}
	createPost := router.Group("/posts")
	{
		createPost.POST("/", h.createPost)
		createPost.GET("/", h.getAllPosts)
		createPost.PUT("/:id/like", h.likePost)
	}

	return router
}
