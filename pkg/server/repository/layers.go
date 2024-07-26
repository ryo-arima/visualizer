package repository

import (
	"time"

	"github.com/ryo-arima/visualizer/pkg/config"
	"github.com/ryo-arima/visualizer/pkg/entity/model"
	"gorm.io/gorm"
)

type LayersRepository interface {
	GetLayerss() []model.Layerss
	CreateLayers(layers model.Layerss) *gorm.DB
	UpdateLayers(layers model.Layerss) *gorm.DB
	DeleteLayers(uuid string) *gorm.DB
}

type layersRepository struct {
	BaseConfig config.BaseConfig
}


func (layersRepository layersRepository) GetLayerss() []model.Layerss {
	var layerss []model.Layerss
	layersRepository.BaseConfig.DBConnection.Find(&layerss)
	return layerss
}

func (layersRepository layersRepository) CreateLayers(layers model.Layerss) *gorm.DB {
	results := layersRepository.BaseConfig.DBConnection.Create(&layers)
	return results
}

func (layersRepository layersRepository) UpdateLayers(layers model.Layerss) *gorm.DB {
	results := layersRepository.BaseConfig.DBConnection.Model(model.Layerss{}).Where("uuid = ?", layers.UUID).Updates(layers)
	return results
}

func (layersRepository layersRepository) DeleteLayers(uuid string) *gorm.DB {
	results := layersRepository.BaseConfig.DBConnection.Model(model.Layerss{}).Where("uuid = ?", uuid).Update("deleted_at", time.Now())
	return results
}

func NewLayersRepository(conf config.BaseConfig) LayersRepository {
	return &layersRepository{BaseConfig: conf}
}