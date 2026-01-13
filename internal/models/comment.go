package models

import "time"

type Comment struct {
	ID        int
	Body      string
	Author    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
