package blogo

import (
	"context"
	"github.com/google/go-github/v42/github"
	"net/http"
)

// RepositoriesService handles communication with the repository related methods
// of the GitHub API.
// https://godoc.org/github.com/google/go-github/github#RepositoriesService
type RepositoriesService interface {
	Get(context.Context, string, string) (*github.Repository, *github.Response, error)
	// ...
}

// UsersService handles communication with the user related methods
// of the GitHub API.
// https://godoc.org/github.com/google/go-github/github#UsersService
type UsersService interface {
	Get(context.Context, string) (*github.User, *github.Response, error)
	// ...
}

// GitHubClient manages communication with the GitHub API.
// https://github.com/google/go-github/issues/113
type GitHubClient struct {
	Repositories RepositoriesService
	Users        UsersService
}

// GitHubInterface defines all necessary methods.
// https://godoc.org/github.com/google/go-github/github#NewClient
type GitHubInterface interface {
	NewClient(httpClient *http.Client) GitHubClient
}

type GitHubCreator struct{}

// NewClient returns a new GitHubInterface instance.
func (g *GitHubCreator) NewClient(httpClient *http.Client) GitHubClient {
	client := github.NewClient(httpClient)
	return GitHubClient{
		Repositories: client.Repositories,
		Users:        client.Users,
	}
}
