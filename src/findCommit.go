package src

import (
	"context"
	"log"
	"time"

	"github.com/google/go-github/github"
)

func FindCommit(owner string, repo string, author string, creds *github.Client) *time.Time {
	commit, _, err := creds.Repositories.ListCommits(context.TODO(), owner, repo, &github.CommitsListOptions{Author: author, ListOptions: github.ListOptions{PerPage: 1}})
	if err != nil {
		log.Fatal(err)
	}
	var commitDate *time.Time
	for _, c := range commit {
		yeet := c.Commit
		commitDate = yeet.Committer.Date
	}
	return commitDate
}
