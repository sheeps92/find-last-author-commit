package main

import (
	"log"

	"github.com/google/go-github/github"
	"github.com/sheeps92/find-last-author-commit/src"
)

func run() *github.Contributor {
	creds := src.Creds()

	users := src.FindUsers("sheeps92", "find-last-author-commit", creds)
	return users
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
