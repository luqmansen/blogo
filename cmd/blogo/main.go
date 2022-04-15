package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/luqmansen/blogo/internal/blogo"
	"github.com/luqmansen/blogo/internal/config"
	"github.com/luqmansen/blogo/internal/mock"
	"github.com/luqmansen/blogo/internal/repository/postgres"
	"github.com/luqmansen/blogo/internal/server"
)

var (
	version      string
	gitCommitSha string
)

func main() {

	conf := config.LoadConfig(".")

	db := sqlx.MustConnect("postgres", conf.DbUri)

	postRepo := postgres.NewPostRepository(db)
	postSvc := blogo.NewPostService(postRepo)
	handler := server.NewHandler(postSvc)

	//oauth2Config := &oauth2.Config{
	//	ClientID:     conf.GithwubClientId,
	//	ClientSecret: conf.GithubClientSecret,
	//	Scopes:       []string{"all"},
	//	RedirectURL:  fmt.Sprintf("http://%s:%s/oauth/callback", conf.Host, conf.Port),
	//	Endpoint:     oauthgithub.Endpoint,
	//}
	oauthHandler := server.NewOauthHandler(&mock.OAuth2Mock{}, &mock.GitHubMock{})

	srv := server.New(conf, handler, oauthHandler)
	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
