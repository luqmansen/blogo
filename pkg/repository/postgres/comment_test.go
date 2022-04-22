package postgres

import (
	"encoding/json"
	_ "github.com/lib/pq"
	"testing"
	"time"
)

func TestCommentTree(t *testing.T) {
	var listRaw = `[{"id":5,"parentPostID":0,"parentID":null,"authorID":1,"content":"asadas","replies":null,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z"},{"id":9,"parentPostID":0,"parentID":5,"authorID":1,"content":"yeet comment","replies":null,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z"},{"id":10,"parentPostID":0,"parentID":9,"authorID":1,"content":"yeet comment","replies":null,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z"},{"id":11,"parentPostID":0,"parentID":5,"authorID":1,"content":"yeet comment","replies":null,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z"},{"id":12,"parentPostID":0,"parentID":10,"authorID":1,"content":"yeet comment","replies":null,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z"},{"id":13,"parentPostID":0,"parentID":10,"authorID":1,"content":"yeet comment","replies":null,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z"},{"id":14,"parentPostID":0,"parentID":10,"authorID":1,"content":"yeet comment","replies":null,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z"},{"id":15,"parentPostID":0,"parentID":12,"authorID":1,"content":"yeet comment","replies":null,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z"},{"id":16,"parentPostID":0,"parentID":11,"authorID":1,"content":"yeet comment","replies":null,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z"},{"id":17,"parentPostID":0,"parentID":16,"authorID":1,"content":"yeet comment","replies":null,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z"},{"id":18,"parentPostID":0,"parentID":17,"authorID":1,"content":"yeet comment","replies":null,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z"},{"id":19,"parentPostID":0,"parentID":18,"authorID":1,"content":"yeet comment","replies":null,"createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z"}]`

	type Comment struct {
		ID uint64 `db:"id" json:"id"`

		ParentPostID uint64 `db:"parent_post_id" json:"parentPostID"`
		ParentID     *int   `db:"parent_id" json:"parentID"` // allow nullable on lvl 0 comment
		AuthorID     uint64 `db:"author_id" json:"authorID"`
		Content      string `db:"content" json:"content"`

		Replies []*Comment `json:"replies"`

		CreatedAt time.Time `db:"created_at" json:"createdAt"`
		UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
	}

	var c []*Comment
	err := json.Unmarshal([]byte(listRaw), &c)
	if err != nil {
		panic(err)
	}

}
