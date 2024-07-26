package response

import "time"

type ElementsResponse struct {
	Code      string
	Message   string
	Elementss []Elements
}

type Elements struct {
	ID        uint
	UUID      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
