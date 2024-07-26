package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/visualizer/pkg/entity/request"
	"github.com/ryo-arima/visualizer/pkg/entity/response"
	"github.com/ryo-arima/visualizer/pkg/server/repository"
)

type LayersControllerForPublic interface {
	GetLayerss(c *gin.Context)
}

type layersControllerForPublic struct {
	LayersRepository repository.LayersRepository
}

func (layersController layersControllerForPublic) GetLayerss(c *gin.Context) {
	var layersRequest request.LayersRequest
	if err := c.Bind(&layersRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.LayersResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Layerss: []response.Layers{}})
		return
	}
	res := layersController.LayersRepository.GetLayerss()
	c.JSON(http.StatusOK, res)
	return
}


func NewLayersControllerForPublic(layersRepository repository.LayersRepository) LayersControllerForPublic {
	return &layersControllerForPublic{ LayersRepository: layersRepository}
}