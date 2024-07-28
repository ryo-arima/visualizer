package repository

import (
	"fmt"

	"github.com/ryo-arima/visualizer/pkg/config"
	"github.com/ryo-arima/visualizer/pkg/entity/request"
	"github.com/ryo-arima/visualizer/pkg/entity/response"
)

type AgentsRepository interface {
	BootstrapAgentsForDB(request request.AgentsRequest)(response response.AgentsResponse) 
	GetAgentsForPublic(request request.AgentsRequest)(response response.AgentsResponse) 
	GetAgentsForInternal(request request.AgentsRequest)(response response.AgentsResponse) 
	GetAgentsForPrivate(request request.AgentsRequest)(response response.AgentsResponse) 
	CreateAgentsForPublic(request request.AgentsRequest)   (response response.AgentsResponse) 
	CreateAgentsForInternal(request request.AgentsRequest) (response response.AgentsResponse)
	CreateAgentsForPrivate(request request.AgentsRequest)  (response response.AgentsResponse)
	UpdateAgentsForPublic(request request.AgentsRequest)   (response response.AgentsResponse)
	UpdateAgentsForInternal(request request.AgentsRequest) (response response.AgentsResponse)
	UpdateAgentsForPrivate(request request.AgentsRequest)  (response response.AgentsResponse)
	DeleteAgentsForPublic(request request.AgentsRequest)   (response response.AgentsResponse)
	DeleteAgentsForInternal(request request.AgentsRequest) (response response.AgentsResponse)
	DeleteAgentsForPrivate(request request.AgentsRequest)  (response response.AgentsResponse)
}

type agentsRepository struct {
	BaseConfig   config.BaseConfig
}

//Bootstrap
func (agentsRepository agentsRepository) BootstrapAgentsForDB(request request.AgentsRequest) (response response.AgentsResponse) {
	fmt.Println("BootstrapAgentsForDB")
	return response
}

//GET
func (agentsRepository agentsRepository) GetAgentsForPublic(request request.AgentsRequest) (response response.AgentsResponse) {
	fmt.Println("GetAgentsForPublic")
	return response
}

func (agentsRepository agentsRepository) GetAgentsForInternal(request request.AgentsRequest) (response response.AgentsResponse ){
	fmt.Println("GetAgentsForInternal")
	return response
}

func (agentsRepository agentsRepository) GetAgentsForPrivate(request request.AgentsRequest) (response response.AgentsResponse) {
	fmt.Println("GetAgentsForPrivate")
	return response
}

//CREATE
func (agentsRepository agentsRepository) CreateAgentsForPublic(request request.AgentsRequest) (response response.AgentsResponse ){
	fmt.Println("CreateAgentsForPublic")
	return response
}

func (agentsRepository agentsRepository) CreateAgentsForInternal(request request.AgentsRequest) (response response.AgentsResponse) {
	fmt.Println("CreateAgentsForInternal()")
	return response
}

func (agentsRepository agentsRepository) CreateAgentsForPrivate(request request.AgentsRequest) (response response.AgentsResponse){
	fmt.Println("CreateAgentsForPrivate()")
	return response
}

//UPDATE
func (agentsRepository agentsRepository) UpdateAgentsForPublic(request request.AgentsRequest) (response response.AgentsResponse){
	fmt.Println("UpdateAgentsForPublic()")
	return response
}

func (agentsRepository agentsRepository) UpdateAgentsForInternal(request request.AgentsRequest) (response response.AgentsResponse){
	fmt.Println("UpdateAgentsForInternal")
	return response
}

func (agentsRepository agentsRepository) UpdateAgentsForPrivate(request request.AgentsRequest) (response response.AgentsResponse){
	fmt.Println("UpdateAgentsForPrivate")
	return response
}

//DELETE
func (agentsRepository agentsRepository) DeleteAgentsForPublic(request request.AgentsRequest) (response response.AgentsResponse){
	fmt.Println("DeleteAgentsForPublic")
	return response
}

func (agentsRepository agentsRepository) DeleteAgentsForInternal(request request.AgentsRequest) (response response.AgentsResponse ){
	fmt.Println("DeleteAgentsForInternal")
	return response
}

func (agentsRepository agentsRepository) DeleteAgentsForPrivate(request request.AgentsRequest) (response response.AgentsResponse){
	fmt.Println("DeleteAgentsForPrivate")
	return response
}

func NewAgentsRepository(conf config.BaseConfig) AgentsRepository {
	return &agentsRepository{BaseConfig: conf}
}