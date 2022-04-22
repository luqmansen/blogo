package blogo

import "time"

type Comment struct {
	ID        uint64    `db:"id" json:"id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`

	ParentPostID uint64 `db:"parent_post_id" json:"parent_post_id"`
	ParentID     *int   `db:"parent_id" json:"parent_id"` // allow nullable on lvl 0 comment

	AuthorID       uint64  `db:"author_id" json:"author_id"`
	AuthorUsername *string `db:"username" json:"author_username"`

	Content string     `db:"content" json:"content"`
	Replies []*Comment `json:"replies"`
}

type CommentRepository interface {
	// GetByPostId get comment which bind to a post
	GetByPostID(commentId uint64) []*Comment
	GetByID(commentId uint64) *Comment
	InsertComment(comment *Comment) error
}

type CommentService struct {
	commentRepo CommentRepository
}

func NewCommentService(repo CommentRepository) *CommentService {
	return &CommentService{commentRepo: repo}
}

func (c CommentService) GetCommentByID(commentId uint64) *Comment {
	return c.commentRepo.GetByID(commentId)
}

func (c CommentService) CreateComment(comment *Comment) error {
	return c.commentRepo.InsertComment(comment)
}
