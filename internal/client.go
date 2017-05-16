package internal

import (
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func NewClient(token string) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	return github.NewClient(tc)
}

func newClient(username, password string) *github.Client {
	tp := github.BasicAuthTransport{
		Username: username,
		Password: password,
	}
	return github.NewClient(tp.Client())
}
