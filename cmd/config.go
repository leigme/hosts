/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if p.ConfigInit {
			if err := viper.WriteConfigAs(filepath.Join(workDir, configName)); err != nil {
				log.Fatalln(err)
			}
		}
		if p.ConfigDetail {
			f, err := os.Open(filepath.Join(workDir, configName))
			if err != nil {
				log.Fatalln("config file is nil, please command init")
			}
			defer func(f *os.File) {
				if err = f.Close(); err != nil {
					log.Println(err)
				}
			}(f)
			fmt.Println(configPath + ": " + filepath.Join(workDir, configName))
			fr := bufio.NewScanner(f)
			for fr.Scan() {
				fmt.Println(fr.Text())
			}
			if fr.Err() != nil {
				log.Fatalln(fr.Err())
			}
		}
		if !strings.EqualFold(p.ConfigSet, "") {
			params := strings.Split(p.ConfigSet, "=")
			if len(params) != 2 {
				log.Fatalln("set param must contain = ")
			}
			viper.Set(params[0], params[1])
			if err := viper.WriteConfigAs(filepath.Join(workDir, configName)); err != nil {
				log.Fatalln("writer config as "+viper.GetString(configPath)+" fail!", err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().BoolVarP(&p.ConfigDetail, "detail", "d", false, "")
	configCmd.Flags().BoolVarP(&p.ConfigInit, "init", "i", false, "")
	configCmd.Flags().StringVarP(&p.ConfigSet, "set", "s", "", "")
}
