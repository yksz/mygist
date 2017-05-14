package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/yksz/mygist/config"
	"github.com/yksz/mygist/internal"
)

const (
	configFileName = "config"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   config.AppName,
	Short: config.AppName + " - a private gist client",
	Long:  config.AppName + " - a private gist client",
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	workspace, err := config.GetWorkspace()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	configFile := filepath.Join(workspace, configFileName)
	if err := loadOrCreateConfig(configFile); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func loadOrCreateConfig(file string) error {
	if internal.Exists(file) {
		return config.Conf.Load(file)
	} else {
		if err := config.Conf.Create(); err != nil {
			return err
		}
		return config.Conf.Save(file)
	}
}
