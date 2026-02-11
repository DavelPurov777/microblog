package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/DavelPurov777/microblog/internal/models"	
)

func (h *Handler)  createPost(c *gin.Context) {
	var input models.Post
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	
	id, err := h.services.PostsList.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}