package main

import (
	"context"
	"fmt"
	"log"

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

	client := internal.NewClientWithOAuth2(auth.AccessToken)
	ctx := context.Background()
	gists, _, err := client.Gists.List(ctx, auth.Username, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, gist := range gists {
		fmt.Printf("%s:\n", *gist.ID)
		for _, file := range gist.Files {
			fmt.Printf("  %s\n", *file.Filename)
		}
	}
}
