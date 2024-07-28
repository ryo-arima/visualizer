package client

import (
	"github.com/ryo-arima/visualizer/pkg/client/controller"
	"github.com/ryo-arima/visualizer/pkg/config"
	"github.com/spf13/cobra"
)

type BaseCmdForAppUser struct {
	Create    *cobra.Command
	Get       *cobra.Command
	Update    *cobra.Command
	Delete    *cobra.Command
}

func InitRootCmdForAppUser() *cobra.Command {
	var rootCmdForAppUser = &cobra.Command{
		Use:   "visualizer-app",
		Short: "'visualizer' is a CLI tool to manage anniversaries",
		Long:  `''`,
	}
	return rootCmdForAppUser
}

func InitBaseCmdForAppUser() BaseCmdForAppUser {
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
	baseCmdForAppUser := BaseCmdForAppUser{
		Create:    createCmd,
		Get:       getCmd,
		Update:    updateCmd,
		Delete:    deleteCmd,
	}
	return baseCmdForAppUser
}

func ClientForAppUser(conf config.BaseConfig) {
	rootCmdForAppUser := InitRootCmdForAppUser()
	rootCmdForAppUser.CompletionOptions.HiddenDefaultCmd = true
	baseCmdForAppUser := InitBaseCmdForAppUser()

	//create
	createAgentsCmdForAppUser := controller.InitCreateAgentsCmdForAppUser(conf)
	baseCmdForAppUser.Create.AddCommand(createAgentsCmdForAppUser)
	createElementsCmdForAppUser := controller.InitCreateElementsCmdForAppUser(conf)
	baseCmdForAppUser.Create.AddCommand(createElementsCmdForAppUser)
	createRelationsCmdForAppUser := controller.InitCreateRelationsCmdForAppUser(conf)
	baseCmdForAppUser.Create.AddCommand(createRelationsCmdForAppUser)
	createLayersCmdForAppUser := controller.InitCreateLayersCmdForAppUser(conf)
	baseCmdForAppUser.Create.AddCommand(createLayersCmdForAppUser)
	createSpacesCmdForAppUser := controller.InitCreateSpacesCmdForAppUser(conf)
	baseCmdForAppUser.Create.AddCommand(createSpacesCmdForAppUser)
	rootCmdForAppUser.AddCommand(baseCmdForAppUser.Create)
	
	//get
	getAgentsCmdForAppUser := controller.InitGetAgentsCmdForAppUser(conf)
	baseCmdForAppUser.Get.AddCommand(getAgentsCmdForAppUser)
	getElementsCmdForAppUser := controller.InitGetElementsCmdForAppUser(conf)
	baseCmdForAppUser.Get.AddCommand(getElementsCmdForAppUser)
	getRelationsCmdForAppUser := controller.InitGetRelationsCmdForAppUser(conf)
	baseCmdForAppUser.Get.AddCommand(getRelationsCmdForAppUser)
	getLayersCmdForAppUser := controller.InitGetLayersCmdForAppUser(conf)
	baseCmdForAppUser.Get.AddCommand(getLayersCmdForAppUser)
	getSpacesCmdForAppUser := controller.InitGetSpacesCmdForAppUser(conf)
	baseCmdForAppUser.Get.AddCommand(getSpacesCmdForAppUser)
	rootCmdForAppUser.AddCommand(baseCmdForAppUser.Get)
	
	//update
	updateAgentsCmdForAppUser := controller.InitUpdateAgentsCmdForAppUser(conf)
	baseCmdForAppUser.Update.AddCommand(updateAgentsCmdForAppUser)
	updateElementsCmdForAppUser := controller.InitUpdateElementsCmdForAppUser(conf)
	baseCmdForAppUser.Update.AddCommand(updateElementsCmdForAppUser)
	updateRelationsCmdForAppUser := controller.InitUpdateRelationsCmdForAppUser(conf)
	baseCmdForAppUser.Update.AddCommand(updateRelationsCmdForAppUser)
	updateLayersCmdForAppUser := controller.InitUpdateLayersCmdForAppUser(conf)
	baseCmdForAppUser.Update.AddCommand(updateLayersCmdForAppUser)
	updateSpacesCmdForAppUser := controller.InitUpdateSpacesCmdForAppUser(conf)
	baseCmdForAppUser.Update.AddCommand(updateSpacesCmdForAppUser)
	rootCmdForAppUser.AddCommand(baseCmdForAppUser.Update)
	
	//delete
	deleteAgentsCmdForAppUser := controller.InitDeleteAgentsCmdForAppUser(conf)
	baseCmdForAppUser.Delete.AddCommand(deleteAgentsCmdForAppUser)
	deleteElementsCmdForAppUser := controller.InitDeleteElementsCmdForAppUser(conf)
	baseCmdForAppUser.Delete.AddCommand(deleteElementsCmdForAppUser)
	deleteRelationsCmdForAppUser := controller.InitDeleteRelationsCmdForAppUser(conf)
	baseCmdForAppUser.Delete.AddCommand(deleteRelationsCmdForAppUser)
	deleteLayersCmdForAppUser := controller.InitDeleteLayersCmdForAppUser(conf)
	baseCmdForAppUser.Delete.AddCommand(deleteLayersCmdForAppUser)
	deleteSpacesCmdForAppUser := controller.InitDeleteSpacesCmdForAppUser(conf)
	baseCmdForAppUser.Delete.AddCommand(deleteSpacesCmdForAppUser)
	rootCmdForAppUser.AddCommand(baseCmdForAppUser.Delete)

	rootCmdForAppUser.Execute()
}