package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/visualizer/pkg/entity/model"
	"github.com/ryo-arima/visualizer/pkg/entity/request"
	"github.com/ryo-arima/visualizer/pkg/entity/response"
	"github.com/ryo-arima/visualizer/pkg/server/repository"
)

type ElementsControllerForInternal interface {
	GetElementss(c *gin.Context)
	CreateElements(c *gin.Context)
	UpdateElements(c *gin.Context)
	DeleteElements(c *gin.Context)
}

type elementsControllerForInternal struct {
	ElementsRepository repository.ElementsRepository
}

func (elementsController elementsControllerForInternal) GetElementss(c *gin.Context) {
	var elementsRequest request.ElementsRequest
	if err := c.Bind(&elementsRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.ElementsResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Elementss: []response.Elements{}})
		return
	}
	res := elementsController.ElementsRepository.GetElementss()
	c.JSON(http.StatusOK, res)
	return
}


func (elementsController elementsControllerForInternal) CreateElements(c *gin.Context) {
	var elementsRequest request.ElementsRequest
	if err := c.Bind(&elementsRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.ElementsResponse{Code: "SERVER_CONTROLLER_CREATE__FOR__001", Message: err.Error(), Elementss: []response.Elements{}})
		return
	}
	var elementsModel model.Elementss
	res := elementsController.ElementsRepository.CreateElements(elementsModel)
	c.JSON(http.StatusOK, res)
	return
}


func (elementsController elementsControllerForInternal) UpdateElements(c *gin.Context) {
	var elementsRequest request.ElementsRequest
	if err := c.Bind(&elementsRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.ElementsResponse{Code: "SERVER_CONTROLLER_UPDATE__FOR__001", Message: err.Error(), Elementss: []response.Elements{}})
		return
	}
	var elementsModel model.Elementss
	res := elementsController.ElementsRepository.UpdateElements(elementsModel)
	c.JSON(http.StatusOK, res)
	return
}


func (elementsController elementsControllerForInternal) DeleteElements(c *gin.Context) {
	var elementsRequest request.ElementsRequest
	if err := c.Bind(&elementsRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.ElementsResponse{Code: "SERVER_CONTROLLER_DELETE__FOR__001", Message: err.Error(), Elementss: []response.Elements{}})
		return
	}
	var uuid string
	res := elementsController.ElementsRepository.DeleteElements(uuid)
	c.JSON(http.StatusOK, res)
	return
}


func NewElementsControllerForInternal(elementsRepository repository.ElementsRepository) ElementsControllerForInternal {
	return &elementsControllerForInternal{ ElementsRepository: elementsRepository}
}