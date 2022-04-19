package server

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/luqmansen/blogo/internal/blogo"
	"github.com/luqmansen/blogo/internal/config"
	"net/http"
)

type handler struct {
	config         *config.Configuration
	oauthConfig    OAuth2ConfigInterface
	postService    *blogo.PostService
	commentService *blogo.CommentService
}

func NewHandler(
	config *config.Configuration,
	postService *blogo.PostService,
	commentService *blogo.CommentService,
	//oauthConfig OAuth2ConfigInterface,
) ServerInterface {
	return handler{
		config:         config,
		postService:    postService,
		commentService: commentService,
		//oauthConfig: oauthConfig,
	}
}

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

func (h handler) CreateUser(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h handler) LoginUser(c *gin.Context, params LoginUserParams) {
	//TODO implement me
	panic("implement me")
}

func (h handler) LogoutUser(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h handler) DeleteUser(c *gin.Context, username string) {
	//TODO implement me
	panic("implement me")
}

func (h handler) GetUserByName(c *gin.Context, username string) {
	//TODO implement me
	panic("implement me")
}

func (h handler) UpdateUser(c *gin.Context, username string) {
	//TODO implement me
	panic("implement me")
}

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
	//TODO implement me
	panic("implement me")
}

func marshal(src, dest interface{}) {
	b, err := json.Marshal(src)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(b, dest)
	if err != nil {
		panic(err)
	}
}

func debugStruct(d interface{}) {
	s, _ := json.MarshalIndent(d, "", "\t")
	fmt.Println(string(s))
}
