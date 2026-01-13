package github

import (
	"time"

	"github.com/parradam/go-fetch/internal/models"
)

type GitHubUser struct {
	Login string `json:"login"`
}

type GitHubIssue struct {
	Number    int        `json:"number"`
	Title     string     `json:"title"`
	Body      string     `json:"body"`
	User      GitHubUser `json:"user"`
	State     string     `json:"state"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func (g *GitHubIssue) ToIssue() *models.Issue {
	return &models.Issue{
		ID:        g.Number,
		Title:     g.Title,
		Body:      g.Body,
		Author:    g.User.Login,
		State:     g.State,
		CreatedAt: g.CreatedAt,
		UpdatedAt: g.UpdatedAt,
	}
}
