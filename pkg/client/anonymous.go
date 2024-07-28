package client

import (
	"github.com/ryo-arima/visualizer/pkg/client/controller"
	"github.com/ryo-arima/visualizer/pkg/config"
	"github.com/spf13/cobra"
)

type BaseCmdForAnonymousUser struct {
	Get       *cobra.Command
}

func InitRootCmdForAnonymousUser() *cobra.Command {
	var rootCmdForAnonymousUser = &cobra.Command{
		Use:   "visualizer-anonymous",
		Short: "'visualizer' is a CLI tool to manage anniversaries",
		Long:  `''`,
	}
	return rootCmdForAnonymousUser
}

func InitBaseCmdForAnonymousUser() BaseCmdForAnonymousUser {
	getCmd := &cobra.Command{
		Use:   "get",
		Short: "get the value of a key",
		Long:  "get the value of a key",
	}
	baseCmdForAnonymousUser := BaseCmdForAnonymousUser{
		Get:       getCmd,
	}
	return baseCmdForAnonymousUser
}

func ClientForAnonymousUser(conf config.BaseConfig) {
	rootCmdForAnonymousUser := InitRootCmdForAnonymousUser()
	rootCmdForAnonymousUser.CompletionOptions.HiddenDefaultCmd = true
	baseCmdForAnonymousUser := InitBaseCmdForAnonymousUser()

	//get
	getAgentsCmdForAnonymousUser := controller.InitGetAgentsCmdForAnonymousUser(conf)
	baseCmdForAnonymousUser.Get.AddCommand(getAgentsCmdForAnonymousUser)
	getElementsCmdForAnonymousUser := controller.InitGetElementsCmdForAnonymousUser(conf)
	baseCmdForAnonymousUser.Get.AddCommand(getElementsCmdForAnonymousUser)
	getRelationsCmdForAnonymousUser := controller.InitGetRelationsCmdForAnonymousUser(conf)
	baseCmdForAnonymousUser.Get.AddCommand(getRelationsCmdForAnonymousUser)
	getLayersCmdForAnonymousUser := controller.InitGetLayersCmdForAnonymousUser(conf)
	baseCmdForAnonymousUser.Get.AddCommand(getLayersCmdForAnonymousUser)
	getSpacesCmdForAnonymousUser := controller.InitGetSpacesCmdForAnonymousUser(conf)
	baseCmdForAnonymousUser.Get.AddCommand(getSpacesCmdForAnonymousUser)
	rootCmdForAnonymousUser.AddCommand(baseCmdForAnonymousUser.Get)
	
	rootCmdForAnonymousUser.Execute()
}