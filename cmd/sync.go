package cmd

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"

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
	RunE:  sync,
}

func init() {
	RootCmd.AddCommand(syncCmd)
}

func sync(cmd *cobra.Command, args []string) error {
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

	for _, gist := range gists {
		<-cloneGist(gist, syncDir)
	}
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

func cloneGist(gist *github.Gist, dir string) <-chan error {
	ch := make(chan error)

	if gist.GitPullURL == nil {
		ch <- nil
		return ch
	}

	go func() {
		cmd := exec.Command("git", "clone", *gist.GitPullURL)
		cmd.Dir = dir
		ch <- cmd.Run()
	}()
	return ch
}
