package models

import "time"

type Issue struct {
	ID        int
	Title     string
	Body      string
	Author    string
	State     string
	Comments  []Comment
	CreatedAt time.Time
	UpdatedAt time.Time
}
