package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/yksz/mygist/internal"
)

const (
	AppName       = "mygist"
	workspaceName = "." + AppName
)

type Config struct {
	Username    string `json:"username"`
	AccessToken string `json:"access_token"`
}

var Conf Config

func GetWorkspace() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return home, err
	}
	workspace := filepath.Join(home, workspaceName)
	if err := os.MkdirAll(workspace, 0700); err != nil {
		return "", err
	}
	return workspace, nil
}

func (c *Config) Create() error {
	username, err := internal.ReadUsername()
	if err != nil {
		return err
	}
	password, err := internal.ReadPassword()
	if err != nil {
		return err
	}
	note := AppName + "_" + time.Now().Format("20060102")
	token, err := internal.CreateAccessToken(username, password, note)
	if err != nil {
		return err
	}
	c.Username = username
	c.AccessToken = token
	return nil
}

func (c *Config) Load(filename string) error {
	if !internal.Exists(filename) {
		return fmt.Errorf("config file not found: %s", filename)
	}
	return c.load(filename)
}

func (c *Config) load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	if err := json.NewDecoder(file).Decode(c); err != nil {
		return err
	}
	return nil
}

func (c *Config) Save(filename string) error {
	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, 0700); err != nil {
		return err
	}
	return c.save(filename)
}

func (c *Config) save(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(c)
}
