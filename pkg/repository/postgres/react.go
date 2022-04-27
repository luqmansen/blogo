package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/luqmansen/blogo/pkg/blogo"
	log "github.com/sirupsen/logrus"
)

type ReactRepository struct {
	db *sqlx.DB
}

func NewReactRepository(sqlClient *sqlx.DB) *ReactRepository {
	return &ReactRepository{
		db: sqlClient,
	}

}

func (r ReactRepository) GetByPostID(postID blogo.PostId) []*blogo.ReactViews {
	query := `
SELECT react_id, count(post_id)
FROM react_users
WHERE post_id = cast(($1) as bigint)
GROUP BY react_id
`

	var reacts []*blogo.ReactViews
	err := r.db.Select(&reacts, query, postID)
	if err != nil {
		panic(err)
	}

	return reacts
}

func getReactByCommentID(db *sqlx.DB, commentIDs []blogo.CommentId) map[blogo.CommentId][]*blogo.ReactViews {
	query := `
SELECT comment_id, react_id, count(comment_id)
from react_users
where comment_id in (?)
group by comment_id, react_id
`
	query, args, err := sqlx.In(query, commentIDs)
	if err != nil {
		panic(err)
	}

	type rv struct {
		CommentID blogo.CommentId `db:"comment_id"`
		blogo.ReactViews
	}
	reactList := make(map[blogo.CommentId][]*blogo.ReactViews)

	query = db.Rebind(query)
	rows, err := db.Queryx(query, args...)
	for rows.Next() {
		var react rv
		err = rows.StructScan(&react)
		if err != nil {
			panic(err)
		}

		v, ok := reactList[react.CommentID]
		if ok {
			v = append(v, &react.ReactViews)
		} else {
			l := make([]*blogo.ReactViews, 0)
			l = append(l, &react.ReactViews)
			reactList[react.CommentID] = l
		}
	}

	return reactList
}

func (r ReactRepository) InsertUserReact(react *blogo.React) error {
	statement := `
INSERT INTO blogo.public.react_users(author_id, post_id, comment_id, react_id)
VALUES (:author_id, :post_id, :comment_id, :react_id)
 `
	result, err := r.db.NamedExec(statement, react)
	if err != nil {
		log.Error(err, result)
		return err
	}
	return nil
}
