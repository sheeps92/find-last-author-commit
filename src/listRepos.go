package src

import (
	"context"
	"log"

	"github.com/google/go-github/github"
)

func ListRepos(owner string, creds *github.Client) []string {
	var repoList []string
	repos, _, err := creds.Repositories.List(context.TODO(), owner, &github.RepositoryListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(repos)
	for _, repo := range repos {
		repoList = append(repoList, *repo.Name)
	}

	return repoList
}
