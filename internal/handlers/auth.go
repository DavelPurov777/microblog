package handler

import (
	"fmt"
	"net/http"

	"github.com/DavelPurov777/microblog/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) register(c *gin.Context) {
	var input models.User

	if err := c.BindJSON(&input); err != nil {
		h.logger.Error(fmt.Sprintf("registration failed due to %v", err))
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		h.logger.Error(fmt.Sprintf("registration failed due to %v", err))
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.logger.Info("successful registration")
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
