package main

import (
	"fmt"
	"log"

	"github.com/yksz/mygist/gist"
	"github.com/yksz/mygist/internal"
)

const (
	apiURL = "https://api.github.com"
)

func main() {
	auth, err := internal.GetAuthInfo()
	if err != nil {
		log.Fatal(err)
		return
	}

	gister := gist.NewGister(apiURL, auth.AccessToken)
	gists, err := gister.ListGists(auth.Username)
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, gist := range gists {
		fmt.Printf("%s: %s\n", gist.ID, gist.Description)
		for name, file := range gist.Files {
			fmt.Printf("  %s %s\n", name, file.Language)
		}
	}
}
