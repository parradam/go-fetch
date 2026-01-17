package output

import "github.com/parradam/go-fetch/internal/models"

type Formatter interface {
	Write(issues []*models.Issue, owner, repo, filename string) error
}
