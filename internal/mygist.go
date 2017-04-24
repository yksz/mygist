package internal

import (
	"os"
	"os/user"
)

const (
	dirname = ".mygist"
)

func GetWorkspace() (string, error) {
	homeDir, err := getHomeDir()
	if err != nil {
		return "", err
	}
	workspace := homeDir + string(os.PathSeparator) + dirname
	return workspace, nil
}

func getHomeDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return usr.HomeDir, nil
}
