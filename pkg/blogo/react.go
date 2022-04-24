package blogo

import (
	"time"
)

type ReactResourceKind int

const (
	ReactResourceKindPost    ReactResourceKind = 1
	ReactResourceKindComment ReactResourceKind = 2
)

type React struct {
	ID uint64 `db:"id" json:"id"`

	AuthorID     uint64            `db:"author_id" json:"author_id"`
	ResourceKind ReactResourceKind `db:"resource_kind" json:"resource_kind"`
	ReactID      int               `db:"react_id" json:"react_id"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type ReactRepository interface {
	InsertUserReact(react *React) error
}

type ReactService struct {
	reactRepository ReactRepository
}

func NewReactService(repository ReactRepository) *ReactService {
	return &ReactService{reactRepository: repository}
}

func (r *ReactService) AddReact(react *React) error {
	return r.reactRepository.InsertUserReact(react)
}
