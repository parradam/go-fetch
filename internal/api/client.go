package api

import "github.com/parradam/go-fetch/internal/models"

type CodePlatform interface {
	FetchIssues(owner, repo string) ([]models.Issue, error)
	FetchComments(owner, repo string, issueID int) ([]models.Comment, error)
}
