package src

import (
	"context"
	"log"
	"time"

	"github.com/google/go-github/github"
)

type commitDateList struct {
	User string
	Date *time.Time
}

type commitDate []commitDateList

func FindCommit(owner string, repo string, author []string, creds *github.Client) commitDate {
	commitDate := commitDate{}
	for _, user := range author {
		commit, _, err := creds.Repositories.ListCommits(context.TODO(), owner, repo, &github.CommitsListOptions{Author: user, ListOptions: github.ListOptions{PerPage: 1}})
		if err != nil {
			log.Fatal(err)
		}
		for _, c := range commit {
			yeet := c.Commit
			cd := yeet.Committer.Date
			commitDate = append(commitDate, commitDateList{user, cd})
		}
	}
	return commitDate
}
