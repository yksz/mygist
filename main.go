package main

import (
	"flag"
	"log"

	"./internal"
)

func main() {
	user := flag.String("user", "", "username")
	flag.Parse()

	err := internal.ListGists(*user)
	if err != nil {
		log.Fatal(err)
	}
}
