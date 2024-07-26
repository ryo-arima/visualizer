package controller

import (
	"fmt"
	"log"

	"github.com/ryo-arima/visualizer/pkg/config"
	"github.com/spf13/cobra"
)

func InitBootstrapElementsCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	bootstrapElementsCmd := &cobra.Command{
		Use:   "elements",
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
	bootstrapElementsCmd.Flags().StringP("key", "k", "", "cache key")
	return bootstrapElementsCmd
}

func InitCreateElementsCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	createElementsCmd := &cobra.Command{
		Use:   "elements",
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
	createElementsCmd.Flags().StringP("key", "k", "", "cache key")
	return createElementsCmd
}

func InitCreateElementsCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	createElementsCmd := &cobra.Command{
		Use:   "elements",
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
	createElementsCmd.Flags().StringP("key", "k", "", "cache key")
	return createElementsCmd
}

func InitCreateElementsCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	createElementsCmd := &cobra.Command{
		Use:   "elements",
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
	createElementsCmd.Flags().StringP("key", "k", "", "cache key")
	return createElementsCmd
}

func InitGetElementsCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	getElementsCmd := &cobra.Command{
		Use:   "elements",
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
	getElementsCmd.Flags().StringP("key", "k", "", "cache key")
	return getElementsCmd
}

func InitGetElementsCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	getElementsCmd := &cobra.Command{
		Use:   "elements",
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
	getElementsCmd.Flags().StringP("key", "k", "", "cache key")
	return getElementsCmd
}

func InitGetElementsCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	getElementsCmd := &cobra.Command{
		Use:   "elements",
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
	getElementsCmd.Flags().StringP("key", "k", "", "cache key")
	return getElementsCmd
}

func InitUpdateElementsCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	updateElementsCmd := &cobra.Command{
		Use:   "elements",
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
	updateElementsCmd.Flags().StringP("key", "k", "", "cache key")
	return updateElementsCmd
}

func InitUpdateElementsCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	updateElementsCmd := &cobra.Command{
		Use:   "elements",
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
	updateElementsCmd.Flags().StringP("key", "k", "", "cache key")
	return updateElementsCmd
}

func InitUpdateElementsCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	updateElementsCmd := &cobra.Command{
		Use:   "elements",
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
	updateElementsCmd.Flags().StringP("key", "k", "", "cache key")
	return updateElementsCmd
}

func InitDeleteElementsCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	deleteElementsCmd := &cobra.Command{
		Use:   "elements",
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
	deleteElementsCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteElementsCmd
}

func InitDeleteElementsCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	deleteElementsCmd := &cobra.Command{
		Use:   "elements",
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
	deleteElementsCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteElementsCmd
}

func InitDeleteElementsCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	deleteElementsCmd := &cobra.Command{
		Use:   "elements",
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
	deleteElementsCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteElementsCmd
}