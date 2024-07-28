package response

import "time"

type AgentsResponse struct {
	Code    string
	Message string
	Agentss []Agents
}

type Agents struct {
	ID        uint
	UUID      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
