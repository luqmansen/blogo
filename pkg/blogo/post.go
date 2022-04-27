package blogo

import (
	"encoding/json"
	"fmt"
	"time"
)

type PostId uint64

type Post struct {
	ID PostId `db:"id" json:"id"`

	AuthorID       uint64 `db:"author_id" json:"author_id"`
	AuthorUsername string `db:"username" json:"author_username"`

	Title      string        `db:"title" json:"title"` // # TODO add field validation before insert
	Content    string        `db:"content" json:"content"`
	ReactViews []*ReactViews `json:"react_views"`
	Replies    []*Comment    `json:"replies"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type PostRepository interface {
	FindByID(PostId) *Post
	InsertPost(post *Post) error
	GetPost(limit, offset int) []*Post
}

type PostService struct {
	postRepository    PostRepository
	commentRepository CommentRepository
	reactRepository   ReactRepository
}

func NewPostService(postRepo PostRepository, commentRepo CommentRepository, reactRepo ReactRepository) *PostService {
	return &PostService{
		postRepository:    postRepo,
		commentRepository: commentRepo,
		reactRepository:   reactRepo,
	}
}

func (p PostService) CreatePost(payload *Post) error {

	payload.AuthorID = 1

	return p.postRepository.InsertPost(payload)
}

func (p PostService) GetPostMany(limit, pageNum int) []*Post {
	return p.postRepository.GetPost(limit, pageNum)
}

func (p PostService) GetPostByID(postID PostId) *Post {
	post := p.postRepository.FindByID(postID)
	if post == nil {
		return nil
	}

	comments := p.commentRepository.GetByPostID(postID)
	post.Replies = comments

	reactViews := p.reactRepository.GetByPostID(postID)
	post.ReactViews = reactViews

	return post
}

func debugStruct(d interface{}) {
	s, _ := json.MarshalIndent(d, "", "\t")
	fmt.Println(string(s))
}
