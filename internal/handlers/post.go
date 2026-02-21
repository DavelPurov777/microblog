package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/DavelPurov777/microblog/internal/models"
	"github.com/gin-gonic/gin"
)

type getAllPostsResponse struct {
	Data []models.Post `json:"data"`
}

func (h *Handler) createPost(c *gin.Context) {
	var input models.Post
	if err := c.BindJSON(&input); err != nil {
		h.logger.Error(fmt.Sprintf("create post: invalid input: %v", err))
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.PostsList.Create(input)
	if err != nil {
		h.logger.Error(fmt.Sprintf("create post failed: %v", err))
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.logger.Info(fmt.Sprintf("post created with id=%d", id))
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllPosts(c *gin.Context) {
	posts, err := h.services.PostsList.GetAll()
	if err != nil {
		h.logger.Error(fmt.Sprintf("getting all posts failed due to %v", err))
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.logger.Info("posts sended")
	c.JSON(http.StatusOK, getAllPostsResponse{
		Data: posts,
	})
}

func (h *Handler) likePost(c *gin.Context) {
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error(fmt.Sprintf("like post: invalid id: %v", err))
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	if err := h.services.PostsList.LikePost(listId); err != nil {
		h.logger.Error(fmt.Sprintf("like to post has failed due to %v", err))
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.logger.Info("post have 1 more like")
	c.JSON(http.StatusOK, statusResponse{"ok"})
}
