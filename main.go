package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sheeps92/find-last-author-commit/src"
)

func run(org string) string {
	creds := src.Creds()
	repoList := src.ListRepos(org, creds)
	for _, repo := range repoList {
		log.Printf("### Finding author commit timestamp in repo: %s ###\n", repo)
		users := src.FindUsers(org, repo, creds)
		d := src.FindCommit(org, repo, users, creds)

		for _, commiter := range d {
			fmt.Printf("User: %s last commit was %s", commiter.User, commiter.Date)
		}
		log.Printf("### Ending search in repo: %s ###\n", repo)
	}
	return ""
}

func main() {
	if err := run(os.Args[1]); err != "" {
		log.Fatal(err)
	}
}
