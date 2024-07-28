package repository

import (
	"time"

	"github.com/ryo-arima/visualizer/pkg/config"
	"github.com/ryo-arima/visualizer/pkg/entity/model"
	"gorm.io/gorm"
)

type AgentsRepository interface {
	GetAgents() []model.Agents
	CreateAgents(agents model.Agents) *gorm.DB
	UpdateAgents(agents model.Agents) *gorm.DB
	DeleteAgents(uuid string) *gorm.DB
}

type agentsRepository struct {
	BaseConfig config.BaseConfig
}

func (agentsRepository agentsRepository) GetAgents() []model.Agents {
	var agentss []model.Agents
	agentsRepository.BaseConfig.DBConnection.Find(&agentss)
	return agentss
}

func (agentsRepository agentsRepository) CreateAgents(agents model.Agents) *gorm.DB {
	results := agentsRepository.BaseConfig.DBConnection.Create(&agents)
	return results
}

func (agentsRepository agentsRepository) UpdateAgents(agents model.Agents) *gorm.DB {
	results := agentsRepository.BaseConfig.DBConnection.Model(model.Agents{}).Where("uuid = ?", agents.UUID).Updates(agents)
	return results
}

func (agentsRepository agentsRepository) DeleteAgents(uuid string) *gorm.DB {
	results := agentsRepository.BaseConfig.DBConnection.Model(model.Agents{}).Where("uuid = ?", uuid).Update("deleted_at", time.Now())
	return results
}

func NewAgentsRepository(conf config.BaseConfig) AgentsRepository {
	return &agentsRepository{BaseConfig: conf}
}
