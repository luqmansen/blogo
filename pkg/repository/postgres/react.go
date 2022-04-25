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

func (r ReactRepository) GetByPostID(postID uint64) []*blogo.ReactViews {
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
