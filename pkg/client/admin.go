package client

import (
	"github.com/ryo-arima/visualizer/pkg/client/controller"
	"github.com/ryo-arima/visualizer/pkg/config"
	"github.com/spf13/cobra"
)

type BaseCmdForAdminUser struct {
	Bootstrap *cobra.Command
	Create    *cobra.Command
	Get       *cobra.Command
	Update    *cobra.Command
	Delete    *cobra.Command
}

func InitRootCmdForAdminUser() *cobra.Command {
	var rootCmdForAdminUser = &cobra.Command{
		Use:   "visualizer-admin",
		Short: "'visualizer' is a CLI tool to manage anniversaries",
		Long:  `''`,
	}
	return rootCmdForAdminUser
}

func InitBaseCmdForAdminUser() BaseCmdForAdminUser {
	bootstrapCmd := &cobra.Command{
		Use:   "bootstrap",
		Short: "bootstrap the value of a key",
		Long:  "bootstrap the value of a key",
	}
	createCmd := &cobra.Command{
		Use:   "create",
		Short: "create the value of a key",
		Long:  "create the value of a key",
	}
	getCmd := &cobra.Command{
		Use:   "get",
		Short: "get the value of a key",
		Long:  "get the value of a key",
	}
	updateCmd := &cobra.Command{
		Use:   "update",
		Short: "update the value of a key",
		Long:  "update the value of a key",
	}
	deleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "delete the value of a key",
		Long:  "delete the value of a key",
	}
	baseCmdForAdminUser := BaseCmdForAdminUser{
		Bootstrap: bootstrapCmd,
		Create:    createCmd,
		Get:       getCmd,
		Update:    updateCmd,
		Delete:    deleteCmd,
	}
	return baseCmdForAdminUser
}

func ClientForAdminUser(conf config.BaseConfig) {
	rootCmdForAdminUser := InitRootCmdForAdminUser()
	rootCmdForAdminUser.CompletionOptions.HiddenDefaultCmd = true
	baseCmdForAdminUser := InitBaseCmdForAdminUser()

	//bootstrap
	bootstrapElementsCmdForAdminUser := controller.InitBootstrapElementsCmdForAdminUser(conf)
	baseCmdForAdminUser.Bootstrap.AddCommand(bootstrapElementsCmdForAdminUser)
	bootstrapRelationsCmdForAdminUser := controller.InitBootstrapRelationsCmdForAdminUser(conf)
	baseCmdForAdminUser.Bootstrap.AddCommand(bootstrapRelationsCmdForAdminUser)
	bootstrapLayersCmdForAdminUser := controller.InitBootstrapLayersCmdForAdminUser(conf)
	baseCmdForAdminUser.Bootstrap.AddCommand(bootstrapLayersCmdForAdminUser)
	bootstrapSpacesCmdForAdminUser := controller.InitBootstrapSpacesCmdForAdminUser(conf)
	baseCmdForAdminUser.Bootstrap.AddCommand(bootstrapSpacesCmdForAdminUser)
	rootCmdForAdminUser.AddCommand(baseCmdForAdminUser.Bootstrap)

	//create
	createElementsCmdForAdminUser := controller.InitCreateElementsCmdForAdminUser(conf)
	baseCmdForAdminUser.Create.AddCommand(createElementsCmdForAdminUser)
	createRelationsCmdForAdminUser := controller.InitCreateRelationsCmdForAdminUser(conf)
	baseCmdForAdminUser.Create.AddCommand(createRelationsCmdForAdminUser)
	createLayersCmdForAdminUser := controller.InitCreateLayersCmdForAdminUser(conf)
	baseCmdForAdminUser.Create.AddCommand(createLayersCmdForAdminUser)
	createSpacesCmdForAdminUser := controller.InitCreateSpacesCmdForAdminUser(conf)
	baseCmdForAdminUser.Create.AddCommand(createSpacesCmdForAdminUser)
	rootCmdForAdminUser.AddCommand(baseCmdForAdminUser.Create)

	//get
	getElementsCmdForAdminUser := controller.InitGetElementsCmdForAdminUser(conf)
	baseCmdForAdminUser.Get.AddCommand(getElementsCmdForAdminUser)
	getRelationsCmdForAdminUser := controller.InitGetRelationsCmdForAdminUser(conf)
	baseCmdForAdminUser.Get.AddCommand(getRelationsCmdForAdminUser)
	getLayersCmdForAdminUser := controller.InitGetLayersCmdForAdminUser(conf)
	baseCmdForAdminUser.Get.AddCommand(getLayersCmdForAdminUser)
	getSpacesCmdForAdminUser := controller.InitGetSpacesCmdForAdminUser(conf)
	baseCmdForAdminUser.Get.AddCommand(getSpacesCmdForAdminUser)
	rootCmdForAdminUser.AddCommand(baseCmdForAdminUser.Get)

	//update
	updateElementsCmdForAdminUser := controller.InitUpdateElementsCmdForAdminUser(conf)
	baseCmdForAdminUser.Update.AddCommand(updateElementsCmdForAdminUser)
	updateRelationsCmdForAdminUser := controller.InitUpdateRelationsCmdForAdminUser(conf)
	baseCmdForAdminUser.Update.AddCommand(updateRelationsCmdForAdminUser)
	updateLayersCmdForAdminUser := controller.InitUpdateLayersCmdForAdminUser(conf)
	baseCmdForAdminUser.Update.AddCommand(updateLayersCmdForAdminUser)
	updateSpacesCmdForAdminUser := controller.InitUpdateSpacesCmdForAdminUser(conf)
	baseCmdForAdminUser.Update.AddCommand(updateSpacesCmdForAdminUser)
	rootCmdForAdminUser.AddCommand(baseCmdForAdminUser.Update)
	
	//delete
	deleteElementsCmdForAdminUser := controller.InitDeleteElementsCmdForAdminUser(conf)
	baseCmdForAdminUser.Delete.AddCommand(deleteElementsCmdForAdminUser)
	deleteRelationsCmdForAdminUser := controller.InitDeleteRelationsCmdForAdminUser(conf)
	baseCmdForAdminUser.Delete.AddCommand(deleteRelationsCmdForAdminUser)
	deleteLayersCmdForAdminUser := controller.InitDeleteLayersCmdForAdminUser(conf)
	baseCmdForAdminUser.Delete.AddCommand(deleteLayersCmdForAdminUser)
	deleteSpacesCmdForAdminUser := controller.InitDeleteSpacesCmdForAdminUser(conf)
	baseCmdForAdminUser.Delete.AddCommand(deleteSpacesCmdForAdminUser)
	rootCmdForAdminUser.AddCommand(baseCmdForAdminUser.Delete)

	rootCmdForAdminUser.Execute()
}