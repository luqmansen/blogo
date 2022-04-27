package postgres

import (
	"database/sql"
	"errors"
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

func isCommentExists(db *sqlx.DB, id blogo.CommentId) bool {
	qry := `SELECT * FROM comments WHERE id = ($1)`

	var c blogo.Comment
	err := db.Get(&c, qry, id)
	debugStruct(c)
	fmt.Println(err)
	if err != nil && err == sql.ErrNoRows {
		return false
	}
	return true
}

func (r CommentRepository) InsertComment(comment *blogo.Comment) error {

	if !isCommentExists(r.db, blogo.CommentId(*comment.ParentID)) {
		return errors.New("parent id does not exists")
	}

	statement := `INSERT INTO blogo.public.comments (parent_post_id, parent_id, author_id, content) VALUES ($1, $2, $3, $4)`

	result, err := r.db.Exec(statement, comment.ParentPostID, comment.ParentID, comment.AuthorID, comment.Content)
	if err != nil {
		log.Error(err, result)
		return err
	}
	return nil
}

// GetByID will get a parent comments and all of its child down to N level
func (r CommentRepository) GetByID(commentId blogo.CommentId) *blogo.Comment {
	query := `
WITH RECURSIVE cte (id, content, author_id,parent_id, username) as (
    select comments.id,
           comments.content,
           comments.author_id,
           comments.parent_id,
			users.username       
    from comments
    join users on comments.author_id = users.id
    where  comments.id = cast(($1) as bigint)

    union all

    select comments.id,
           comments.content,
           comments.author_id,
           comments.parent_id,
			users.username       
    from comments
	join users on comments.author_id = users.id
    join cte on comments.parent_id = cte.id
) select id, content, author_id, parent_id, username from cte
order by id
`
	var comments []*blogo.Comment
	err := r.db.Select(&comments, query, commentId)
	if err != nil {
		panic(err)
	}

	node := make(map[blogo.CommentId]*blogo.Comment)
	for _, v := range comments {
		node[v.ID] = v
	}

	var rootNode *blogo.Comment
	// assign replies to each comment
	for _, comm := range node {
		if comm.ID == commentId {
			rootNode = comm
			continue
		}

		parent := node[blogo.CommentId(*comm.ParentID)]
		if parent == nil {
			continue
		}
		if parent.Replies == nil {
			parent.Replies = make([]*blogo.Comment, 0)
			parent.Replies = append(parent.Replies, comm)
		} else {
			parent.Replies = append(parent.Replies, comm)
		}
	}

	commentIDs := make([]blogo.CommentId, len(comments))
	for idx, v := range comments {
		commentIDs[idx] = v.ID
	}

	reactList := getReactByCommentID(r.db, commentIDs)
	// assign react  to each comment
	for _, comm := range node {
		rv, ok := reactList[comm.ID]
		if ok {
			comm.ReactViews = rv
		}
	}

	return rootNode
}

// GetByPostID will get all comments by its post id
func (r CommentRepository) GetByPostID(postId blogo.PostId) []*blogo.Comment {
	query := `
WITH RECURSIVE cte (id, content, author_id,parent_id, username) as (
    select comments.id,
           comments.content,
           comments.author_id,
           comments.parent_id,
			users.username       
    from comments
    join users on comments.author_id = users.id
    where  comments.parent_post_id = cast(($1) as bigint)

    union all

    select comments.id,
           comments.content,
           comments.author_id,
           comments.parent_id,
			users.username       
    from comments
	join users on comments.author_id = users.id
    join cte on comments.parent_id = cte.id
) select id, content, author_id, parent_id, username from cte
order by id
`
	var comments []*blogo.Comment
	err := r.db.Select(&comments, query, postId)
	if err != nil {
		panic(err)
	}

	node := make(map[blogo.CommentId]*blogo.Comment)
	for _, v := range comments {
		node[v.ID] = v
	}

	rootNodes := make([]*blogo.Comment, 0)
	for _, comm := range node {
		if comm.ParentID == nil {
			rootNodes = append(rootNodes, comm)
			continue
		}

		parent := node[blogo.CommentId(*comm.ParentID)]
		if parent == nil {
			continue
		}
		if parent.Replies == nil {
			parent.Replies = make([]*blogo.Comment, 0)
			parent.Replies = append(parent.Replies, comm)
		} else {
			parent.Replies = append(parent.Replies, comm)
		}
	}

	commentIDs := make([]blogo.CommentId, len(comments))
	for idx, v := range comments {
		commentIDs[idx] = v.ID
	}

	reactList := getReactByCommentID(r.db, commentIDs)
	// I separate the loop for the sake of code clarity,
	// the algorithm complexity should be the same
	for _, comm := range node {
		rv, ok := reactList[comm.ID]
		if ok {
			// assign react to each comment
			comm.ReactViews = rv
		}
	}

	return rootNodes
}
