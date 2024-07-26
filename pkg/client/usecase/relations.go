package usecase

import (
	"fmt"

	"github.com/ryo-arima/visualizer/pkg/client/repository"
	"github.com/ryo-arima/visualizer/pkg/entity/request"
)

type RelationsUsecase interface {
	BootstrapRelationsForDB(request request.RelationsRequest)
	GetRelationsForPublic(request request.RelationsRequest)
	GetRelationsForInternal(request request.RelationsRequest)
	GetRelationsForPrivate(request request.RelationsRequest)
	CreateRelationsForPublic(request request.RelationsRequest)
	CreateRelationsForInternal(request request.RelationsRequest)
	CreateRelationsForPrivate(request request.RelationsRequest)
	UpdateRelationsForPublic(request request.RelationsRequest)
	UpdateRelationsForInternal(request request.RelationsRequest)
	UpdateRelationsForPrivate(request request.RelationsRequest)
	DeleteRelationsForPublic(request request.RelationsRequest)
	DeleteRelationsForInternal(request request.RelationsRequest)
	DeleteRelationsForPrivate(request request.RelationsRequest)
}

type relationsUsecase struct {
	RelationsRepository   repository.RelationsRepository
}

//Bootstrap
func (relationsUsecase relationsUsecase) BootstrapRelationsForDB(request request.RelationsRequest) {
	relationss := relationsUsecase.RelationsRepository.BootstrapRelationsForDB(request)
	fmt.Println(relationss)
}

//GET
func (relationsUsecase relationsUsecase) GetRelationsForPublic(request request.RelationsRequest) {
	relationss := relationsUsecase.RelationsRepository.GetRelationsForPublic(request)
	fmt.Println(relationss)
}

func (relationsUsecase relationsUsecase) GetRelationsForInternal(request request.RelationsRequest) {
	relationss := relationsUsecase.RelationsRepository.GetRelationsForInternal(request)
	fmt.Println(relationss)
}

func (relationsUsecase relationsUsecase) GetRelationsForPrivate(request request.RelationsRequest) {
	relationss := relationsUsecase.RelationsRepository.GetRelationsForPrivate(request)
	fmt.Println(relationss)
}

//CREATE
func (relationsUsecase relationsUsecase) CreateRelationsForPublic(request request.RelationsRequest) {
	relationss := relationsUsecase.RelationsRepository.CreateRelationsForPublic(request)
	fmt.Println(relationss)
}

func (relationsUsecase relationsUsecase) CreateRelationsForInternal(request request.RelationsRequest) {
	relationss := relationsUsecase.RelationsRepository.CreateRelationsForInternal(request)
	fmt.Println(relationss)
}

func (relationsUsecase relationsUsecase) CreateRelationsForPrivate(request request.RelationsRequest) {
	relationss := relationsUsecase.RelationsRepository.CreateRelationsForPrivate(request)
	fmt.Println(relationss)
}

//UPDATE
func (relationsUsecase relationsUsecase) UpdateRelationsForPublic(request request.RelationsRequest) {
	relationss := relationsUsecase.RelationsRepository.UpdateRelationsForPublic(request)
	fmt.Println(relationss)
}

func (relationsUsecase relationsUsecase) UpdateRelationsForInternal(request request.RelationsRequest) {
	relationss := relationsUsecase.RelationsRepository.UpdateRelationsForInternal(request)
	fmt.Println(relationss)
}

func (relationsUsecase relationsUsecase) UpdateRelationsForPrivate(request request.RelationsRequest) {
	relationss := relationsUsecase.RelationsRepository.UpdateRelationsForPrivate(request)
	fmt.Println(relationss)
}

//DELETE
func (relationsUsecase relationsUsecase) DeleteRelationsForPublic(request request.RelationsRequest) {
	relationss := relationsUsecase.RelationsRepository.DeleteRelationsForPublic(request)
	fmt.Println(relationss)
}

func (relationsUsecase relationsUsecase) DeleteRelationsForInternal(request request.RelationsRequest) {
	relationss := relationsUsecase.RelationsRepository.DeleteRelationsForInternal(request)
	fmt.Println(relationss)
}

func (relationsUsecase relationsUsecase) DeleteRelationsForPrivate(request request.RelationsRequest) {
	relationss := relationsUsecase.RelationsRepository.DeleteRelationsForPrivate(request)
	fmt.Println(relationss)
}

func NewRelationsUsecase(relationsRepository repository.RelationsRepository) RelationsUsecase {
	return &relationsUsecase{ RelationsRepository: relationsRepository}
}