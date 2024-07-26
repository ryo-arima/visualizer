package usecase

import (
	"fmt"

	"github.com/ryo-arima/visualizer/pkg/client/repository"
	"github.com/ryo-arima/visualizer/pkg/entity/request"
)

type SpacesUsecase interface {
	BootstrapSpacesForDB(request request.SpacesRequest)
	GetSpacesForPublic(request request.SpacesRequest)
	GetSpacesForInternal(request request.SpacesRequest)
	GetSpacesForPrivate(request request.SpacesRequest)
	CreateSpacesForPublic(request request.SpacesRequest)
	CreateSpacesForInternal(request request.SpacesRequest)
	CreateSpacesForPrivate(request request.SpacesRequest)
	UpdateSpacesForPublic(request request.SpacesRequest)
	UpdateSpacesForInternal(request request.SpacesRequest)
	UpdateSpacesForPrivate(request request.SpacesRequest)
	DeleteSpacesForPublic(request request.SpacesRequest)
	DeleteSpacesForInternal(request request.SpacesRequest)
	DeleteSpacesForPrivate(request request.SpacesRequest)
}

type spacesUsecase struct {
	SpacesRepository   repository.SpacesRepository
}

//Bootstrap
func (spacesUsecase spacesUsecase) BootstrapSpacesForDB(request request.SpacesRequest) {
	spacess := spacesUsecase.SpacesRepository.BootstrapSpacesForDB(request)
	fmt.Println(spacess)
}

//GET
func (spacesUsecase spacesUsecase) GetSpacesForPublic(request request.SpacesRequest) {
	spacess := spacesUsecase.SpacesRepository.GetSpacesForPublic(request)
	fmt.Println(spacess)
}

func (spacesUsecase spacesUsecase) GetSpacesForInternal(request request.SpacesRequest) {
	spacess := spacesUsecase.SpacesRepository.GetSpacesForInternal(request)
	fmt.Println(spacess)
}

func (spacesUsecase spacesUsecase) GetSpacesForPrivate(request request.SpacesRequest) {
	spacess := spacesUsecase.SpacesRepository.GetSpacesForPrivate(request)
	fmt.Println(spacess)
}

//CREATE
func (spacesUsecase spacesUsecase) CreateSpacesForPublic(request request.SpacesRequest) {
	spacess := spacesUsecase.SpacesRepository.CreateSpacesForPublic(request)
	fmt.Println(spacess)
}

func (spacesUsecase spacesUsecase) CreateSpacesForInternal(request request.SpacesRequest) {
	spacess := spacesUsecase.SpacesRepository.CreateSpacesForInternal(request)
	fmt.Println(spacess)
}

func (spacesUsecase spacesUsecase) CreateSpacesForPrivate(request request.SpacesRequest) {
	spacess := spacesUsecase.SpacesRepository.CreateSpacesForPrivate(request)
	fmt.Println(spacess)
}

//UPDATE
func (spacesUsecase spacesUsecase) UpdateSpacesForPublic(request request.SpacesRequest) {
	spacess := spacesUsecase.SpacesRepository.UpdateSpacesForPublic(request)
	fmt.Println(spacess)
}

func (spacesUsecase spacesUsecase) UpdateSpacesForInternal(request request.SpacesRequest) {
	spacess := spacesUsecase.SpacesRepository.UpdateSpacesForInternal(request)
	fmt.Println(spacess)
}

func (spacesUsecase spacesUsecase) UpdateSpacesForPrivate(request request.SpacesRequest) {
	spacess := spacesUsecase.SpacesRepository.UpdateSpacesForPrivate(request)
	fmt.Println(spacess)
}

//DELETE
func (spacesUsecase spacesUsecase) DeleteSpacesForPublic(request request.SpacesRequest) {
	spacess := spacesUsecase.SpacesRepository.DeleteSpacesForPublic(request)
	fmt.Println(spacess)
}

func (spacesUsecase spacesUsecase) DeleteSpacesForInternal(request request.SpacesRequest) {
	spacess := spacesUsecase.SpacesRepository.DeleteSpacesForInternal(request)
	fmt.Println(spacess)
}

func (spacesUsecase spacesUsecase) DeleteSpacesForPrivate(request request.SpacesRequest) {
	spacess := spacesUsecase.SpacesRepository.DeleteSpacesForPrivate(request)
	fmt.Println(spacess)
}

func NewSpacesUsecase(spacesRepository repository.SpacesRepository) SpacesUsecase {
	return &spacesUsecase{ SpacesRepository: spacesRepository}
}