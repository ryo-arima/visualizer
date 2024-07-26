package repository

import (
	"time"

	"github.com/ryo-arima/visualizer/pkg/config"
	"github.com/ryo-arima/visualizer/pkg/entity/model"
	"gorm.io/gorm"
)

type ElementsRepository interface {
	GetElementss() []model.Elementss
	CreateElements(elements model.Elementss) *gorm.DB
	UpdateElements(elements model.Elementss) *gorm.DB
	DeleteElements(uuid string) *gorm.DB
}

type elementsRepository struct {
	BaseConfig config.BaseConfig
}


func (elementsRepository elementsRepository) GetElementss() []model.Elementss {
	var elementss []model.Elementss
	elementsRepository.BaseConfig.DBConnection.Find(&elementss)
	return elementss
}

func (elementsRepository elementsRepository) CreateElements(elements model.Elementss) *gorm.DB {
	results := elementsRepository.BaseConfig.DBConnection.Create(&elements)
	return results
}

func (elementsRepository elementsRepository) UpdateElements(elements model.Elementss) *gorm.DB {
	results := elementsRepository.BaseConfig.DBConnection.Model(model.Elementss{}).Where("uuid = ?", elements.UUID).Updates(elements)
	return results
}

func (elementsRepository elementsRepository) DeleteElements(uuid string) *gorm.DB {
	results := elementsRepository.BaseConfig.DBConnection.Model(model.Elementss{}).Where("uuid = ?", uuid).Update("deleted_at", time.Now())
	return results
}

func NewElementsRepository(conf config.BaseConfig) ElementsRepository {
	return &elementsRepository{BaseConfig: conf}
}