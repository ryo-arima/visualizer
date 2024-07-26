package repository

import (
	"fmt"

	"github.com/ryo-arima/visualizer/pkg/config"
	"github.com/ryo-arima/visualizer/pkg/entity/request"
	"github.com/ryo-arima/visualizer/pkg/entity/response"
)

type SpacesRepository interface {
	BootstrapSpacesForDB(request request.SpacesRequest)(response response.SpacesResponse) 
	GetSpacesForPublic(request request.SpacesRequest)(response response.SpacesResponse) 
	GetSpacesForInternal(request request.SpacesRequest)(response response.SpacesResponse) 
	GetSpacesForPrivate(request request.SpacesRequest)(response response.SpacesResponse) 
	CreateSpacesForPublic(request request.SpacesRequest)   (response response.SpacesResponse) 
	CreateSpacesForInternal(request request.SpacesRequest) (response response.SpacesResponse)
	CreateSpacesForPrivate(request request.SpacesRequest)  (response response.SpacesResponse)
	UpdateSpacesForPublic(request request.SpacesRequest)   (response response.SpacesResponse)
	UpdateSpacesForInternal(request request.SpacesRequest) (response response.SpacesResponse)
	UpdateSpacesForPrivate(request request.SpacesRequest)  (response response.SpacesResponse)
	DeleteSpacesForPublic(request request.SpacesRequest)   (response response.SpacesResponse)
	DeleteSpacesForInternal(request request.SpacesRequest) (response response.SpacesResponse)
	DeleteSpacesForPrivate(request request.SpacesRequest)  (response response.SpacesResponse)
}

type spacesRepository struct {
	BaseConfig   config.BaseConfig
}

//Bootstrap
func (spacesRepository spacesRepository) BootstrapSpacesForDB(request request.SpacesRequest) (response response.SpacesResponse) {
	fmt.Println("BootstrapSpacesForDB")
	return response
}

//GET
func (spacesRepository spacesRepository) GetSpacesForPublic(request request.SpacesRequest) (response response.SpacesResponse) {
	fmt.Println("GetSpacesForPublic")
	return response
}

func (spacesRepository spacesRepository) GetSpacesForInternal(request request.SpacesRequest) (response response.SpacesResponse ){
	fmt.Println("GetSpacesForInternal")
	return response
}

func (spacesRepository spacesRepository) GetSpacesForPrivate(request request.SpacesRequest) (response response.SpacesResponse) {
	fmt.Println("GetSpacesForPrivate")
	return response
}

//CREATE
func (spacesRepository spacesRepository) CreateSpacesForPublic(request request.SpacesRequest) (response response.SpacesResponse ){
	fmt.Println("CreateSpacesForPublic")
	return response
}

func (spacesRepository spacesRepository) CreateSpacesForInternal(request request.SpacesRequest) (response response.SpacesResponse) {
	fmt.Println("CreateSpacesForInternal()")
	return response
}

func (spacesRepository spacesRepository) CreateSpacesForPrivate(request request.SpacesRequest) (response response.SpacesResponse){
	fmt.Println("CreateSpacesForPrivate()")
	return response
}

//UPDATE
func (spacesRepository spacesRepository) UpdateSpacesForPublic(request request.SpacesRequest) (response response.SpacesResponse){
	fmt.Println("UpdateSpacesForPublic()")
	return response
}

func (spacesRepository spacesRepository) UpdateSpacesForInternal(request request.SpacesRequest) (response response.SpacesResponse){
	fmt.Println("UpdateSpacesForInternal")
	return response
}

func (spacesRepository spacesRepository) UpdateSpacesForPrivate(request request.SpacesRequest) (response response.SpacesResponse){
	fmt.Println("UpdateSpacesForPrivate")
	return response
}

//DELETE
func (spacesRepository spacesRepository) DeleteSpacesForPublic(request request.SpacesRequest) (response response.SpacesResponse){
	fmt.Println("DeleteSpacesForPublic")
	return response
}

func (spacesRepository spacesRepository) DeleteSpacesForInternal(request request.SpacesRequest) (response response.SpacesResponse ){
	fmt.Println("DeleteSpacesForInternal")
	return response
}

func (spacesRepository spacesRepository) DeleteSpacesForPrivate(request request.SpacesRequest) (response response.SpacesResponse){
	fmt.Println("DeleteSpacesForPrivate")
	return response
}

func NewSpacesRepository(conf config.BaseConfig) SpacesRepository {
	return &spacesRepository{BaseConfig: conf}
}