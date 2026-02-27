package handler

import (
	"fmt"
	"net/http"

	"github.com/DavelPurov777/microblog/internal/models"
	"github.com/gin-gonic/gin"
)

type registerInput struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) register(c *gin.Context) {
	var input registerInput

	if err := c.BindJSON(&input); err != nil {
		h.logger.Error(fmt.Sprintf("registration failed due to %v", err))
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user := models.User{
		Name:     input.Name,
		Username: input.Username,
		Password: input.Password,
	}

	id, err := h.services.Authorization.CreateUser(user)
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
