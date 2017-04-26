package gist

import "time"

type Gister struct {
	url   string
	token string
}

func NewGister(url, token string) *Gister {
	return &Gister{url: url, token: token}
}

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
