package repository

import (
	"fmt"

	"github.com/ryo-arima/visualizer/pkg/config"
	"github.com/ryo-arima/visualizer/pkg/entity/request"
	"github.com/ryo-arima/visualizer/pkg/entity/response"
)

type LayersRepository interface {
	BootstrapLayersForDB(request request.LayersRequest)(response response.LayersResponse) 
	GetLayersForPublic(request request.LayersRequest)(response response.LayersResponse) 
	GetLayersForInternal(request request.LayersRequest)(response response.LayersResponse) 
	GetLayersForPrivate(request request.LayersRequest)(response response.LayersResponse) 
	CreateLayersForPublic(request request.LayersRequest)   (response response.LayersResponse) 
	CreateLayersForInternal(request request.LayersRequest) (response response.LayersResponse)
	CreateLayersForPrivate(request request.LayersRequest)  (response response.LayersResponse)
	UpdateLayersForPublic(request request.LayersRequest)   (response response.LayersResponse)
	UpdateLayersForInternal(request request.LayersRequest) (response response.LayersResponse)
	UpdateLayersForPrivate(request request.LayersRequest)  (response response.LayersResponse)
	DeleteLayersForPublic(request request.LayersRequest)   (response response.LayersResponse)
	DeleteLayersForInternal(request request.LayersRequest) (response response.LayersResponse)
	DeleteLayersForPrivate(request request.LayersRequest)  (response response.LayersResponse)
}

type layersRepository struct {
	BaseConfig   config.BaseConfig
}

//Bootstrap
func (layersRepository layersRepository) BootstrapLayersForDB(request request.LayersRequest) (response response.LayersResponse) {
	fmt.Println("BootstrapLayersForDB")
	return response
}

//GET
func (layersRepository layersRepository) GetLayersForPublic(request request.LayersRequest) (response response.LayersResponse) {
	fmt.Println("GetLayersForPublic")
	return response
}

func (layersRepository layersRepository) GetLayersForInternal(request request.LayersRequest) (response response.LayersResponse ){
	fmt.Println("GetLayersForInternal")
	return response
}

func (layersRepository layersRepository) GetLayersForPrivate(request request.LayersRequest) (response response.LayersResponse) {
	fmt.Println("GetLayersForPrivate")
	return response
}

//CREATE
func (layersRepository layersRepository) CreateLayersForPublic(request request.LayersRequest) (response response.LayersResponse ){
	fmt.Println("CreateLayersForPublic")
	return response
}

func (layersRepository layersRepository) CreateLayersForInternal(request request.LayersRequest) (response response.LayersResponse) {
	fmt.Println("CreateLayersForInternal()")
	return response
}

func (layersRepository layersRepository) CreateLayersForPrivate(request request.LayersRequest) (response response.LayersResponse){
	fmt.Println("CreateLayersForPrivate()")
	return response
}

//UPDATE
func (layersRepository layersRepository) UpdateLayersForPublic(request request.LayersRequest) (response response.LayersResponse){
	fmt.Println("UpdateLayersForPublic()")
	return response
}

func (layersRepository layersRepository) UpdateLayersForInternal(request request.LayersRequest) (response response.LayersResponse){
	fmt.Println("UpdateLayersForInternal")
	return response
}

func (layersRepository layersRepository) UpdateLayersForPrivate(request request.LayersRequest) (response response.LayersResponse){
	fmt.Println("UpdateLayersForPrivate")
	return response
}

//DELETE
func (layersRepository layersRepository) DeleteLayersForPublic(request request.LayersRequest) (response response.LayersResponse){
	fmt.Println("DeleteLayersForPublic")
	return response
}

func (layersRepository layersRepository) DeleteLayersForInternal(request request.LayersRequest) (response response.LayersResponse ){
	fmt.Println("DeleteLayersForInternal")
	return response
}

func (layersRepository layersRepository) DeleteLayersForPrivate(request request.LayersRequest) (response response.LayersResponse){
	fmt.Println("DeleteLayersForPrivate")
	return response
}

func NewLayersRepository(conf config.BaseConfig) LayersRepository {
	return &layersRepository{BaseConfig: conf}
}