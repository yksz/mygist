package internal

import (
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func NewClientWithBasicAuth(username, password string) *github.Client {
	tp := github.BasicAuthTransport{
		Username: strings.TrimSpace(username),
		Password: strings.TrimSpace(password),
	}
	return github.NewClient(tp.Client())
}

func NewClientWithOAuth2(token string) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	return github.NewClient(tc)
}
