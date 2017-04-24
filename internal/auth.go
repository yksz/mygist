package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

const (
	cacheFileName = "auth"
)

type AuthInfo struct {
	Username    string
	AccessToken string `json:"access_token"`
}

func GetAuthInfo() (*AuthInfo, error) {
	path, err := getCacheFilePath()
	if err != nil {
		return nil, err
	}
	if exists(path) {
		return readAuthInfo(path)
	}
	return nil, fmt.Errorf("file not found: %s", path)
}

func getCacheFilePath() (string, error) {
	workspace, err := GetWorkspace()
	if err != nil {
		return "", err
	}
	path := workspace + string(os.PathSeparator) + cacheFileName
	return path, nil
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func readAuthInfo(filename string) (*AuthInfo, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var authInfo AuthInfo
	if err := json.NewDecoder(file).Decode(&authInfo); err != nil {
		return nil, err
	}
	return &authInfo, nil
}

func ReadPassword() (string, error) {
	fmt.Printf("Password: ")
	password, err := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	if err != nil {
		return "", err
	}
	return string(password), nil
}
