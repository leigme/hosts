package cmd

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	version = "v1.0.1"
	p       = &param{}
	rootCmd = &cobra.Command{
		Use:   "hosts",
		Short: "Hosts file update tool",
		Long: `Hosts file update tool. For example:
				./hosts config -s k=v
				./hosts download
				./hosts update [-f]
				./hosts version
			.`,
	}
)

type param struct {
	Skip  bool
	Force bool
	Add   string
	Url   string
	Name  string
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.AddConfigPath(WorkDir())
	_, err := os.Stat(filepath.Join(WorkDir(), configName))
	if err != nil {
		if !os.IsNotExist(err) {
			log.Fatalln(err)
		}
		viper.SetConfigType("yaml")
		viper.SetConfigName(configNamePer())
		viper.SetDefault(hostsPath, defaultHostsPath)
		viper.SetDefault(hostsUrl, defaultHostsUrl)
		viper.SetDefault(tagStart, defaultTagStart)
		viper.SetDefault(tagEnd, defaultTagEnd)
		if err = viper.SafeWriteConfig(); err != nil {
			log.Fatalln(err)
		}
	} else {
		viper.SetConfigFile(filepath.Join(WorkDir(), configName))
	}
	if err = viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}
}

func WorkDir() (workDir string) {
	// 获取可执行文件的名称
	exe, err := os.Executable()
	if err != nil {
		log.Fatalln(err)
	}
	appName := filepath.Base(exe)

	if workDir, err = os.Getwd(); err != nil {
		log.Fatalln(err)
	}

	if _, err = os.Stat(filepath.Join(workDir, "."+appName)); err != nil {
		if !os.IsNotExist(err) {
			log.Fatalln(err)
		}
		workDir, err = os.UserHomeDir()
		if err != nil {
			log.Fatalln(err)
		}
	}

	workDir = filepath.Join(workDir, "."+appName)

	if _, err = os.Stat(workDir); err != nil {
		if !os.IsNotExist(err) {
			log.Fatalln(err)
		}
		if err = os.MkdirAll(workDir, os.ModePerm); err != nil {
			log.Fatalln(err)
		}
	}
	return
}
