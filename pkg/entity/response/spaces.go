package response

import "time"

type SpacesResponse struct {
	Code    string
	Message string
	Spacess []Spaces
}

type Spaces struct {
	ID        uint
	UUID      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
