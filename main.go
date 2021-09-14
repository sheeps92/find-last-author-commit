package main

import (
	"log"

	"github.com/sheeps92/find-last-author-commit/src"
)

func run() string {
	creds := src.Creds()

	users := src.FindUsers("sheeps92", "find-last-author-commit", creds)
	date := src.FindCommit("sheeps92", "find-last-author-commit", users, creds)
	log.Println(date)
	return users
}

func main() {
	if err := run(); err != "" {
		log.Fatal(err)
	}
}
