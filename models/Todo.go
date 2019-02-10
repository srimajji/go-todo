package models

import "time"

// Todo model
type Todo struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	DateCreated time.Time `json:"dateCreated"`
}
