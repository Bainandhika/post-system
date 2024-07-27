package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"post-system/app/logging"
	"post-system/app/services"
	"post-system/lib/models"

	"github.com/gin-gonic/gin"
)

type PostsHandler struct {
	postsService services.PostsService
}

func NewPostsHandler(postsService services.PostsService) *PostsHandler {
	return &PostsHandler{postsService: postsService}
}

func (h *PostsHandler) Insert(c *gin.Context) {
	urlPath := c.Request.URL.Path

	var payload models.AddPost
	if err := c.ShouldBindJSON(&payload); err != nil {
		logging.Error.Println(logDetail(urlPath, payload, err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.postsService.Insert(payload); err != nil {
		logging.Error.Println(logDetail(urlPath, payload, err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "post created"})
}

func (h *PostsHandler) GetAll(c *gin.Context) {
	urlPath := c.Request.URL.Path

	posts, err := h.postsService.GetAll()
	if err != nil {
		logging.Error.Println(logDetail(urlPath, nil, err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (h *PostsHandler) GetById(c *gin.Context) {
	urlPath := c.Request.URL.Path

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logging.Error.Println(logDetail(urlPath, id, err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post, err := h.postsService.GetById(id)
	if err != nil {
		logging.Error.Println(logDetail(urlPath, id, err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if post == nil {
		logging.Error.Println(logDetail(urlPath, id, errors.New("post not found")))
		c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *PostsHandler) Update(c *gin.Context) {
	urlPath := c.Request.URL.Path

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logging.Error.Println(logDetail(urlPath, id, err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var payload models.UpdatePost
	if err := c.ShouldBindJSON(&payload); err != nil {
		logging.Error.Println(logDetail(urlPath, id, err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req := gin.H{
		"id":   id,
        "Data": payload,
	}

	if err := h.postsService.Update(id, payload); err != nil {
		logging.Error.Println(logDetail(urlPath, req, err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "post updated"})
}

func (h *PostsHandler) Delete(c *gin.Context) {
	urlPath := c.Request.URL.Path

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logging.Error.Println(logDetail(urlPath, id, err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.postsService.Delete(id); err != nil {
		logging.Error.Println(logDetail(urlPath, id, err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "post deleted"})
}

func logDetail(urlPath string, request any, err error) string {
	detail := struct {
		URLPath string `json:"url_path"`
		Request any    `json:"request"`
		Error   string `json:"error"`
	}{
		URLPath: urlPath,
		Request: request,
	}

	if err != nil {
		detail.Error = err.Error()
	}

	detailBytes, _ := json.Marshal(detail)
	return string(detailBytes)
}
