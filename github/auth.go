package github

import (
	"context"
	"strings"

	"github.com/google/go-github/github"
)

const (
	accessTokenNote = "mygist"
)

func CreateAccessToken(username, password string) (string, error) {
	client := newClient(username, password)
	return createAccessToken(client, accessTokenNote)
}

func newClient(username, password string) *github.Client {
	tp := github.BasicAuthTransport{
		Username: strings.TrimSpace(username),
		Password: strings.TrimSpace(password),
	}
	return github.NewClient(tp.Client())
}

func createAccessToken(client *github.Client, note string) (string, error) {
	ctx := context.Background()
	req := &github.AuthorizationRequest{
		Scopes: []github.Scope{github.ScopeGist},
		Note:   &note,
	}
	auth, _, err := client.Authorizations.Create(ctx, req)
	if err != nil {
		return "", err
	}
	return *auth.Token, nil
}
