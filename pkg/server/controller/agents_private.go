package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/visualizer/pkg/entity/model"
	"github.com/ryo-arima/visualizer/pkg/entity/request"
	"github.com/ryo-arima/visualizer/pkg/entity/response"
	"github.com/ryo-arima/visualizer/pkg/server/repository"
)

type AgentsControllerForPrivate interface {
	GetAgents(c *gin.Context)
	CreateAgents(c *gin.Context)
	UpdateAgents(c *gin.Context)
	DeleteAgents(c *gin.Context)
}

type agentsControllerForPrivate struct {
	AgentsRepository repository.AgentsRepository
}

func (agentsController agentsControllerForPrivate) GetAgents(c *gin.Context) {
	var agentsRequest request.AgentsRequest
	if err := c.Bind(&agentsRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.AgentsResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Agentss: []response.Agents{}})
		return
	}
	res := agentsController.AgentsRepository.GetAgents()
	c.JSON(http.StatusOK, res)
	return
}

func (agentsController agentsControllerForPrivate) CreateAgents(c *gin.Context) {
	var agentsRequest request.AgentsRequest
	if err := c.Bind(&agentsRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.AgentsResponse{Code: "SERVER_CONTROLLER_CREATE__FOR__001", Message: err.Error(), Agentss: []response.Agents{}})
		return
	}
	var agentsModel model.Agents
	res := agentsController.AgentsRepository.CreateAgents(agentsModel)
	c.JSON(http.StatusOK, res)
	return
}

func (agentsController agentsControllerForPrivate) UpdateAgents(c *gin.Context) {
	var agentsRequest request.AgentsRequest
	if err := c.Bind(&agentsRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.AgentsResponse{Code: "SERVER_CONTROLLER_UPDATE__FOR__001", Message: err.Error(), Agentss: []response.Agents{}})
		return
	}
	var agentsModel model.Agents
	res := agentsController.AgentsRepository.UpdateAgents(agentsModel)
	c.JSON(http.StatusOK, res)
	return
}

func (agentsController agentsControllerForPrivate) DeleteAgents(c *gin.Context) {
	var agentsRequest request.AgentsRequest
	if err := c.Bind(&agentsRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.AgentsResponse{Code: "SERVER_CONTROLLER_DELETE__FOR__001", Message: err.Error(), Agentss: []response.Agents{}})
		return
	}
	var uuid string
	res := agentsController.AgentsRepository.DeleteAgents(uuid)
	c.JSON(http.StatusOK, res)
	return
}

func NewAgentsControllerForPrivate(agentsRepository repository.AgentsRepository) AgentsControllerForPrivate {
	return &agentsControllerForPrivate{AgentsRepository: agentsRepository}
}
