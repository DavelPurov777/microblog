package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/DavelPurov777/todo-app-golang/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.default()
	register := router.POST("/register", h.register)

	return router
}