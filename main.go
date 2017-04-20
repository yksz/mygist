package main

import (
	"fmt"
	"log"
	"os"

	"github.com/yksz/mygist/internal"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("usage: %s <username>\n", os.Args[0])
		os.Exit(1)
	}
	username := os.Args[1]

	password, err := internal.ReadPassword()
	if err != nil {
		log.Fatal(err)
		return
	}

	if err := internal.ListGists(username, password); err != nil {
		log.Fatal(err)
		return
	}
}
