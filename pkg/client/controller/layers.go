package controller

import (
	"fmt"
	"log"

	"github.com/ryo-arima/visualizer/pkg/config"
	"github.com/spf13/cobra"
)

func InitBootstrapLayersCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	bootstrapLayersCmd := &cobra.Command{
		Use:   "layers",
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
	bootstrapLayersCmd.Flags().StringP("key", "k", "", "cache key")
	return bootstrapLayersCmd
}

func InitCreateLayersCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	createLayersCmd := &cobra.Command{
		Use:   "layers",
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
	createLayersCmd.Flags().StringP("key", "k", "", "cache key")
	return createLayersCmd
}

func InitCreateLayersCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	createLayersCmd := &cobra.Command{
		Use:   "layers",
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
	createLayersCmd.Flags().StringP("key", "k", "", "cache key")
	return createLayersCmd
}

func InitCreateLayersCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	createLayersCmd := &cobra.Command{
		Use:   "layers",
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
	createLayersCmd.Flags().StringP("key", "k", "", "cache key")
	return createLayersCmd
}

func InitGetLayersCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	getLayersCmd := &cobra.Command{
		Use:   "layers",
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
	getLayersCmd.Flags().StringP("key", "k", "", "cache key")
	return getLayersCmd
}

func InitGetLayersCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	getLayersCmd := &cobra.Command{
		Use:   "layers",
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
	getLayersCmd.Flags().StringP("key", "k", "", "cache key")
	return getLayersCmd
}

func InitGetLayersCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	getLayersCmd := &cobra.Command{
		Use:   "layers",
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
	getLayersCmd.Flags().StringP("key", "k", "", "cache key")
	return getLayersCmd
}

func InitUpdateLayersCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	updateLayersCmd := &cobra.Command{
		Use:   "layers",
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
	updateLayersCmd.Flags().StringP("key", "k", "", "cache key")
	return updateLayersCmd
}

func InitUpdateLayersCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	updateLayersCmd := &cobra.Command{
		Use:   "layers",
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
	updateLayersCmd.Flags().StringP("key", "k", "", "cache key")
	return updateLayersCmd
}

func InitUpdateLayersCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	updateLayersCmd := &cobra.Command{
		Use:   "layers",
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
	updateLayersCmd.Flags().StringP("key", "k", "", "cache key")
	return updateLayersCmd
}

func InitDeleteLayersCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	deleteLayersCmd := &cobra.Command{
		Use:   "layers",
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
	deleteLayersCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteLayersCmd
}

func InitDeleteLayersCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	deleteLayersCmd := &cobra.Command{
		Use:   "layers",
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
	deleteLayersCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteLayersCmd
}

func InitDeleteLayersCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	deleteLayersCmd := &cobra.Command{
		Use:   "layers",
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
	deleteLayersCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteLayersCmd
}