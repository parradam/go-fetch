package commands

import (
	"fmt"

	"github.com/parradam/go-fetch/internal/api/github"
)

func Fetch(owner, repo string) error {
	client := github.NewClient()

	issues, err := client.FetchIssues(owner, repo)
	if err != nil {
		return fmt.Errorf("failed to fetch issues: %w", err)
	}

	fmt.Printf("Fetched %d issues:\n\n", len(issues))

	for _, issue := range issues {
		fmt.Printf("%d: %s (by %s)\n", issue.ID, issue.Title, issue.Author)
	}

	return nil
}
