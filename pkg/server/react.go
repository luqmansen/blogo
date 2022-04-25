package server

import (
	"github.com/gin-gonic/gin"
	"github.com/luqmansen/blogo/pkg/blogo"
	"net/http"
)

func (h handler) GetReactList(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h handler) AddReact(c *gin.Context) {
	var reactRequest ReactRequest
	if err := c.BindJSON(&reactRequest); err != nil {
		panic(err)
	}

	if reactRequest.PostId != nil && reactRequest.CommentId != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{
			Message: "only post id or comment id allowed, not both",
			Status:  false,
		})
		return
	}

	var react blogo.React
	marshal(&reactRequest, &react)

	err := h.reactService.AddReact(&react)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, ResourceCreatedResponse{
		Message: "react created",
		Status:  true,
	})
}
