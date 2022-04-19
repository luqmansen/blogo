package blogo

import (
	"time"
)

type Post struct {
	ID uint64 `db:"id" json:"id"`

	AuthorID uint64 `db:"author_id" json:"author_id"`
	Title    string `db:"title" json:"title"`
	Content  string `db:"content" json:"content"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type PostRepository interface {
	FindByID(postID uint64) error
	InsertPost(post *Post) error
	GetPost(limit, offset int) []*Post
}

type PostService struct {
	db PostRepository
}

func NewPostService(db PostRepository) *PostService {
	return &PostService{db: db}
}

func (p PostService) CreatePost(payload *Post) error {

	post := &Post{
		AuthorID: 1,
		Title:    payload.Title,
		Content:  payload.Content,
	}

	return p.db.InsertPost(post)
}

func (p PostService) GetPostMany(limit, pageNum int) []*Post {
	return p.db.GetPost(limit, pageNum)
}

func (p PostService) GetPostById(postId uint64) *Post {
	panic("implement me")
	return nil
}
