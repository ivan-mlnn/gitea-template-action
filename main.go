package main

import (
	"code.gitea.io/sdk/gitea"
	"fmt"
	gha "github.com/sethvargo/go-githubactions"
	"os"
	"strings"
)

func main() {
	ctx, err := gha.Context()
	if err != nil {
		gha.Fatalf("failed to get context: %v", err)
	}
	apiKey := os.Getenv("GITHUB_TOKEN")

	c, err := gitea.NewClient(ctx.ServerURL, gitea.SetToken(apiKey))
	if err != nil {
		gha.Fatalf("failed to create gitea client: %v", err)
	}
	owner := ctx.RepositoryOwner
	repo := strings.Split(ctx.Repository, "/")[1]

	r, _, err := c.GetRepo(owner, repo)
	if err != nil {
		gha.Fatalf("failed to GetRepo: %v", err)
	}
	_ = r.Description
	fmt.Println(r.Description)

	files, err := os.ReadDir(ctx.Workspace)
	if err != nil {
		gha.Fatalf("failed to ReadDir: %v", err)
	}

	for _, s := range files {
		if s.IsDir() == false {
			fmt.Println(s.Name())
		}
	}

	env := os.Environ()
	for _, s := range env {
		fmt.Println(s)
	}
}
