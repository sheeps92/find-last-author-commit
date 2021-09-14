package src

import (
	"github.com/google/go-github/github"
)

func Creds() *github.Client {
	client := github.NewClient(nil)

	return client
}
