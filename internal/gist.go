package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	apiURL = "https://api.github.com"
)

type Gist struct {
	ID          string
	Description string
	Public      bool
	URL         string
	Owner       *Owner
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"created_at"`
}

type Owner struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func ListGists(username string) error {
	if username == "" {
		panic("username must not be empty")
	}

	resp, err := http.Get(apiURL + "/users/" + username + "/gists")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status: %s", resp.Status)
	}
	var gists []Gist
	if err := json.NewDecoder(resp.Body).Decode(&gists); err != nil {
		return err
	}
	for _, gist := range gists {
		fmt.Printf("#%s %s\n", gist.ID, gist.Description)
	}
	return nil
}
