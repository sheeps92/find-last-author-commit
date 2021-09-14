package src

import (
	"context"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func Creds(pat []byte) *github.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: string(pat)},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	return client
}
