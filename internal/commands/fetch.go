package commands

import (
	"fmt"

	"github.com/parradam/go-fetch/internal/api/github"
	"github.com/parradam/go-fetch/internal/output"
)

func Fetch(owner, repo string) error {
	client := github.NewClient()

	issues, err := client.FetchIssues(owner, repo)
	if err != nil {
		return fmt.Errorf("failed to fetch issues: %w", err)
	}

	formatter := &output.MarkdownFormatter{}
	filename := fmt.Sprintf("%s-%s-issues.md", owner, repo)

	err = formatter.Write(issues, owner, repo, filename)
	if err != nil {
		return fmt.Errorf("failed to write markdown: %w", err)
	}

	fmt.Printf("Saved %d issues to %s\n", len(issues), filename)

	return nil
}
