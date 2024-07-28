package controller

import (
	"fmt"
	"log"

	"github.com/ryo-arima/visualizer/pkg/config"
	"github.com/spf13/cobra"
)

func InitBootstrapAgentsCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	bootstrapAgentsCmd := &cobra.Command{
		Use:   "agents",
		Short: "bootstrap the value of a key",
		Long:  "bootstrap the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	bootstrapAgentsCmd.Flags().StringP("key", "k", "", "cache key")
	return bootstrapAgentsCmd
}

func InitCreateAgentsCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	createAgentsCmd := &cobra.Command{
		Use:   "agents",
		Short: "create the value of a key",
		Long:  "create the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	createAgentsCmd.Flags().StringP("key", "k", "", "cache key")
	return createAgentsCmd
}

func InitCreateAgentsCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	createAgentsCmd := &cobra.Command{
		Use:   "agents",
		Short: "create the value of a key",
		Long:  "create the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	createAgentsCmd.Flags().StringP("key", "k", "", "cache key")
	return createAgentsCmd
}

func InitCreateAgentsCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	createAgentsCmd := &cobra.Command{
		Use:   "agents",
		Short: "create the value of a key",
		Long:  "create the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	createAgentsCmd.Flags().StringP("key", "k", "", "cache key")
	return createAgentsCmd
}

func InitGetAgentsCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	getAgentsCmd := &cobra.Command{
		Use:   "agents",
		Short: "get the value of a key",
		Long:  "get the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	getAgentsCmd.Flags().StringP("key", "k", "", "cache key")
	return getAgentsCmd
}

func InitGetAgentsCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	getAgentsCmd := &cobra.Command{
		Use:   "agents",
		Short: "get the value of a key",
		Long:  "get the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	getAgentsCmd.Flags().StringP("key", "k", "", "cache key")
	return getAgentsCmd
}

func InitGetAgentsCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	getAgentsCmd := &cobra.Command{
		Use:   "agents",
		Short: "get the value of a key",
		Long:  "get the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	getAgentsCmd.Flags().StringP("key", "k", "", "cache key")
	return getAgentsCmd
}

func InitUpdateAgentsCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	updateAgentsCmd := &cobra.Command{
		Use:   "agents",
		Short: "update the value of a key",
		Long:  "udpate the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	updateAgentsCmd.Flags().StringP("key", "k", "", "cache key")
	return updateAgentsCmd
}

func InitUpdateAgentsCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	updateAgentsCmd := &cobra.Command{
		Use:   "agents",
		Short: "update the value of a key",
		Long:  "udpate the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	updateAgentsCmd.Flags().StringP("key", "k", "", "cache key")
	return updateAgentsCmd
}

func InitUpdateAgentsCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	updateAgentsCmd := &cobra.Command{
		Use:   "agents",
		Short: "update the value of a key",
		Long:  "update the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	updateAgentsCmd.Flags().StringP("key", "k", "", "cache key")
	return updateAgentsCmd
}

func InitDeleteAgentsCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	deleteAgentsCmd := &cobra.Command{
		Use:   "agents",
		Short: "delete the value of a key",
		Long:  "delete the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	deleteAgentsCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteAgentsCmd
}

func InitDeleteAgentsCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	deleteAgentsCmd := &cobra.Command{
		Use:   "agents",
		Short: "delete the value of a key",
		Long:  "delete the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	deleteAgentsCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteAgentsCmd
}

func InitDeleteAgentsCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	deleteAgentsCmd := &cobra.Command{
		Use:   "agents",
		Short: "delete the value of a key",
		Long:  "delete the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	deleteAgentsCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteAgentsCmd
}