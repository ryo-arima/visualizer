package controller

import (
	"fmt"
	"log"

	"github.com/ryo-arima/visualizer/pkg/config"
	"github.com/spf13/cobra"
)

func InitBootstrapSpacesCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	bootstrapSpacesCmd := &cobra.Command{
		Use:   "spaces",
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
	bootstrapSpacesCmd.Flags().StringP("key", "k", "", "cache key")
	return bootstrapSpacesCmd
}

func InitCreateSpacesCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	createSpacesCmd := &cobra.Command{
		Use:   "spaces",
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
	createSpacesCmd.Flags().StringP("key", "k", "", "cache key")
	return createSpacesCmd
}

func InitCreateSpacesCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	createSpacesCmd := &cobra.Command{
		Use:   "spaces",
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
	createSpacesCmd.Flags().StringP("key", "k", "", "cache key")
	return createSpacesCmd
}

func InitCreateSpacesCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	createSpacesCmd := &cobra.Command{
		Use:   "spaces",
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
	createSpacesCmd.Flags().StringP("key", "k", "", "cache key")
	return createSpacesCmd
}

func InitGetSpacesCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	getSpacesCmd := &cobra.Command{
		Use:   "spaces",
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
	getSpacesCmd.Flags().StringP("key", "k", "", "cache key")
	return getSpacesCmd
}

func InitGetSpacesCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	getSpacesCmd := &cobra.Command{
		Use:   "spaces",
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
	getSpacesCmd.Flags().StringP("key", "k", "", "cache key")
	return getSpacesCmd
}

func InitGetSpacesCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	getSpacesCmd := &cobra.Command{
		Use:   "spaces",
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
	getSpacesCmd.Flags().StringP("key", "k", "", "cache key")
	return getSpacesCmd
}

func InitUpdateSpacesCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	updateSpacesCmd := &cobra.Command{
		Use:   "spaces",
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
	updateSpacesCmd.Flags().StringP("key", "k", "", "cache key")
	return updateSpacesCmd
}

func InitUpdateSpacesCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	updateSpacesCmd := &cobra.Command{
		Use:   "spaces",
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
	updateSpacesCmd.Flags().StringP("key", "k", "", "cache key")
	return updateSpacesCmd
}

func InitUpdateSpacesCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	updateSpacesCmd := &cobra.Command{
		Use:   "spaces",
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
	updateSpacesCmd.Flags().StringP("key", "k", "", "cache key")
	return updateSpacesCmd
}

func InitDeleteSpacesCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	deleteSpacesCmd := &cobra.Command{
		Use:   "spaces",
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
	deleteSpacesCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteSpacesCmd
}

func InitDeleteSpacesCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	deleteSpacesCmd := &cobra.Command{
		Use:   "spaces",
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
	deleteSpacesCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteSpacesCmd
}

func InitDeleteSpacesCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	deleteSpacesCmd := &cobra.Command{
		Use:   "spaces",
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
	deleteSpacesCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteSpacesCmd
}