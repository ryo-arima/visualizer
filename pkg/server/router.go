package server

import (
	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/visualizer/pkg/config"
	"github.com/ryo-arima/visualizer/pkg/server/controller"
	"github.com/ryo-arima/visualizer/pkg/server/middleware"
	"github.com/ryo-arima/visualizer/pkg/server/repository"
)

func InitRouter(conf config.BaseConfig) *gin.Engine {
	
	elementsRepository := repository.NewElementsRepository(conf)
	elementsControllerForPublic := controller.NewElementsControllerForPublic(elementsRepository)
	elementsControllerForInternal := controller.NewElementsControllerForInternal(elementsRepository)
	elementsControllerForPrivate := controller.NewElementsControllerForPrivate(elementsRepository)
	
	relationsRepository := repository.NewRelationsRepository(conf)
	relationsControllerForPublic := controller.NewRelationsControllerForPublic(relationsRepository)
	relationsControllerForInternal := controller.NewRelationsControllerForInternal(relationsRepository)
	relationsControllerForPrivate := controller.NewRelationsControllerForPrivate(relationsRepository)
	
	layersRepository := repository.NewLayersRepository(conf)
	layersControllerForPublic := controller.NewLayersControllerForPublic(layersRepository)
	layersControllerForInternal := controller.NewLayersControllerForInternal(layersRepository)
	layersControllerForPrivate := controller.NewLayersControllerForPrivate(layersRepository)
	
	spacesRepository := repository.NewSpacesRepository(conf)
	spacesControllerForPublic := controller.NewSpacesControllerForPublic(spacesRepository)
	spacesControllerForInternal := controller.NewSpacesControllerForInternal(spacesRepository)
	spacesControllerForPrivate := controller.NewSpacesControllerForPrivate(spacesRepository)
	
	
    router := gin.Default()
	privateAPI := router.Group("api/private")
	internalAPI := router.Group("api/internal")
	publicAPI := router.Group("api/public")

	publicAPI.Use(middleware.ForPublic(conf))
	internalAPI.Use(middleware.ForInternal(conf))
	privateAPI.Use(middleware.ForPrivate(conf))

	
	//elements
	publicAPI.GET("/elementss", elementsControllerForPublic.GetElementss)
	internalAPI.GET("/elementss", elementsControllerForInternal.GetElementss)
	internalAPI.POST("/elements", elementsControllerForInternal.CreateElements)
	internalAPI.PUT("/elements", elementsControllerForInternal.UpdateElements)
	internalAPI.DELETE("/elements", elementsControllerForInternal.DeleteElements)
	privateAPI.GET("/elementss", elementsControllerForPrivate.GetElementss)
	privateAPI.POST("/elements", elementsControllerForPrivate.CreateElements)
	privateAPI.PUT("/elements", elementsControllerForPrivate.UpdateElements)
	privateAPI.DELETE("/elements", elementsControllerForPrivate.DeleteElements)
	
	//relations
	publicAPI.GET("/relationss", relationsControllerForPublic.GetRelationss)
	internalAPI.GET("/relationss", relationsControllerForInternal.GetRelationss)
	internalAPI.POST("/relations", relationsControllerForInternal.CreateRelations)
	internalAPI.PUT("/relations", relationsControllerForInternal.UpdateRelations)
	internalAPI.DELETE("/relations", relationsControllerForInternal.DeleteRelations)
	privateAPI.GET("/relationss", relationsControllerForPrivate.GetRelationss)
	privateAPI.POST("/relations", relationsControllerForPrivate.CreateRelations)
	privateAPI.PUT("/relations", relationsControllerForPrivate.UpdateRelations)
	privateAPI.DELETE("/relations", relationsControllerForPrivate.DeleteRelations)
	
	//layers
	publicAPI.GET("/layerss", layersControllerForPublic.GetLayerss)
	internalAPI.GET("/layerss", layersControllerForInternal.GetLayerss)
	internalAPI.POST("/layers", layersControllerForInternal.CreateLayers)
	internalAPI.PUT("/layers", layersControllerForInternal.UpdateLayers)
	internalAPI.DELETE("/layers", layersControllerForInternal.DeleteLayers)
	privateAPI.GET("/layerss", layersControllerForPrivate.GetLayerss)
	privateAPI.POST("/layers", layersControllerForPrivate.CreateLayers)
	privateAPI.PUT("/layers", layersControllerForPrivate.UpdateLayers)
	privateAPI.DELETE("/layers", layersControllerForPrivate.DeleteLayers)
	
	//spaces
	publicAPI.GET("/spacess", spacesControllerForPublic.GetSpacess)
	internalAPI.GET("/spacess", spacesControllerForInternal.GetSpacess)
	internalAPI.POST("/spaces", spacesControllerForInternal.CreateSpaces)
	internalAPI.PUT("/spaces", spacesControllerForInternal.UpdateSpaces)
	internalAPI.DELETE("/spaces", spacesControllerForInternal.DeleteSpaces)
	privateAPI.GET("/spacess", spacesControllerForPrivate.GetSpacess)
	privateAPI.POST("/spaces", spacesControllerForPrivate.CreateSpaces)
	privateAPI.PUT("/spaces", spacesControllerForPrivate.UpdateSpaces)
	privateAPI.DELETE("/spaces", spacesControllerForPrivate.DeleteSpaces)
	

	return router
}