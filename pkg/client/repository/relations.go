package repository

import (
	"fmt"

	"github.com/ryo-arima/visualizer/pkg/config"
	"github.com/ryo-arima/visualizer/pkg/entity/request"
	"github.com/ryo-arima/visualizer/pkg/entity/response"
)

type RelationsRepository interface {
	BootstrapRelationsForDB(request request.RelationsRequest)(response response.RelationsResponse) 
	GetRelationsForPublic(request request.RelationsRequest)(response response.RelationsResponse) 
	GetRelationsForInternal(request request.RelationsRequest)(response response.RelationsResponse) 
	GetRelationsForPrivate(request request.RelationsRequest)(response response.RelationsResponse) 
	CreateRelationsForPublic(request request.RelationsRequest)   (response response.RelationsResponse) 
	CreateRelationsForInternal(request request.RelationsRequest) (response response.RelationsResponse)
	CreateRelationsForPrivate(request request.RelationsRequest)  (response response.RelationsResponse)
	UpdateRelationsForPublic(request request.RelationsRequest)   (response response.RelationsResponse)
	UpdateRelationsForInternal(request request.RelationsRequest) (response response.RelationsResponse)
	UpdateRelationsForPrivate(request request.RelationsRequest)  (response response.RelationsResponse)
	DeleteRelationsForPublic(request request.RelationsRequest)   (response response.RelationsResponse)
	DeleteRelationsForInternal(request request.RelationsRequest) (response response.RelationsResponse)
	DeleteRelationsForPrivate(request request.RelationsRequest)  (response response.RelationsResponse)
}

type relationsRepository struct {
	BaseConfig   config.BaseConfig
}

//Bootstrap
func (relationsRepository relationsRepository) BootstrapRelationsForDB(request request.RelationsRequest) (response response.RelationsResponse) {
	fmt.Println("BootstrapRelationsForDB")
	return response
}

//GET
func (relationsRepository relationsRepository) GetRelationsForPublic(request request.RelationsRequest) (response response.RelationsResponse) {
	fmt.Println("GetRelationsForPublic")
	return response
}

func (relationsRepository relationsRepository) GetRelationsForInternal(request request.RelationsRequest) (response response.RelationsResponse ){
	fmt.Println("GetRelationsForInternal")
	return response
}

func (relationsRepository relationsRepository) GetRelationsForPrivate(request request.RelationsRequest) (response response.RelationsResponse) {
	fmt.Println("GetRelationsForPrivate")
	return response
}

//CREATE
func (relationsRepository relationsRepository) CreateRelationsForPublic(request request.RelationsRequest) (response response.RelationsResponse ){
	fmt.Println("CreateRelationsForPublic")
	return response
}

func (relationsRepository relationsRepository) CreateRelationsForInternal(request request.RelationsRequest) (response response.RelationsResponse) {
	fmt.Println("CreateRelationsForInternal()")
	return response
}

func (relationsRepository relationsRepository) CreateRelationsForPrivate(request request.RelationsRequest) (response response.RelationsResponse){
	fmt.Println("CreateRelationsForPrivate()")
	return response
}

//UPDATE
func (relationsRepository relationsRepository) UpdateRelationsForPublic(request request.RelationsRequest) (response response.RelationsResponse){
	fmt.Println("UpdateRelationsForPublic()")
	return response
}

func (relationsRepository relationsRepository) UpdateRelationsForInternal(request request.RelationsRequest) (response response.RelationsResponse){
	fmt.Println("UpdateRelationsForInternal")
	return response
}

func (relationsRepository relationsRepository) UpdateRelationsForPrivate(request request.RelationsRequest) (response response.RelationsResponse){
	fmt.Println("UpdateRelationsForPrivate")
	return response
}

//DELETE
func (relationsRepository relationsRepository) DeleteRelationsForPublic(request request.RelationsRequest) (response response.RelationsResponse){
	fmt.Println("DeleteRelationsForPublic")
	return response
}

func (relationsRepository relationsRepository) DeleteRelationsForInternal(request request.RelationsRequest) (response response.RelationsResponse ){
	fmt.Println("DeleteRelationsForInternal")
	return response
}

func (relationsRepository relationsRepository) DeleteRelationsForPrivate(request request.RelationsRequest) (response response.RelationsResponse){
	fmt.Println("DeleteRelationsForPrivate")
	return response
}

func NewRelationsRepository(conf config.BaseConfig) RelationsRepository {
	return &relationsRepository{BaseConfig: conf}
}