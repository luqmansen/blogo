package mock

import (
	"context"
	"github.com/google/go-github/v42/github"
	"github.com/luqmansen/blogo/internal/blogo"
	"net/http"
)

// RepositoriesMock mocks RepositoriesService
type RepositoriesMock struct {
	blogo.RepositoriesService
}

// Get returns a repository.
func (r *RepositoriesMock) Get(context.Context, string, string) (*github.Repository, *github.Response, error) {
	return &github.Repository{
		ID:              github.Int64(185409993),
		Name:            github.String("wayne"),
		Description:     github.String("some description"),
		Language:        github.String("JavaScript"),
		StargazersCount: github.Int(3141),
		HTMLURL:         github.String("https://www.foo.com"),
		FullName:        github.String("john/wayne"),
	}, nil, nil
}

// UsersMock mocks UsersService
type UsersMock struct {
	blogo.UsersService
}

// Get returns a user.
func (u *UsersMock) Get(context.Context, string) (*github.User, *github.Response, error) {
	return &github.User{
		Login: github.String("john"),
	}, nil, nil
}

// GitHubMock implements GitHubInterface.
type GitHubMock struct{}

// NewClient something
func (g *GitHubMock) NewClient(httpClient *http.Client) blogo.GitHubClient {
	return blogo.GitHubClient{
		Repositories: &RepositoriesMock{},
		Users:        &UsersMock{},
	}
}
