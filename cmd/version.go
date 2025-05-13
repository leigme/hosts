/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "A Program Version command",
	Long: `A Program Version. For example:
			./hosts version
		`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("version: %s\n", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
