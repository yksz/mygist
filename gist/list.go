package gist

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func (g *Gister) ListGists(username string) ([]Gist, error) {
	if username == "" {
		return nil, errors.New("username must not be empty")
	}

	url := g.url + "/users/" + username + "/gists"
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "token "+g.token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %s", resp.Status)
	}

	var gists []Gist
	if err := json.NewDecoder(resp.Body).Decode(&gists); err != nil {
		return nil, err
	}
	return gists, nil
}
