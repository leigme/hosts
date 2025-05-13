/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		url := viper.GetString(hostsUrl)
		destPath := filepath.Join(WorkDir(), hostsTmp)
		download(url, destPath)
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
}

func download(url, filename string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer func(response *http.Response) {
		if err = response.Body.Close(); err != nil {
			log.Fatalln(err)
		}
	}(resp)
	bar := progressbar.DefaultBytes(
		resp.ContentLength, "downloading: ")
	out, err := os.Create(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer func(o *os.File) {
		if err = o.Close(); err != nil {
			log.Fatalln(err)
		}
	}(out)
	_, err = io.Copy(io.MultiWriter(out, bar), resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
}
