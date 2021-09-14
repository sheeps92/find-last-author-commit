package src

import (
	"context"
	"log"

	"github.com/google/go-github/github"
)

func FindUsers(org string, repo string, creds *github.Client) []string {
	var userList []string
	stats, _, err := creds.Repositories.ListContributorsStats(context.TODO(), org, repo)
	if err != nil {
		log.Fatal(err)
	}
	var author *github.Contributor
	for _, user := range stats {
		author = user.Author
		userList = append(userList, *author.Login)
	}

	return userList
}
