package mock

import (
	"context"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"net/http"
	"net/url"
	"time"
)

type OAuth2Mock struct{}
type Oauth2HandlerMock struct{}

func (o OAuth2Mock) AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string {

	u := url.URL{
		Scheme: "http",
		Host:   "localhost:8081",
		Path:   "login/oauth/authorize",
	}

	v := url.Values{}
	v.Set("state", state)

	u.RawQuery = v.Encode()
	return u.String()
}

func (h Oauth2HandlerMock) DevOauthAuthorize(ctx *gin.Context) {
	state := ctx.Query("state")

	u, err := url.Parse("http://localhost:8081/oauth/callback")
	if err != nil {
		panic(err)
	}

	v := url.Values{}
	v.Set("code", "code")
	v.Set("state", state)
	u.RawQuery = v.Encode()
	ctx.Redirect(http.StatusTemporaryRedirect, u.String())
}

func (o OAuth2Mock) Exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	return &oauth2.Token{
		AccessToken: "AccessToken",
		Expiry:      time.Now().Add(1 * time.Hour),
	}, nil
}

// Client returns a new http.Client.
func (o *OAuth2Mock) Client(ctx context.Context, t *oauth2.Token) *http.Client {
	return &http.Client{}
}
