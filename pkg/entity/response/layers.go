package response

import "time"

type LayersResponse struct {
	Code    string
	Message string
	Layerss []Layers
}

type Layers struct {
	ID        uint
	UUID      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
