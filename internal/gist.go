package internal

import (
	"encoding/json"
	"errors"
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
	Files       map[string]*File
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"created_at"`
}

type Owner struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type File struct {
	Filename string
	Language string
}

func ListGists(token, username string) error {
	if token == "" {
		return errors.New("token must not be empty")
	}
	if username == "" {
		return errors.New("username must not be empty")
	}

	url := apiURL + "/users/" + username + "/gists"
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "token "+token)
	client := &http.Client{}
	resp, err := client.Do(req)
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
		fmt.Printf("%s: %s\n", gist.ID, gist.Description)
		for name, file := range gist.Files {
			fmt.Printf("  %s %s\n", name, file.Language)
		}
	}
	return nil
}
