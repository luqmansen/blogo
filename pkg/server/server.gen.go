// Package server provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.1 DO NOT EDIT.
package server

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/gin-gonic/gin"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (POST /api/v1/comment)
	AddComment(c *gin.Context)

	// (GET /api/v1/comment/{commentId})
	GetCommentByID(c *gin.Context, commentId CommentId)
	// get multiple post
	// (GET /api/v1/post)
	GetPost(c *gin.Context)
	// Create a new post
	// (POST /api/v1/post)
	CreatePost(c *gin.Context)
	// Finds Post by ID
	// (GET /api/v1/post/{postId})
	FindPostByID(c *gin.Context, postId PostId)

	// (GET /api/v1/react/)
	GetReactList(c *gin.Context)

	// (POST /api/v1/react/)
	AddReact(c *gin.Context)
	// Create user
	// (POST /user)
	CreateUser(c *gin.Context)
	// Logs user into the system
	// (GET /user/login)
	LoginUser(c *gin.Context, params LoginUserParams)
	// Logs out current logged in user session
	// (GET /user/logout)
	LogoutUser(c *gin.Context)
	// Delete user
	// (DELETE /user/{username})
	DeleteUser(c *gin.Context, username string)
	// Get user by user name
	// (GET /user/{username})
	GetUserByName(c *gin.Context, username string)
	// Update user
	// (PUT /user/{username})
	UpdateUser(c *gin.Context, username string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
}

type MiddlewareFunc func(c *gin.Context)

// AddComment operation middleware
func (siw *ServerInterfaceWrapper) AddComment(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.AddComment(c)
}

// GetCommentByID operation middleware
func (siw *ServerInterfaceWrapper) GetCommentByID(c *gin.Context) {

	var err error

	// ------------- Path parameter "commentId" -------------
	var commentId CommentId

	err = runtime.BindStyledParameter("simple", false, "commentId", c.Param("commentId"), &commentId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter commentId: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetCommentByID(c, commentId)
}

// GetPost operation middleware
func (siw *ServerInterfaceWrapper) GetPost(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetPost(c)
}

// CreatePost operation middleware
func (siw *ServerInterfaceWrapper) CreatePost(c *gin.Context) {

	c.Set(Petstore_authScopes, []string{"write:pets", "read:pets"})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.CreatePost(c)
}

// FindPostByID operation middleware
func (siw *ServerInterfaceWrapper) FindPostByID(c *gin.Context) {

	var err error

	// ------------- Path parameter "postId" -------------
	var postId PostId

	err = runtime.BindStyledParameter("simple", false, "postId", c.Param("postId"), &postId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter postId: %s", err)})
		return
	}

	c.Set(Petstore_authScopes, []string{"write:pets", "read:pets"})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.FindPostByID(c, postId)
}

// GetReactList operation middleware
func (siw *ServerInterfaceWrapper) GetReactList(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetReactList(c)
}

// AddReact operation middleware
func (siw *ServerInterfaceWrapper) AddReact(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.AddReact(c)
}

// CreateUser operation middleware
func (siw *ServerInterfaceWrapper) CreateUser(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.CreateUser(c)
}

// LoginUser operation middleware
func (siw *ServerInterfaceWrapper) LoginUser(c *gin.Context) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params LoginUserParams

	// ------------- Optional query parameter "username" -------------
	if paramValue := c.Query("username"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "username", c.Request.URL.Query(), &params.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter username: %s", err)})
		return
	}

	// ------------- Optional query parameter "password" -------------
	if paramValue := c.Query("password"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "password", c.Request.URL.Query(), &params.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter password: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.LoginUser(c, params)
}

// LogoutUser operation middleware
func (siw *ServerInterfaceWrapper) LogoutUser(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.LogoutUser(c)
}

// DeleteUser operation middleware
func (siw *ServerInterfaceWrapper) DeleteUser(c *gin.Context) {

	var err error

	// ------------- Path parameter "username" -------------
	var username string

	err = runtime.BindStyledParameter("simple", false, "username", c.Param("username"), &username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter username: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.DeleteUser(c, username)
}

// GetUserByName operation middleware
func (siw *ServerInterfaceWrapper) GetUserByName(c *gin.Context) {

	var err error

	// ------------- Path parameter "username" -------------
	var username string

	err = runtime.BindStyledParameter("simple", false, "username", c.Param("username"), &username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter username: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetUserByName(c, username)
}

// UpdateUser operation middleware
func (siw *ServerInterfaceWrapper) UpdateUser(c *gin.Context) {

	var err error

	// ------------- Path parameter "username" -------------
	var username string

	err = runtime.BindStyledParameter("simple", false, "username", c.Param("username"), &username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter username: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.UpdateUser(c, username)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL     string
	Middlewares []MiddlewareFunc
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router *gin.Engine, si ServerInterface) *gin.Engine {
	return RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router *gin.Engine, si ServerInterface, options GinServerOptions) *gin.Engine {
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
	}

	router.POST(options.BaseURL+"/api/v1/comment", wrapper.AddComment)

	router.GET(options.BaseURL+"/api/v1/comment/:commentId", wrapper.GetCommentByID)

	router.GET(options.BaseURL+"/api/v1/post", wrapper.GetPost)

	router.POST(options.BaseURL+"/api/v1/post", wrapper.CreatePost)

	router.GET(options.BaseURL+"/api/v1/post/:postId", wrapper.FindPostByID)

	router.GET(options.BaseURL+"/api/v1/react/", wrapper.GetReactList)

	router.POST(options.BaseURL+"/api/v1/react/", wrapper.AddReact)

	router.POST(options.BaseURL+"/user", wrapper.CreateUser)

	router.GET(options.BaseURL+"/user/login", wrapper.LoginUser)

	router.GET(options.BaseURL+"/user/logout", wrapper.LogoutUser)

	router.DELETE(options.BaseURL+"/user/:username", wrapper.DeleteUser)

	router.GET(options.BaseURL+"/user/:username", wrapper.GetUserByName)

	router.PUT(options.BaseURL+"/user/:username", wrapper.UpdateUser)

	return router
}
