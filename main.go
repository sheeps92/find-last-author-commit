package main

import (
	"fmt"
	"log"
	"os"
	"syscall"

	"github.com/sheeps92/find-last-author-commit/src"
	"golang.org/x/term"
)

func run(org string) string {
	fmt.Print("Enter personal access token:")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("\n\n")
	}
	creds := src.Creds(bytePassword)
	repoList := src.ListRepos(org, creds)
	for _, repo := range repoList {
		fmt.Printf("### Finding author commit timestamp in repo: %s ###\n", repo)
		users := src.FindUsers(org, repo, creds)
		d := src.FindCommit(org, repo, users, creds)

		for _, commiter := range d {
			fmt.Printf("User: %s last commit was %s", commiter.User, commiter.Date)
		}
		fmt.Printf("### Ending search in repo: %s ###\n\n", repo)
	}
	return ""
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Invalid number of args. Please enter owner as first arg. ex. .\\find-last-author-commit.exe <owner>")
		os.Exit(1)
	}
	if err := run(os.Args[1]); err != "" {
		log.Fatal(err)
	}
}
