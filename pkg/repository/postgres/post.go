package postgres

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/luqmansen/blogo/pkg/blogo"
	log "github.com/sirupsen/logrus"
)

type PostRepository struct {
	db *sqlx.DB
}

func (r PostRepository) GetPost(limit, offset int) []*blogo.Post {
	query := `SELECT * FROM blogo.public.posts limit ($1) offset ($2)`
	var posts []*blogo.Post
	err := r.db.Select(&posts, query, limit, offset)
	if err != nil {
		panic(err)
	}

	return posts
}

func (r PostRepository) FindByID(postID uint64) *blogo.Post {
	query := `SELECT * FROM blogo.public.posts where id = ($1) limit 1`
	var post blogo.Post
	err := r.db.Get(&post, query, postID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		} else {
			panic(err)
		}
	}

	return &post
}

func (r PostRepository) InsertPost(post *blogo.Post) error {
	statement := `INSERT INTO blogo.public.posts (author_id, title, content) VALUES (:author_id, :title, :content)`

	result, err := r.db.NamedExec(statement, post)
	if err != nil {
		log.Error(err, result)
		return err
	}
	return nil
}

func NewPostRepository(sqlClient *sqlx.DB) *PostRepository {
	return &PostRepository{
		db: sqlClient,
	}

}
func debugStruct(d interface{}) {
	s, _ := json.MarshalIndent(d, "", "\t")
	fmt.Println(string(s))
}
