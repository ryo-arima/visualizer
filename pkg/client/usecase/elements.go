package usecase

import (
	"fmt"

	"github.com/ryo-arima/visualizer/pkg/client/repository"
	"github.com/ryo-arima/visualizer/pkg/entity/request"
)

type ElementsUsecase interface {
	BootstrapElementsForDB(request request.ElementsRequest)
	GetElementsForPublic(request request.ElementsRequest)
	GetElementsForInternal(request request.ElementsRequest)
	GetElementsForPrivate(request request.ElementsRequest)
	CreateElementsForPublic(request request.ElementsRequest)
	CreateElementsForInternal(request request.ElementsRequest)
	CreateElementsForPrivate(request request.ElementsRequest)
	UpdateElementsForPublic(request request.ElementsRequest)
	UpdateElementsForInternal(request request.ElementsRequest)
	UpdateElementsForPrivate(request request.ElementsRequest)
	DeleteElementsForPublic(request request.ElementsRequest)
	DeleteElementsForInternal(request request.ElementsRequest)
	DeleteElementsForPrivate(request request.ElementsRequest)
}

type elementsUsecase struct {
	ElementsRepository   repository.ElementsRepository
}

//Bootstrap
func (elementsUsecase elementsUsecase) BootstrapElementsForDB(request request.ElementsRequest) {
	elementss := elementsUsecase.ElementsRepository.BootstrapElementsForDB(request)
	fmt.Println(elementss)
}

//GET
func (elementsUsecase elementsUsecase) GetElementsForPublic(request request.ElementsRequest) {
	elementss := elementsUsecase.ElementsRepository.GetElementsForPublic(request)
	fmt.Println(elementss)
}

func (elementsUsecase elementsUsecase) GetElementsForInternal(request request.ElementsRequest) {
	elementss := elementsUsecase.ElementsRepository.GetElementsForInternal(request)
	fmt.Println(elementss)
}

func (elementsUsecase elementsUsecase) GetElementsForPrivate(request request.ElementsRequest) {
	elementss := elementsUsecase.ElementsRepository.GetElementsForPrivate(request)
	fmt.Println(elementss)
}

//CREATE
func (elementsUsecase elementsUsecase) CreateElementsForPublic(request request.ElementsRequest) {
	elementss := elementsUsecase.ElementsRepository.CreateElementsForPublic(request)
	fmt.Println(elementss)
}

func (elementsUsecase elementsUsecase) CreateElementsForInternal(request request.ElementsRequest) {
	elementss := elementsUsecase.ElementsRepository.CreateElementsForInternal(request)
	fmt.Println(elementss)
}

func (elementsUsecase elementsUsecase) CreateElementsForPrivate(request request.ElementsRequest) {
	elementss := elementsUsecase.ElementsRepository.CreateElementsForPrivate(request)
	fmt.Println(elementss)
}

//UPDATE
func (elementsUsecase elementsUsecase) UpdateElementsForPublic(request request.ElementsRequest) {
	elementss := elementsUsecase.ElementsRepository.UpdateElementsForPublic(request)
	fmt.Println(elementss)
}

func (elementsUsecase elementsUsecase) UpdateElementsForInternal(request request.ElementsRequest) {
	elementss := elementsUsecase.ElementsRepository.UpdateElementsForInternal(request)
	fmt.Println(elementss)
}

func (elementsUsecase elementsUsecase) UpdateElementsForPrivate(request request.ElementsRequest) {
	elementss := elementsUsecase.ElementsRepository.UpdateElementsForPrivate(request)
	fmt.Println(elementss)
}

//DELETE
func (elementsUsecase elementsUsecase) DeleteElementsForPublic(request request.ElementsRequest) {
	elementss := elementsUsecase.ElementsRepository.DeleteElementsForPublic(request)
	fmt.Println(elementss)
}

func (elementsUsecase elementsUsecase) DeleteElementsForInternal(request request.ElementsRequest) {
	elementss := elementsUsecase.ElementsRepository.DeleteElementsForInternal(request)
	fmt.Println(elementss)
}

func (elementsUsecase elementsUsecase) DeleteElementsForPrivate(request request.ElementsRequest) {
	elementss := elementsUsecase.ElementsRepository.DeleteElementsForPrivate(request)
	fmt.Println(elementss)
}

func NewElementsUsecase(elementsRepository repository.ElementsRepository) ElementsUsecase {
	return &elementsUsecase{ ElementsRepository: elementsRepository}
}