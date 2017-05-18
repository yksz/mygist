package cmd

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"

	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
	"github.com/yksz/mygist/config"
	"github.com/yksz/mygist/internal"
)

const syncDirName = "gists"

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync gists",
	Long:  "Sync gists",
	RunE:  execSync,
}

func init() {
	RootCmd.AddCommand(syncCmd)
}

func execSync(cmd *cobra.Command, args []string) error {
	syncDir, err := getSyncDir()
	if err != nil {
		return err
	}

	ctx := context.Background()
	client := internal.NewClient(config.Conf.AccessToken)
	gists, _, err := client.Gists.List(ctx, config.Conf.Username, nil)
	if err != nil {
		return err
	}

	count := cloneGists(gists, syncDir)
	fmt.Printf("sync %d gists\n", count)
	return nil
}

func getSyncDir() (string, error) {
	workspace, err := config.GetWorkspace()
	if err != nil {
		return "", err
	}
	syncDir := filepath.Join(workspace, syncDirName)
	if err := os.MkdirAll(syncDir, 0700); err != nil {
		return "", err
	}
	return syncDir, nil
}

func cloneGists(gists []*github.Gist, dir string) int {
	count := 0
	var wg sync.WaitGroup
	for _, g := range gists {
		wg.Add(1)
		go func(gist *github.Gist) {
			defer wg.Done()
			if err := cloneGist(gist, dir); err == nil {
				count++
			}
		}(g)
	}
	wg.Wait()
	return count
}

func cloneGist(gist *github.Gist, dir string) error {
	if gist.GitPullURL == nil {
		return nil
	}

	cmd := exec.Command("git", "clone", *gist.GitPullURL)
	cmd.Dir = dir
	return cmd.Run()
}
