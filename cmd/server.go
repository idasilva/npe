package cmd

import (
	"fmt"
	"github.com/idasilva/npe/app"
	"github.com/idasilva/npe/internal/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Create a new sample server",
	Long:  "Long description",
	RunE: func(cmd *cobra.Command, args []string) error {
		var config pkg.CfgFile
		err := viper.Unmarshal(&config)
		if err != nil {
			fmt.Sprintf("unable to decode into struct, %v", err)
		}

		fmt.Println(config.Env)

		app, err := app.Run(config)
		if err != nil {
			return err
		}

		err = app.Server()
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	rootCmd.PersistentFlags().StringP("env", "e", "prod", "set an environment")

	_ = viper.BindPFlag("env", rootCmd.PersistentFlags().Lookup("env"))
}
