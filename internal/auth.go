package internal

import (
	"context"

	"github.com/google/go-github/github"
)

func CreateAccessToken(username, password, note string) (string, error) {
	client := newClient(username, password)
	return createAccessToken(client, note)
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
