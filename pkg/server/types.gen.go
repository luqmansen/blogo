// Package server provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.1 DO NOT EDIT.
package server

const (
	Petstore_authScopes = "petstore_auth.Scopes"
)

// CommentRequest defines model for CommentRequest.
type CommentRequest struct {
	Content      string `json:"content"`
	ParentId     *int   `json:"parent_id,omitempty"`
	ParentPostId int    `json:"parent_post_id"`
}

// CommentResponse defines model for CommentResponse.
type CommentResponse struct {
	AuthorId       *int64  `json:"author_id,omitempty"`
	AuthorUsername *string `json:"author_username,omitempty"`
	Content        *string `json:"content,omitempty"`
	Id             *int64  `json:"id,omitempty"`

	// replies of this comment
	Replies *[]CommentResponse `json:"replies,omitempty"`
}

// PostRequest defines model for PostRequest.
type PostRequest struct {
	Content string `json:"content"`
	Title   string `json:"title"`
}

// PostResponse defines model for PostResponse.
type PostResponse struct {
	AuthorId       *int64  `json:"author_id,omitempty"`
	AuthorUsername *string `json:"author_username,omitempty"`
	Content        *string `json:"content,omitempty"`
	Id             *int64  `json:"id,omitempty"`
	Name           *string `json:"name,omitempty"`

	// replies of the post
	Replies *[]CommentResponse `json:"replies,omitempty"`
}

// Response for multiple post
type PostsResponse []PostResponse

// User defines model for User.
type User struct {
	Email     *string `json:"email,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	Id        *int64  `json:"id,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
	Password  *string `json:"password,omitempty"`
	Phone     *string `json:"phone,omitempty"`

	// User Status
	UserStatus *int32  `json:"userStatus,omitempty"`
	Username   *string `json:"username,omitempty"`
}

// CommentId defines model for commentId.
type CommentId int

// PostId defines model for postId.
type PostId int64

// AddCommentJSONBody defines parameters for AddComment.
type AddCommentJSONBody CommentRequest

// CreatePostJSONBody defines parameters for CreatePost.
type CreatePostJSONBody PostRequest

// CreateUserJSONBody defines parameters for CreateUser.
type CreateUserJSONBody User

// LoginUserParams defines parameters for LoginUser.
type LoginUserParams struct {
	// The user name for login
	Username *string `json:"username,omitempty"`

	// The password for login in clear text
	Password *string `json:"password,omitempty"`
}

// UpdateUserJSONBody defines parameters for UpdateUser.
type UpdateUserJSONBody User

// AddCommentJSONRequestBody defines body for AddComment for application/json ContentType.
type AddCommentJSONRequestBody AddCommentJSONBody

// CreatePostJSONRequestBody defines body for CreatePost for application/json ContentType.
type CreatePostJSONRequestBody CreatePostJSONBody

// CreateUserJSONRequestBody defines body for CreateUser for application/json ContentType.
type CreateUserJSONRequestBody CreateUserJSONBody

// UpdateUserJSONRequestBody defines body for UpdateUser for application/json ContentType.
type UpdateUserJSONRequestBody UpdateUserJSONBody
