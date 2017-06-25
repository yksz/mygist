package cmd

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"sort"

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
	pagination := newPagination()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if err := showListAndCommand(pagination); err != nil {
			return err
		}
		if !scanner.Scan() {
			return scanner.Err()
		}
		line := scanner.Text()
		if err := eval(line, pagination); err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
	}
	return nil
}

const perPage = 10

type pagination struct {
	Index int
}

func newPagination() *pagination {
	return &pagination{Index: 1}
}

func (p *pagination) Next() int {
	p.Index++
	return p.Index
}

func (p *pagination) Prev() int {
	if p.Index > 1 {
		p.Index--
	}
	return p.Index
}

func showListAndCommand(pagination *pagination) error {
	if err := showList(pagination); err != nil {
		return err
	}
	fmt.Println()
	fmt.Println("*** Commands ***")
	fmt.Println(" n: next\t p: previous\t q: quit")
	fmt.Print("What now> ")
	return nil
}

func showList(pagination *pagination) error {
	gists, err := fetchGists(pagination.Index)
	if err != nil {
		return err
	}

	sort.Slice(gists, func(i, j int) bool {
		return (*gists[i].UpdatedAt).After(*gists[j].UpdatedAt)
	})
	fmt.Printf("%3s: %-20s\t%s\n", "id", "filename", "description")
	for i, gist := range gists {
		file := getFirstFile(gist)
		filename := *file.Filename
		description := *gist.Description
		fmt.Printf("%3d: %-20.20s\t%s\n", i+1, filename, description)
	}
	return nil
}

func fetchGists(page int) ([]*github.Gist, error) {
	opt := &github.GistListOptions{
		ListOptions: github.ListOptions{Page: page, PerPage: perPage},
	}

	ctx := context.Background()
	client := internal.NewClient(config.Conf.AccessToken)
	gists, _, err := client.Gists.List(ctx, config.Conf.Username, opt)
	return gists, err
}

func getFirstFile(gist *github.Gist) *github.GistFile {
	for _, file := range gist.Files {
		return &file
	}
	return &github.GistFile{}
}

func eval(line string, pagination *pagination) error {
	if line == "q" {
		return io.EOF
	} else if line == "n" {
		pagination.Next()
	} else if line == "p" {
		pagination.Prev()
	}
	return nil
}
