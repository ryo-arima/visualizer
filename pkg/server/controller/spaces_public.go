package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/visualizer/pkg/entity/request"
	"github.com/ryo-arima/visualizer/pkg/entity/response"
	"github.com/ryo-arima/visualizer/pkg/server/repository"
)

type SpacesControllerForPublic interface {
	GetSpacess(c *gin.Context)
}

type spacesControllerForPublic struct {
	SpacesRepository repository.SpacesRepository
}

func (spacesController spacesControllerForPublic) GetSpacess(c *gin.Context) {
	var spacesRequest request.SpacesRequest
	if err := c.Bind(&spacesRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.SpacesResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Spacess: []response.Spaces{}})
		return
	}
	res := spacesController.SpacesRepository.GetSpacess()
	c.JSON(http.StatusOK, res)
	return
}


func NewSpacesControllerForPublic(spacesRepository repository.SpacesRepository) SpacesControllerForPublic {
	return &spacesControllerForPublic{ SpacesRepository: spacesRepository}
}