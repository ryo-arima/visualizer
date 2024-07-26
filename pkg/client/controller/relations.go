package controller

import (
	"fmt"
	"log"

	"github.com/ryo-arima/visualizer/pkg/config"
	"github.com/spf13/cobra"
)

func InitBootstrapRelationsCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	bootstrapRelationsCmd := &cobra.Command{
		Use:   "relations",
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
	bootstrapRelationsCmd.Flags().StringP("key", "k", "", "cache key")
	return bootstrapRelationsCmd
}

func InitCreateRelationsCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	createRelationsCmd := &cobra.Command{
		Use:   "relations",
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
	createRelationsCmd.Flags().StringP("key", "k", "", "cache key")
	return createRelationsCmd
}

func InitCreateRelationsCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	createRelationsCmd := &cobra.Command{
		Use:   "relations",
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
	createRelationsCmd.Flags().StringP("key", "k", "", "cache key")
	return createRelationsCmd
}

func InitCreateRelationsCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	createRelationsCmd := &cobra.Command{
		Use:   "relations",
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
	createRelationsCmd.Flags().StringP("key", "k", "", "cache key")
	return createRelationsCmd
}

func InitGetRelationsCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	getRelationsCmd := &cobra.Command{
		Use:   "relations",
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
	getRelationsCmd.Flags().StringP("key", "k", "", "cache key")
	return getRelationsCmd
}

func InitGetRelationsCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	getRelationsCmd := &cobra.Command{
		Use:   "relations",
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
	getRelationsCmd.Flags().StringP("key", "k", "", "cache key")
	return getRelationsCmd
}

func InitGetRelationsCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	getRelationsCmd := &cobra.Command{
		Use:   "relations",
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
	getRelationsCmd.Flags().StringP("key", "k", "", "cache key")
	return getRelationsCmd
}

func InitUpdateRelationsCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	updateRelationsCmd := &cobra.Command{
		Use:   "relations",
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
	updateRelationsCmd.Flags().StringP("key", "k", "", "cache key")
	return updateRelationsCmd
}

func InitUpdateRelationsCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	updateRelationsCmd := &cobra.Command{
		Use:   "relations",
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
	updateRelationsCmd.Flags().StringP("key", "k", "", "cache key")
	return updateRelationsCmd
}

func InitUpdateRelationsCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	updateRelationsCmd := &cobra.Command{
		Use:   "relations",
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
	updateRelationsCmd.Flags().StringP("key", "k", "", "cache key")
	return updateRelationsCmd
}

func InitDeleteRelationsCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	deleteRelationsCmd := &cobra.Command{
		Use:   "relations",
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
	deleteRelationsCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteRelationsCmd
}

func InitDeleteRelationsCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	deleteRelationsCmd := &cobra.Command{
		Use:   "relations",
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
	deleteRelationsCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteRelationsCmd
}

func InitDeleteRelationsCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	deleteRelationsCmd := &cobra.Command{
		Use:   "relations",
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
	deleteRelationsCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteRelationsCmd
}