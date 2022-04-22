package blogo

import (
	"encoding/json"
	"fmt"
	"time"
)

type Post struct {
	ID uint64 `db:"id" json:"id"`

	AuthorID       uint64 `db:"author_id" json:"author_id"`
	AuthorUsername string `db:"username" json:"author_username"`

	Title   string     `db:"title" json:"title"`
	Content string     `db:"content" json:"content"`
	Replies []*Comment `json:"replies"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type PostRepository interface {
	FindByID(postID uint64) *Post
	InsertPost(post *Post) error
	GetPost(limit, offset int) []*Post
}

type PostService struct {
	postRepository    PostRepository
	commentRepository CommentRepository
}

func NewPostService(postRepo PostRepository, commentRepo CommentRepository) *PostService {
	return &PostService{
		postRepository:    postRepo,
		commentRepository: commentRepo,
	}
}

func (p PostService) CreatePost(payload *Post) error {

	post := &Post{
		AuthorID: 1,
		Title:    payload.Title,
		Content:  payload.Content,
	}

	return p.postRepository.InsertPost(post)
}

func (p PostService) GetPostMany(limit, pageNum int) []*Post {
	return p.postRepository.GetPost(limit, pageNum)
}

func (p PostService) GetPostByID(postId uint64) *Post {
	post := p.postRepository.FindByID(postId)
	if post == nil {
		return nil
	}

	comments := p.commentRepository.GetByPostID(postId)
	post.Replies = comments

	return post
}

func debugStruct(d interface{}) {
	s, _ := json.MarshalIndent(d, "", "\t")
	fmt.Println(string(s))
}
