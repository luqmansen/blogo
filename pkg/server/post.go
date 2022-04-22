package server

import (
	"github.com/gin-gonic/gin"
	"github.com/luqmansen/blogo/pkg/blogo"
	"net/http"
)

func (h handler) GetPost(c *gin.Context) {
	posts := h.postService.GetPostMany(10, 0)
	c.JSON(http.StatusOK, posts)
}

func (h handler) CreatePost(c *gin.Context) {
	var postRequest PostRequest
	if err := c.BindJSON(&postRequest); err != nil {
		panic(err)
	}

	var post blogo.Post
	marshal(&postRequest, &post)

	err := h.postService.CreatePost(&post)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, post)
}

func (h handler) FindPostByID(c *gin.Context, postId PostId) {
	p := h.postService.GetPostByID(uint64(postId))
	if p == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	var resp PostResponse
	marshal(&p, &resp)

	c.JSON(http.StatusOK, resp)
}
