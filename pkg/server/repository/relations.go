package repository

import (
	"time"

	"github.com/ryo-arima/visualizer/pkg/config"
	"github.com/ryo-arima/visualizer/pkg/entity/model"
	"gorm.io/gorm"
)

type RelationsRepository interface {
	GetRelationss() []model.Relationss
	CreateRelations(relations model.Relationss) *gorm.DB
	UpdateRelations(relations model.Relationss) *gorm.DB
	DeleteRelations(uuid string) *gorm.DB
}

type relationsRepository struct {
	BaseConfig config.BaseConfig
}


func (relationsRepository relationsRepository) GetRelationss() []model.Relationss {
	var relationss []model.Relationss
	relationsRepository.BaseConfig.DBConnection.Find(&relationss)
	return relationss
}

func (relationsRepository relationsRepository) CreateRelations(relations model.Relationss) *gorm.DB {
	results := relationsRepository.BaseConfig.DBConnection.Create(&relations)
	return results
}

func (relationsRepository relationsRepository) UpdateRelations(relations model.Relationss) *gorm.DB {
	results := relationsRepository.BaseConfig.DBConnection.Model(model.Relationss{}).Where("uuid = ?", relations.UUID).Updates(relations)
	return results
}

func (relationsRepository relationsRepository) DeleteRelations(uuid string) *gorm.DB {
	results := relationsRepository.BaseConfig.DBConnection.Model(model.Relationss{}).Where("uuid = ?", uuid).Update("deleted_at", time.Now())
	return results
}

func NewRelationsRepository(conf config.BaseConfig) RelationsRepository {
	return &relationsRepository{BaseConfig: conf}
}