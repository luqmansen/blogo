package server

import (
	"context"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/luqmansen/blogo/pkg/blogo"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"net/http"
)

type OAuth2ConfigInterface interface {
	AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string
	Exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error)
	Client(ctx context.Context, t *oauth2.Token) *http.Client
}

type OauthHandler struct {
	oauth2Config  OAuth2ConfigInterface
	githubService blogo.GitHubInterface
}

func NewOauthHandler(
	configInterface OAuth2ConfigInterface,
	githubService blogo.GitHubInterface,
) OauthHandler {
	return OauthHandler{
		oauth2Config:  configInterface,
		githubService: githubService,
	}
}

func (h OauthHandler) Authorize(ctx *gin.Context) {
	sess := sessions.Default(ctx)

	state := uuid.New().String()
	sess.Set("state", state)
	if err := sess.Save(); err != nil {
		log.Error(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	ctx.Redirect(http.StatusTemporaryRedirect, h.oauth2Config.AuthCodeURL(state))
}

func (h OauthHandler) GithubCallbackHandler(ctx *gin.Context) {
	sess := sessions.Default(ctx)

	storedState := sess.Get("state")
	returnedState := ctx.Query("state")
	if returnedState != storedState.(string) {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, "Invalid state parameter")
		return
	}

	code := ctx.Query("code")
	if code == "" {
		ctx.String(http.StatusBadRequest, "no code in query param")
		return
	}
	token, err := h.oauth2Config.Exchange(ctx, code)
	if err != nil {
		log.Error(err)
	}

	client := h.oauth2Config.Client(ctx, token)

	githubClient := h.githubService.NewClient(client)

	repo, _, _ := githubClient.Repositories.Get(ctx, "1", "")

	ctx.JSON(200, repo)
}
