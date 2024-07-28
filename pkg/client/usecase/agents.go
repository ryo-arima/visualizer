package usecase

import (
	"fmt"

	"github.com/ryo-arima/visualizer/pkg/client/repository"
	"github.com/ryo-arima/visualizer/pkg/entity/request"
)

type AgentsUsecase interface {
	BootstrapAgentsForDB(request request.AgentsRequest)
	GetAgentsForPublic(request request.AgentsRequest)
	GetAgentsForInternal(request request.AgentsRequest)
	GetAgentsForPrivate(request request.AgentsRequest)
	CreateAgentsForPublic(request request.AgentsRequest)
	CreateAgentsForInternal(request request.AgentsRequest)
	CreateAgentsForPrivate(request request.AgentsRequest)
	UpdateAgentsForPublic(request request.AgentsRequest)
	UpdateAgentsForInternal(request request.AgentsRequest)
	UpdateAgentsForPrivate(request request.AgentsRequest)
	DeleteAgentsForPublic(request request.AgentsRequest)
	DeleteAgentsForInternal(request request.AgentsRequest)
	DeleteAgentsForPrivate(request request.AgentsRequest)
}

type agentsUsecase struct {
	AgentsRepository   repository.AgentsRepository
}

//Bootstrap
func (agentsUsecase agentsUsecase) BootstrapAgentsForDB(request request.AgentsRequest) {
	agentss := agentsUsecase.AgentsRepository.BootstrapAgentsForDB(request)
	fmt.Println(agentss)
}

//GET
func (agentsUsecase agentsUsecase) GetAgentsForPublic(request request.AgentsRequest) {
	agentss := agentsUsecase.AgentsRepository.GetAgentsForPublic(request)
	fmt.Println(agentss)
}

func (agentsUsecase agentsUsecase) GetAgentsForInternal(request request.AgentsRequest) {
	agentss := agentsUsecase.AgentsRepository.GetAgentsForInternal(request)
	fmt.Println(agentss)
}

func (agentsUsecase agentsUsecase) GetAgentsForPrivate(request request.AgentsRequest) {
	agentss := agentsUsecase.AgentsRepository.GetAgentsForPrivate(request)
	fmt.Println(agentss)
}

//CREATE
func (agentsUsecase agentsUsecase) CreateAgentsForPublic(request request.AgentsRequest) {
	agentss := agentsUsecase.AgentsRepository.CreateAgentsForPublic(request)
	fmt.Println(agentss)
}

func (agentsUsecase agentsUsecase) CreateAgentsForInternal(request request.AgentsRequest) {
	agentss := agentsUsecase.AgentsRepository.CreateAgentsForInternal(request)
	fmt.Println(agentss)
}

func (agentsUsecase agentsUsecase) CreateAgentsForPrivate(request request.AgentsRequest) {
	agentss := agentsUsecase.AgentsRepository.CreateAgentsForPrivate(request)
	fmt.Println(agentss)
}

//UPDATE
func (agentsUsecase agentsUsecase) UpdateAgentsForPublic(request request.AgentsRequest) {
	agentss := agentsUsecase.AgentsRepository.UpdateAgentsForPublic(request)
	fmt.Println(agentss)
}

func (agentsUsecase agentsUsecase) UpdateAgentsForInternal(request request.AgentsRequest) {
	agentss := agentsUsecase.AgentsRepository.UpdateAgentsForInternal(request)
	fmt.Println(agentss)
}

func (agentsUsecase agentsUsecase) UpdateAgentsForPrivate(request request.AgentsRequest) {
	agentss := agentsUsecase.AgentsRepository.UpdateAgentsForPrivate(request)
	fmt.Println(agentss)
}

//DELETE
func (agentsUsecase agentsUsecase) DeleteAgentsForPublic(request request.AgentsRequest) {
	agentss := agentsUsecase.AgentsRepository.DeleteAgentsForPublic(request)
	fmt.Println(agentss)
}

func (agentsUsecase agentsUsecase) DeleteAgentsForInternal(request request.AgentsRequest) {
	agentss := agentsUsecase.AgentsRepository.DeleteAgentsForInternal(request)
	fmt.Println(agentss)
}

func (agentsUsecase agentsUsecase) DeleteAgentsForPrivate(request request.AgentsRequest) {
	agentss := agentsUsecase.AgentsRepository.DeleteAgentsForPrivate(request)
	fmt.Println(agentss)
}

func NewAgentsUsecase(agentsRepository repository.AgentsRepository) AgentsUsecase {
	return &agentsUsecase{ AgentsRepository: agentsRepository}
}