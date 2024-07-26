package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/visualizer/pkg/entity/model"
	"github.com/ryo-arima/visualizer/pkg/entity/request"
	"github.com/ryo-arima/visualizer/pkg/entity/response"
	"github.com/ryo-arima/visualizer/pkg/server/repository"
)

type LayersControllerForInternal interface {
	GetLayerss(c *gin.Context)
	CreateLayers(c *gin.Context)
	UpdateLayers(c *gin.Context)
	DeleteLayers(c *gin.Context)
}

type layersControllerForInternal struct {
	LayersRepository repository.LayersRepository
}

func (layersController layersControllerForInternal) GetLayerss(c *gin.Context) {
	var layersRequest request.LayersRequest
	if err := c.Bind(&layersRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.LayersResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Layerss: []response.Layers{}})
		return
	}
	res := layersController.LayersRepository.GetLayerss()
	c.JSON(http.StatusOK, res)
	return
}


func (layersController layersControllerForInternal) CreateLayers(c *gin.Context) {
	var layersRequest request.LayersRequest
	if err := c.Bind(&layersRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.LayersResponse{Code: "SERVER_CONTROLLER_CREATE__FOR__001", Message: err.Error(), Layerss: []response.Layers{}})
		return
	}
	var layersModel model.Layerss
	res := layersController.LayersRepository.CreateLayers(layersModel)
	c.JSON(http.StatusOK, res)
	return
}


func (layersController layersControllerForInternal) UpdateLayers(c *gin.Context) {
	var layersRequest request.LayersRequest
	if err := c.Bind(&layersRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.LayersResponse{Code: "SERVER_CONTROLLER_UPDATE__FOR__001", Message: err.Error(), Layerss: []response.Layers{}})
		return
	}
	var layersModel model.Layerss
	res := layersController.LayersRepository.UpdateLayers(layersModel)
	c.JSON(http.StatusOK, res)
	return
}


func (layersController layersControllerForInternal) DeleteLayers(c *gin.Context) {
	var layersRequest request.LayersRequest
	if err := c.Bind(&layersRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.LayersResponse{Code: "SERVER_CONTROLLER_DELETE__FOR__001", Message: err.Error(), Layerss: []response.Layers{}})
		return
	}
	var uuid string
	res := layersController.LayersRepository.DeleteLayers(uuid)
	c.JSON(http.StatusOK, res)
	return
}


func NewLayersControllerForInternal(layersRepository repository.LayersRepository) LayersControllerForInternal {
	return &layersControllerForInternal{ LayersRepository: layersRepository}
}