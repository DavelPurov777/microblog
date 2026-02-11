package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/DavelPurov777/microblog/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
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
		createPost.GET("/", h.getPosts)
	}

	return router
}