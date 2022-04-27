package blogo

import (
	"time"
)

type React struct {
	ID uint64 `db:"id" json:"id"`

	AuthorID uint64 `db:"author_id" json:"author_id"`
	ReactID  int    `db:"react_id" json:"react_id"`

	PostID    *uint64 `db:"post_id" json:"post_id"`
	CommentID *uint64 `db:"comment_id" json:"comment_id"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type ReactViews struct {
	ReactID int `db:"react_id" json:"react_id"`
	Count   int `db:"count" json:"count"`
}

type ReactRepository interface {
	InsertUserReact(react *React) error
	GetByPostID(PostId) []*ReactViews
}

type ReactService struct {
	reactRepository ReactRepository
}

func NewReactService(repository ReactRepository) *ReactService {
	return &ReactService{reactRepository: repository}
}

func (r *ReactService) AddReact(react *React) error {
	react.AuthorID = 1 // TODO: add user
	return r.reactRepository.InsertUserReact(react)
}
