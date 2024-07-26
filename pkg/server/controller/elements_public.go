package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/visualizer/pkg/entity/request"
	"github.com/ryo-arima/visualizer/pkg/entity/response"
	"github.com/ryo-arima/visualizer/pkg/server/repository"
)

type ElementsControllerForPublic interface {
	GetElementss(c *gin.Context)
}

type elementsControllerForPublic struct {
	ElementsRepository repository.ElementsRepository
}

func (elementsController elementsControllerForPublic) GetElementss(c *gin.Context) {
	var elementsRequest request.ElementsRequest
	if err := c.Bind(&elementsRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.ElementsResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Elementss: []response.Elements{}})
		return
	}
	res := elementsController.ElementsRepository.GetElementss()
	c.JSON(http.StatusOK, res)
	return
}


func NewElementsControllerForPublic(elementsRepository repository.ElementsRepository) ElementsControllerForPublic {
	return &elementsControllerForPublic{ ElementsRepository: elementsRepository}
}