package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/luqmansen/blogo/pkg/blogo"
	log "github.com/sirupsen/logrus"
)

type CommentRepository struct {
	db *sqlx.DB
}

func NewCommentRepository(sqlClient *sqlx.DB) *CommentRepository {
	return &CommentRepository{
		db: sqlClient,
	}
}

func (r CommentRepository) InsertComment(comment *blogo.Comment) error {
	statement := `INSERT INTO blogo.public.comments (parent_post_id, parent_id, author_id, content) VALUES ($1, $2, $3, $4)`

	result, err := r.db.Exec(statement, comment.ParentPostID, comment.ParentID, comment.AuthorID, comment.Content)
	if err != nil {
		log.Error(err, result)
		return err
	}
	return nil
}

// GetByID will get a parent comments and all of its child down to N level
func (r CommentRepository) GetByID(commentId uint64) *blogo.Comment {
	query := `
WITH RECURSIVE cte (id, content, author_id, parent_id) as (
    select id,
           content,
           author_id,
           parent_id
    from comments
    where  id = cast(($1) as bigint)

    union all

    select comments.id,
           comments.content,
           comments.author_id,
           comments.parent_id
    from comments
    join cte on comments.parent_id = cte.id
) select id, content, author_id, parent_id from cte
order by id
`
	var comments []*blogo.Comment
	err := r.db.Select(&comments, query, commentId)
	if err != nil {
		panic(err)
	}

	node := make(map[uint64]*blogo.Comment)
	for _, v := range comments {
		node[v.ID] = v
	}
	var rootNode *blogo.Comment
	for _, comm := range node {
		if comm.ParentID == nil {
			rootNode = comm
			continue
		}

		parent := node[uint64(*comm.ParentID)]
		if parent == nil {
			panic(fmt.Sprintf("parent nil %d", *comm.ParentID))
		}
		if parent.Replies == nil {
			parent.Replies = make([]*blogo.Comment, 0)
			parent.Replies = append(parent.Replies, comm)
		} else {
			parent.Replies = append(parent.Replies, comm)
		}
	}

	return rootNode
}

// GetByPostID will get all comments by its post id
func (r CommentRepository) GetByPostID(postId uint64) []*blogo.Comment {
	query := `
WITH RECURSIVE cte (id, content, author_id, parent_id) as (
    select id,
           content,
           author_id,
           parent_id
    from comments
    where  parent_post_id = cast(($1) as bigint)

    union all

    select comments.id,
           comments.content,
           comments.author_id,
           comments.parent_id
    from comments
    join cte on comments.parent_id = cte.id
) select id, content, author_id, parent_id from cte
order by id
`
	var comments []*blogo.Comment
	err := r.db.Select(&comments, query, postId)
	if err != nil {
		panic(err)
	}

	node := make(map[uint64]*blogo.Comment)
	for _, v := range comments {
		node[v.ID] = v
	}

	rootNodes := make([]*blogo.Comment, 0)
	for _, comm := range node {
		if comm.ParentID == nil {
			rootNodes = append(rootNodes, comm)
			continue
		}

		parent := node[uint64(*comm.ParentID)]
		if parent == nil {
			panic(fmt.Sprintf("parent nil %d", *comm.ParentID))
		}
		if parent.Replies == nil {
			parent.Replies = make([]*blogo.Comment, 0)
			parent.Replies = append(parent.Replies, comm)
		} else {
			parent.Replies = append(parent.Replies, comm)
		}
	}

	return rootNodes
}
