package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"

	"github.com/yksz/mygist/github"
	"github.com/yksz/mygist/internal"
)

const (
	appHomeDirName = ".mygist"
	configFileName = "config"
)

type Config struct {
	Username    string `json:"username"`
	AccessToken string `json:"access_token"`
}

var Conf Config

func init() {
	if err := initConf(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func initConf() error {
	appHomeDir, err := GetAppHomeDir()
	if err != nil {
		return err
	}
	configFile := filepath.Join(appHomeDir, configFileName)
	if internal.Exists(configFile) {
		return Conf.Load(configFile)
	} else {
		if err := Conf.Create(); err != nil {
			return err
		}
		return Conf.Save(configFile)
	}
}

func GetAppHomeDir() (string, error) {
	userHomeDir, err := getUserHomeDir()
	if err != nil {
		return "", err
	}
	appHomeDir := filepath.Join(userHomeDir, appHomeDirName)
	if err := os.MkdirAll(appHomeDir, 0700); err != nil {
		return "", err
	}
	return appHomeDir, nil
}

func getUserHomeDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return usr.HomeDir, nil
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
	token, err := github.CreateAccessToken(username, password)
	if err != nil {
		return err
	}
	c.Username = username
	c.AccessToken = token
	return nil
}

func (c *Config) Load(filename string) error {
	if !internal.Exists(filename) {
		return fmt.Errorf("file not found: %s", filename)
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
