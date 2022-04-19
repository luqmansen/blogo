package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/luqmansen/blogo/internal/blogo"
	"github.com/luqmansen/blogo/internal/config"
	"github.com/luqmansen/blogo/internal/repository/postgres"
	"github.com/luqmansen/blogo/internal/server"
)

var (
	version      string
	gitCommitSha string
)

func main() {

	conf := config.LoadConfig()

	//db := sqldblogger.OpenDriver(conf.DatabaseURI, &pq.Driver{}, loggerAdapter /*, ...options */)
	//sqlx.BindDriver("postgres", sqlx.NAMED)
	db := sqlx.MustConnect("postgres", conf.DatabaseURI)

	postRepo := postgres.NewPostRepository(db)
	commentRepo := postgres.NewCommentRepository(db)
	postService := blogo.NewPostService(postRepo)
	commentService := blogo.NewCommentService(commentRepo)
	handler := server.NewHandler(conf, postService, commentService)

	//oauth2Config := &oauth2.Config{
	//	ClientID:     conf.GithwubClientId,
	//	ClientSecret: conf.GithubClientSecret,
	//	Scopes:       []string{"all"},
	//	RedirectURL:  fmt.Sprintf("http://%s:%s/oauth/callback", conf.Host, conf.Port),
	//	Endpoint:     oauthgithub.Endpoint,
	//}
	//oauthHandler := server.NewOauthHandler(&mock.OAuth2Mock{}, &mock.GitHubMock{})

	srv := server.New(conf, handler)
	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
