package cmd

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
	"github.com/yksz/mygist/config"
	"github.com/yksz/mygist/internal"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List gists",
	Long:  "List gists",
	RunE:  doList,
}

func init() {
	RootCmd.AddCommand(listCmd)
}

func doList(cmd *cobra.Command, args []string) error {
	page := 1
	opt := &github.GistListOptions{
		ListOptions: github.ListOptions{Page: page},
	}

	ctx := context.Background()
	client := internal.NewClient(config.Conf.AccessToken)
	gists, _, err := client.Gists.List(ctx, config.Conf.Username, opt)
	if err != nil {
		return err
	}
	for _, gist := range gists {
		file := getFirstFile(gist)
		filename := *file.Filename
		description := *gist.Description
		fmt.Printf("%-20.20s\t%s\n", filename, description)
	}
	return nil
}

func getFirstFile(gist *github.Gist) *github.GistFile {
	for _, file := range gist.Files {
		return &file
	}
	return &github.GistFile{}
}
