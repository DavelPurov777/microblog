package handlers

import (
	"net/http"
	"strconv"

	"github.com/DavelPurov777/microblog/services/engagement/internal/logger"
	"github.com/DavelPurov777/microblog/services/engagement/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	stats  *service.StatsService
	logger logger.Logger
}

func NewHandler(stats *service.StatsService, log logger.Logger) *Handler {
	return &Handler{
		stats:  stats,
		logger: log,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.New()

	stats := r.Group("/stats")
	{
		stats.GET("/posts/:id", h.getPostLikes)
	}

	return r
}

func (h *Handler) getPostLikes(c *gin.Context) {
	idStr := c.Param("id")
	postID, err := strconv.Atoi(idStr)
	if err != nil {
		h.logger.Error("invalid post id: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post id"})
		return
	}
	likes, err := h.stats.GetPostLikes(postID)
	if err != nil {
		h.logger.Error("get likes error: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	h.logger.Info("stats returned for post " + idStr)
	c.JSON(http.StatusOK, gin.H{
		"post_id": postID,
		"likes":   likes,
	})
}
