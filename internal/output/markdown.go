package output

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/parradam/go-fetch/internal/models"
)

type MarkdownFormatter struct{}

func (m *MarkdownFormatter) Write(issues []*models.Issue, owner, repo, filename string) error {
	var builder strings.Builder

	fmt.Fprintf(&builder, "# Issues for %s/%s\n\n", owner, repo)
	fmt.Fprintf(&builder, "Fetched: %s", time.Now().Format(time.RFC3339))

	for _, issue := range issues {
		fmt.Fprintf(&builder, "\n\n## Issue #%d: %s\n\n", issue.ID, issue.Title)
		fmt.Fprintf(&builder, "**Author:** %s  \n**State:** %s  \n**Created:** %s\n\n", issue.Author, issue.State, issue.CreatedAt.Format(time.DateOnly))
		builder.WriteString(issue.Body)
		builder.WriteString("\n\n---")
	}

	builder.WriteString("\n")

	err := os.WriteFile(filename, []byte(builder.String()), 0644)
	if err != nil {
		return fmt.Errorf("failed to write file %s: %w", filename, err)
	}
	return nil
}
