package main

import (
	"context"
	"fmt"
	"log"

	"github.com/yksz/mygist/config"
	"github.com/yksz/mygist/github"
)

func main() {
	client := github.NewClient(config.Conf.AccessToken)
	ctx := context.Background()
	gists, _, err := client.Gists.List(ctx, config.Conf.Username, nil)
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
