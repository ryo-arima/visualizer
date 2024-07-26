package repository

import (
	"time"

	"github.com/ryo-arima/visualizer/pkg/config"
	"github.com/ryo-arima/visualizer/pkg/entity/model"
	"gorm.io/gorm"
)

type SpacesRepository interface {
	GetSpacess() []model.Spacess
	CreateSpaces(spaces model.Spacess) *gorm.DB
	UpdateSpaces(spaces model.Spacess) *gorm.DB
	DeleteSpaces(uuid string) *gorm.DB
}

type spacesRepository struct {
	BaseConfig config.BaseConfig
}


func (spacesRepository spacesRepository) GetSpacess() []model.Spacess {
	var spacess []model.Spacess
	spacesRepository.BaseConfig.DBConnection.Find(&spacess)
	return spacess
}

func (spacesRepository spacesRepository) CreateSpaces(spaces model.Spacess) *gorm.DB {
	results := spacesRepository.BaseConfig.DBConnection.Create(&spaces)
	return results
}

func (spacesRepository spacesRepository) UpdateSpaces(spaces model.Spacess) *gorm.DB {
	results := spacesRepository.BaseConfig.DBConnection.Model(model.Spacess{}).Where("uuid = ?", spaces.UUID).Updates(spaces)
	return results
}

func (spacesRepository spacesRepository) DeleteSpaces(uuid string) *gorm.DB {
	results := spacesRepository.BaseConfig.DBConnection.Model(model.Spacess{}).Where("uuid = ?", uuid).Update("deleted_at", time.Now())
	return results
}

func NewSpacesRepository(conf config.BaseConfig) SpacesRepository {
	return &spacesRepository{BaseConfig: conf}
}