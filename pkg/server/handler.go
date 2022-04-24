package server

import (
	"encoding/json"
	"fmt"
	"github.com/luqmansen/blogo/pkg/blogo"
	"github.com/luqmansen/blogo/pkg/config"
)

type handler struct {
	config         *config.Configuration
	oauthConfig    OAuth2ConfigInterface
	postService    *blogo.PostService
	commentService *blogo.CommentService
	reactService   *blogo.ReactService
}

func NewHandler(
	config *config.Configuration,
	postService *blogo.PostService,
	commentService *blogo.CommentService,
	reactService *blogo.ReactService,
	//oauthConfig OAuth2ConfigInterface,
) ServerInterface {
	return handler{
		config:         config,
		postService:    postService,
		commentService: commentService,
		reactService:   reactService,
		//oauthConfig: oauthConfig,
	}
}

func marshal(src, dest interface{}) {
	b, err := json.Marshal(src)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(b, dest)
	if err != nil {
		panic(err)
	}
}

func debugStruct(d interface{}) {
	s, _ := json.MarshalIndent(d, "", "\t")
	fmt.Println(string(s))
}
