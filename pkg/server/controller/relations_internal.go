package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/visualizer/pkg/entity/model"
	"github.com/ryo-arima/visualizer/pkg/entity/request"
	"github.com/ryo-arima/visualizer/pkg/entity/response"
	"github.com/ryo-arima/visualizer/pkg/server/repository"
)

type RelationsControllerForInternal interface {
	GetRelationss(c *gin.Context)
	CreateRelations(c *gin.Context)
	UpdateRelations(c *gin.Context)
	DeleteRelations(c *gin.Context)
}

type relationsControllerForInternal struct {
	RelationsRepository repository.RelationsRepository
}

func (relationsController relationsControllerForInternal) GetRelationss(c *gin.Context) {
	var relationsRequest request.RelationsRequest
	if err := c.Bind(&relationsRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.RelationsResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Relationss: []response.Relations{}})
		return
	}
	res := relationsController.RelationsRepository.GetRelationss()
	c.JSON(http.StatusOK, res)
	return
}


func (relationsController relationsControllerForInternal) CreateRelations(c *gin.Context) {
	var relationsRequest request.RelationsRequest
	if err := c.Bind(&relationsRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.RelationsResponse{Code: "SERVER_CONTROLLER_CREATE__FOR__001", Message: err.Error(), Relationss: []response.Relations{}})
		return
	}
	var relationsModel model.Relationss
	res := relationsController.RelationsRepository.CreateRelations(relationsModel)
	c.JSON(http.StatusOK, res)
	return
}


func (relationsController relationsControllerForInternal) UpdateRelations(c *gin.Context) {
	var relationsRequest request.RelationsRequest
	if err := c.Bind(&relationsRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.RelationsResponse{Code: "SERVER_CONTROLLER_UPDATE__FOR__001", Message: err.Error(), Relationss: []response.Relations{}})
		return
	}
	var relationsModel model.Relationss
	res := relationsController.RelationsRepository.UpdateRelations(relationsModel)
	c.JSON(http.StatusOK, res)
	return
}


func (relationsController relationsControllerForInternal) DeleteRelations(c *gin.Context) {
	var relationsRequest request.RelationsRequest
	if err := c.Bind(&relationsRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.RelationsResponse{Code: "SERVER_CONTROLLER_DELETE__FOR__001", Message: err.Error(), Relationss: []response.Relations{}})
		return
	}
	var uuid string
	res := relationsController.RelationsRepository.DeleteRelations(uuid)
	c.JSON(http.StatusOK, res)
	return
}


func NewRelationsControllerForInternal(relationsRepository repository.RelationsRepository) RelationsControllerForInternal {
	return &relationsControllerForInternal{ RelationsRepository: relationsRepository}
}