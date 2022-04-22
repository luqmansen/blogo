package server

import (
	"github.com/gin-gonic/gin"
	"github.com/luqmansen/blogo/pkg/blogo"
	"net/http"
)

func (h handler) AddComment(c *gin.Context) {

	var req CommentRequest
	if err := c.BindJSON(&req); err != nil {
		panic(err)
	}

	comment := &blogo.Comment{
		ParentPostID: uint64(req.ParentPostId),
		ParentID:     req.ParentId,
		AuthorID:     uint64(1),
		Content:      req.Content,
	}

	err := h.commentService.CreateComment(comment)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, comment)
}

func (h handler) GetCommentByID(c *gin.Context, commentId CommentId) {
	posts := h.commentService.GetCommentByID(uint64(commentId))

	var resp CommentResponse
	marshal(posts, &resp)
	c.JSON(http.StatusOK, resp)
}
