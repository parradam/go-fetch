package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/parradam/go-fetch/internal/models"
)

type Client struct {
	BaseURL string
}

func NewClient() *Client {
	return &Client{
		BaseURL: "https://api.github.com",
	}
}

func (c *Client) FetchIssues(owner, repo string) ([]*models.Issue, error) {
	url := fmt.Sprintf("%s/repos/%s/%s/issues", c.BaseURL, owner, repo)
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer func() {
		closeErr := resp.Body.Close()
		if closeErr != nil {
			fmt.Fprintf(os.Stderr, "Warning: failed to close response body: %v\n", closeErr)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GitHub API returned status %d for %s/%s", resp.StatusCode, owner, repo)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var gitHubIssues []GitHubIssue
	err = json.Unmarshal(body, &gitHubIssues)
	if err != nil {
		return nil, err
	}

	var issues []*models.Issue
	for _, g := range gitHubIssues {
		issue := g.ToIssue()
		issues = append(issues, issue)
	}

	return issues, nil
}
