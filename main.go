package main

import (
	"log"

	"github.com/yksz/mygist/internal"
)

func main() {
	auth, err := internal.GetAuthInfo()
	if err != nil {
		log.Fatal(err)
		return
	}

	if err := internal.ListGists(auth.AccessToken, auth.Username); err != nil {
		log.Fatal(err)
		return
	}
}
