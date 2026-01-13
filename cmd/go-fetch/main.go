package main

import (
	"fmt"

	"github.com/parradam/go-fetch/internal/api/github"
)

func main() {
	client := github.NewClient()

	issues, err := client.FetchIssues("parradam", "wellwellwell")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Fetched %d issues:\n\n", len(issues))

	for i, issue := range issues {
		if i >= 5 {
			break
		}
		fmt.Printf("%d: %s (by %s)\n", issue.ID, issue.Title, issue.Author)
	}
}
