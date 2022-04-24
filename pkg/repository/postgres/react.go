package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/luqmansen/blogo/pkg/blogo"
	log "github.com/sirupsen/logrus"
)

type ReactRepository struct {
	db *sqlx.DB
}

func (r PostRepository) InsertUserReact(react *blogo.React) error {
	statement := `
INSERT INTO blogo.public.react_users(author_id, resource_kind, react_id)
VALUES (:author_id, :resource_kind, :resource_kind)
 `

	result, err := r.db.NamedExec(statement, react)
	if err != nil {
		log.Error(err, result)
		return err
	}
	return nil
}
