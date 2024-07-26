package usecase

import (
	"fmt"

	"github.com/ryo-arima/visualizer/pkg/client/repository"
	"github.com/ryo-arima/visualizer/pkg/entity/request"
)

type LayersUsecase interface {
	BootstrapLayersForDB(request request.LayersRequest)
	GetLayersForPublic(request request.LayersRequest)
	GetLayersForInternal(request request.LayersRequest)
	GetLayersForPrivate(request request.LayersRequest)
	CreateLayersForPublic(request request.LayersRequest)
	CreateLayersForInternal(request request.LayersRequest)
	CreateLayersForPrivate(request request.LayersRequest)
	UpdateLayersForPublic(request request.LayersRequest)
	UpdateLayersForInternal(request request.LayersRequest)
	UpdateLayersForPrivate(request request.LayersRequest)
	DeleteLayersForPublic(request request.LayersRequest)
	DeleteLayersForInternal(request request.LayersRequest)
	DeleteLayersForPrivate(request request.LayersRequest)
}

type layersUsecase struct {
	LayersRepository   repository.LayersRepository
}

//Bootstrap
func (layersUsecase layersUsecase) BootstrapLayersForDB(request request.LayersRequest) {
	layerss := layersUsecase.LayersRepository.BootstrapLayersForDB(request)
	fmt.Println(layerss)
}

//GET
func (layersUsecase layersUsecase) GetLayersForPublic(request request.LayersRequest) {
	layerss := layersUsecase.LayersRepository.GetLayersForPublic(request)
	fmt.Println(layerss)
}

func (layersUsecase layersUsecase) GetLayersForInternal(request request.LayersRequest) {
	layerss := layersUsecase.LayersRepository.GetLayersForInternal(request)
	fmt.Println(layerss)
}

func (layersUsecase layersUsecase) GetLayersForPrivate(request request.LayersRequest) {
	layerss := layersUsecase.LayersRepository.GetLayersForPrivate(request)
	fmt.Println(layerss)
}

//CREATE
func (layersUsecase layersUsecase) CreateLayersForPublic(request request.LayersRequest) {
	layerss := layersUsecase.LayersRepository.CreateLayersForPublic(request)
	fmt.Println(layerss)
}

func (layersUsecase layersUsecase) CreateLayersForInternal(request request.LayersRequest) {
	layerss := layersUsecase.LayersRepository.CreateLayersForInternal(request)
	fmt.Println(layerss)
}

func (layersUsecase layersUsecase) CreateLayersForPrivate(request request.LayersRequest) {
	layerss := layersUsecase.LayersRepository.CreateLayersForPrivate(request)
	fmt.Println(layerss)
}

//UPDATE
func (layersUsecase layersUsecase) UpdateLayersForPublic(request request.LayersRequest) {
	layerss := layersUsecase.LayersRepository.UpdateLayersForPublic(request)
	fmt.Println(layerss)
}

func (layersUsecase layersUsecase) UpdateLayersForInternal(request request.LayersRequest) {
	layerss := layersUsecase.LayersRepository.UpdateLayersForInternal(request)
	fmt.Println(layerss)
}

func (layersUsecase layersUsecase) UpdateLayersForPrivate(request request.LayersRequest) {
	layerss := layersUsecase.LayersRepository.UpdateLayersForPrivate(request)
	fmt.Println(layerss)
}

//DELETE
func (layersUsecase layersUsecase) DeleteLayersForPublic(request request.LayersRequest) {
	layerss := layersUsecase.LayersRepository.DeleteLayersForPublic(request)
	fmt.Println(layerss)
}

func (layersUsecase layersUsecase) DeleteLayersForInternal(request request.LayersRequest) {
	layerss := layersUsecase.LayersRepository.DeleteLayersForInternal(request)
	fmt.Println(layerss)
}

func (layersUsecase layersUsecase) DeleteLayersForPrivate(request request.LayersRequest) {
	layerss := layersUsecase.LayersRepository.DeleteLayersForPrivate(request)
	fmt.Println(layerss)
}

func NewLayersUsecase(layersRepository repository.LayersRepository) LayersUsecase {
	return &layersUsecase{ LayersRepository: layersRepository}
}