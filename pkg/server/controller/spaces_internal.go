package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/visualizer/pkg/entity/model"
	"github.com/ryo-arima/visualizer/pkg/entity/request"
	"github.com/ryo-arima/visualizer/pkg/entity/response"
	"github.com/ryo-arima/visualizer/pkg/server/repository"
)

type SpacesControllerForInternal interface {
	GetSpacess(c *gin.Context)
	CreateSpaces(c *gin.Context)
	UpdateSpaces(c *gin.Context)
	DeleteSpaces(c *gin.Context)
}

type spacesControllerForInternal struct {
	SpacesRepository repository.SpacesRepository
}

func (spacesController spacesControllerForInternal) GetSpacess(c *gin.Context) {
	var spacesRequest request.SpacesRequest
	if err := c.Bind(&spacesRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.SpacesResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Spacess: []response.Spaces{}})
		return
	}
	res := spacesController.SpacesRepository.GetSpacess()
	c.JSON(http.StatusOK, res)
	return
}


func (spacesController spacesControllerForInternal) CreateSpaces(c *gin.Context) {
	var spacesRequest request.SpacesRequest
	if err := c.Bind(&spacesRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.SpacesResponse{Code: "SERVER_CONTROLLER_CREATE__FOR__001", Message: err.Error(), Spacess: []response.Spaces{}})
		return
	}
	var spacesModel model.Spacess
	res := spacesController.SpacesRepository.CreateSpaces(spacesModel)
	c.JSON(http.StatusOK, res)
	return
}


func (spacesController spacesControllerForInternal) UpdateSpaces(c *gin.Context) {
	var spacesRequest request.SpacesRequest
	if err := c.Bind(&spacesRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.SpacesResponse{Code: "SERVER_CONTROLLER_UPDATE__FOR__001", Message: err.Error(), Spacess: []response.Spaces{}})
		return
	}
	var spacesModel model.Spacess
	res := spacesController.SpacesRepository.UpdateSpaces(spacesModel)
	c.JSON(http.StatusOK, res)
	return
}


func (spacesController spacesControllerForInternal) DeleteSpaces(c *gin.Context) {
	var spacesRequest request.SpacesRequest
	if err := c.Bind(&spacesRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.SpacesResponse{Code: "SERVER_CONTROLLER_DELETE__FOR__001", Message: err.Error(), Spacess: []response.Spaces{}})
		return
	}
	var uuid string
	res := spacesController.SpacesRepository.DeleteSpaces(uuid)
	c.JSON(http.StatusOK, res)
	return
}


func NewSpacesControllerForInternal(spacesRepository repository.SpacesRepository) SpacesControllerForInternal {
	return &spacesControllerForInternal{ SpacesRepository: spacesRepository}
}