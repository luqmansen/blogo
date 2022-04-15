package blogo

import (
	"time"
)

type PostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Post struct {
	ID uint64 `json:"id"`

	AuthorID uint64
	Title    string
	Content  string

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostRepository interface {
	FindByID(postID uint64) error
	InsertPost(post *Post) error
}

type PostService struct {
	db PostRepository
}

func NewPostService(db PostRepository) PostService {
	return PostService{db: db}
}

func (p PostService) CreatePost(payload *PostRequest) error {

	post := &Post{
		AuthorID: 0,
		Title:    payload.Title,
		Content:  payload.Content,
	}

	return p.db.InsertPost(post)
}

func (p PostService) GetPostMany(limit, pageNum int) []*Post {
	panic("implement me")
	return nil
}

func (p PostService) GetPostById(postId uint64) *Post {
	panic("implement me")
	return nil
}
