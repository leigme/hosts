/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"path/filepath"
	"strings"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Set configuration command",
	Long: `Set configuration command For example:
			./hosts config -s k=v
		`,
	Run: func(cmd *cobra.Command, args []string) {
		if !strings.EqualFold(p.ConfigSet, "") {
			params := strings.Split(p.ConfigSet, "=")
			if len(params) != 2 {
				log.Fatalln("set param must contain = ")
			}
			viper.Set(params[0], params[1])
			if err := viper.WriteConfigAs(filepath.Join(WorkDir(), configName)); err != nil {
				log.Fatalln("writer config as "+viper.GetString(configPath)+" fail!", err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().StringVarP(&p.ConfigSet, "set", "s", "", "")
}
