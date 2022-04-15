package postgres

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/luqmansen/blogo/internal/blogo"
	log "github.com/sirupsen/logrus"
)

type PostRepository struct {
	db           *sqlx.DB
	databaseName string
}

func (r PostRepository) FindByID(postID uint64) error {
	//TODO implement me
	panic("implement me")
}

func (r PostRepository) InsertPost(post *blogo.Post) error {
	//TODO implement me
	panic("implement me")
}

func NewPostRepository(sqlClient *sqlx.DB) *PostRepository {
	return &PostRepository{
		db: sqlClient,
	}
}

func (r PostRepository) CreatePost(post *blogo.Post) error {
	statement := `INSERT INTO spost (content,title) VALUES (?,?)`
	result, err := r.db.NamedExec(statement, post)
	if err != nil {
		log.Error(err, result)
	}
	return nil
}
