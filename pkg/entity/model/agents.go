package model

import "time"

type Agents struct {
	ID        uint `gorm:"primaryKey,autoIncrement"`
	UUID      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

type AgentRpcArguments struct {
}

type AgentRpcReturns struct {
}
