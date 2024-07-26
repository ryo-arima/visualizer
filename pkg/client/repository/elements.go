package repository

import (
	"fmt"

	"github.com/ryo-arima/visualizer/pkg/config"
	"github.com/ryo-arima/visualizer/pkg/entity/request"
	"github.com/ryo-arima/visualizer/pkg/entity/response"
)

type ElementsRepository interface {
	BootstrapElementsForDB(request request.ElementsRequest)(response response.ElementsResponse) 
	GetElementsForPublic(request request.ElementsRequest)(response response.ElementsResponse) 
	GetElementsForInternal(request request.ElementsRequest)(response response.ElementsResponse) 
	GetElementsForPrivate(request request.ElementsRequest)(response response.ElementsResponse) 
	CreateElementsForPublic(request request.ElementsRequest)   (response response.ElementsResponse) 
	CreateElementsForInternal(request request.ElementsRequest) (response response.ElementsResponse)
	CreateElementsForPrivate(request request.ElementsRequest)  (response response.ElementsResponse)
	UpdateElementsForPublic(request request.ElementsRequest)   (response response.ElementsResponse)
	UpdateElementsForInternal(request request.ElementsRequest) (response response.ElementsResponse)
	UpdateElementsForPrivate(request request.ElementsRequest)  (response response.ElementsResponse)
	DeleteElementsForPublic(request request.ElementsRequest)   (response response.ElementsResponse)
	DeleteElementsForInternal(request request.ElementsRequest) (response response.ElementsResponse)
	DeleteElementsForPrivate(request request.ElementsRequest)  (response response.ElementsResponse)
}

type elementsRepository struct {
	BaseConfig   config.BaseConfig
}

//Bootstrap
func (elementsRepository elementsRepository) BootstrapElementsForDB(request request.ElementsRequest) (response response.ElementsResponse) {
	fmt.Println("BootstrapElementsForDB")
	return response
}

//GET
func (elementsRepository elementsRepository) GetElementsForPublic(request request.ElementsRequest) (response response.ElementsResponse) {
	fmt.Println("GetElementsForPublic")
	return response
}

func (elementsRepository elementsRepository) GetElementsForInternal(request request.ElementsRequest) (response response.ElementsResponse ){
	fmt.Println("GetElementsForInternal")
	return response
}

func (elementsRepository elementsRepository) GetElementsForPrivate(request request.ElementsRequest) (response response.ElementsResponse) {
	fmt.Println("GetElementsForPrivate")
	return response
}

//CREATE
func (elementsRepository elementsRepository) CreateElementsForPublic(request request.ElementsRequest) (response response.ElementsResponse ){
	fmt.Println("CreateElementsForPublic")
	return response
}

func (elementsRepository elementsRepository) CreateElementsForInternal(request request.ElementsRequest) (response response.ElementsResponse) {
	fmt.Println("CreateElementsForInternal()")
	return response
}

func (elementsRepository elementsRepository) CreateElementsForPrivate(request request.ElementsRequest) (response response.ElementsResponse){
	fmt.Println("CreateElementsForPrivate()")
	return response
}

//UPDATE
func (elementsRepository elementsRepository) UpdateElementsForPublic(request request.ElementsRequest) (response response.ElementsResponse){
	fmt.Println("UpdateElementsForPublic()")
	return response
}

func (elementsRepository elementsRepository) UpdateElementsForInternal(request request.ElementsRequest) (response response.ElementsResponse){
	fmt.Println("UpdateElementsForInternal")
	return response
}

func (elementsRepository elementsRepository) UpdateElementsForPrivate(request request.ElementsRequest) (response response.ElementsResponse){
	fmt.Println("UpdateElementsForPrivate")
	return response
}

//DELETE
func (elementsRepository elementsRepository) DeleteElementsForPublic(request request.ElementsRequest) (response response.ElementsResponse){
	fmt.Println("DeleteElementsForPublic")
	return response
}

func (elementsRepository elementsRepository) DeleteElementsForInternal(request request.ElementsRequest) (response response.ElementsResponse ){
	fmt.Println("DeleteElementsForInternal")
	return response
}

func (elementsRepository elementsRepository) DeleteElementsForPrivate(request request.ElementsRequest) (response response.ElementsResponse){
	fmt.Println("DeleteElementsForPrivate")
	return response
}

func NewElementsRepository(conf config.BaseConfig) ElementsRepository {
	return &elementsRepository{BaseConfig: conf}
}