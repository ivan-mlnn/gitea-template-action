package main

import (
	"code.gitea.io/sdk/gitea"
	"fmt"
	gha "github.com/sethvargo/go-githubactions"
	"io/fs"
	"os"
	"path/filepath"
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

	filepath.WalkDir(ctx.Workspace, func(path string, d fs.DirEntry, err error) error {
		if d.Name() == ".git" {
			return fs.SkipDir
		}

		if d.IsDir() == false {
			fmt.Println(path, d.Name())
		}

		return nil
	})

	env := os.Environ()
	for _, s := range env {
		fmt.Println(s)
	}
}
