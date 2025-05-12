/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	workDir string
	p       = &param{}
	rootCmd = &cobra.Command{
		Use:   "hosts",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}
)

type param struct {
	ConfigInit         bool
	ConfigDetail       bool
	UpdateForce        bool
	DownloadMaxRetries int
	ConfigSet          string
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(wd string) {
	fmt.Println(wd)
	workDir = wd
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	if _, err := os.Stat(workDir); err != nil {
		if !os.IsNotExist(err) {
			log.Fatalln("check work dir fail!", err)
		}
		err = os.MkdirAll(workDir, os.ModePerm)
		if err != nil {
			log.Fatalln("create work dir fail!", err)
		}
	}
	viper.AddConfigPath(workDir)
	_, err := os.Stat(filepath.Join(workDir, configName))
	if err != nil {
		if !os.IsNotExist(err) {
			log.Fatalln(err)
		}
		viper.SetConfigType("yaml")
		viper.SetConfigName(configNamePer())
		viper.SetDefault(hostsPath, "/etc/hosts")
		viper.SetDefault(hostsUrl, "https://github.com/ittuann/GitHub-IP-hosts/raw/refs/heads/main/hosts")
		if err = viper.SafeWriteConfig(); err != nil {
			log.Fatalln(err)
		}
	} else {
		viper.SetConfigFile(filepath.Join(workDir, configName))
	}
	if err = viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}
}
