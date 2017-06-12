package cmd

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
	"github.com/yksz/mygist/config"
	"github.com/yksz/mygist/internal"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new gists",
	Long:  "Create a new gists",
	RunE:  doNew,
}

func init() {
	RootCmd.AddCommand(newCmd)
}

func doNew(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("filenames not found")
	}

	description := ""
	public := false
	files := createGistFiles(args)

	ctx := context.Background()
	client := internal.NewClient(config.Conf.AccessToken)
	gist, _, err := client.Gists.Create(ctx, &github.Gist{
		Description: &description,
		Public:      &public,
		Files:       files,
	})
	if err != nil {
		return err
	}

	if gist.HTMLURL != nil {
		fmt.Printf("%s\n", *gist.HTMLURL)
	}
	return nil
}

func createGistFiles(filenames []string) map[github.GistFilename]github.GistFile {
	files := make(map[github.GistFilename]github.GistFile)
	for _, filename := range filenames {
		content, err := readContent(filename)
		if err != nil {
			continue
		}
		files[github.GistFilename(filename)] = github.GistFile{
			Filename: &filename,
			Content:  &content,
		}
	}
	return files
}

func readContent(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
