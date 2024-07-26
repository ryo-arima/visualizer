package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/visualizer/pkg/entity/request"
	"github.com/ryo-arima/visualizer/pkg/entity/response"
	"github.com/ryo-arima/visualizer/pkg/server/repository"
)

type RelationsControllerForPublic interface {
	GetRelationss(c *gin.Context)
}

type relationsControllerForPublic struct {
	RelationsRepository repository.RelationsRepository
}

func (relationsController relationsControllerForPublic) GetRelationss(c *gin.Context) {
	var relationsRequest request.RelationsRequest
	if err := c.Bind(&relationsRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.RelationsResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Relationss: []response.Relations{}})
		return
	}
	res := relationsController.RelationsRepository.GetRelationss()
	c.JSON(http.StatusOK, res)
	return
}


func NewRelationsControllerForPublic(relationsRepository repository.RelationsRepository) RelationsControllerForPublic {
	return &relationsControllerForPublic{ RelationsRepository: relationsRepository}
}