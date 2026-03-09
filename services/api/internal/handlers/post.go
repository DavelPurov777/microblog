package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/DavelPurov777/microblog/services/api/internal/models"
	"github.com/gin-gonic/gin"
)

type createPostRequest struct {
	UserId      int    `json:"user_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type getAllPostsResponse struct {
	Data []models.Post `json:"data"`
}

type likePostRequest struct {
	UserId int `json:"user_id" binding:"required"`
}

func (h *Handler) createPost(c *gin.Context) {
	var input createPostRequest
	if err := c.BindJSON(&input); err != nil {
		h.logger.Error(fmt.Sprintf("create post: invalid input: %v", err))
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	post := models.Post{
		UserId:      input.UserId,
		Title:       input.Title,
		Description: input.Description,
		CreatedAt:   time.Now(),
		Likes:       0,
	}

	id, err := h.services.PostsList.Create(post)
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
	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error(fmt.Sprintf("like post: invalid id: %v", err))
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input likePostRequest
	if err := c.BindJSON(&input); err != nil {
		h.logger.Error(fmt.Sprintf("missing userId in body: %v", err))
		newErrorResponse(c, http.StatusBadRequest, "missing userId in body")
		return
	}

	if err := h.services.PostsList.LikePost(postId, input.UserId); err != nil {
		h.logger.Error(fmt.Sprintf("like to post has failed due to %v", err))
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.logger.Info("post have 1 more like")
	c.JSON(http.StatusOK, statusResponse{"ok"})
}
