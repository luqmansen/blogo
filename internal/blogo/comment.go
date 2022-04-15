package blogo

import (
	"time"
)

type CommentRequest struct {
	Content string `json:"content"`
}

type Comment struct {
	ID uint64 `json:"id"`

	ParentPostID uint64 `json:"parent_post_id"`
	ParentID     uint64 `json:"parent_id"`
	AuthorID     string `json:"author_id"`
	Content      string `json:"content"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CommentRepository interface {
	FindByID(commentId uint64) *Comment
	InsertForPost(postId uint64, commentBody *Comment) error
	InsertForComment(commentId uint64, commentBody *Comment) error
}

type CommentService struct {
	commentRepo CommentRepository
}

func NewCommentService(repo CommentRepository) *CommentService {
	return &CommentService{commentRepo: repo}
}

func (c CommentService) GetCommentByID(commentId uint64) *Comment {
	panic("implement me")
	return nil
}

func (c CommentService) CreateCommentForPost(postId uint64, commentBody *CommentRequest) error {
	panic("implement me")
	return nil

}

func (c CommentService) CreateCommentForComment(commentId uint64, commentBody *CommentRequest) error {
	panic("implement me")
	return nil
}
