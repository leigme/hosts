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
	"strings"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "A Download Command",
	Long: `A Download Command. For example:
			./hosts download
		`,
	Run: func(cmd *cobra.Command, args []string) {
		url := p.Url
		if strings.EqualFold(url, "") {
			url = viper.GetString(hostsUrl)
		}
		name := p.Name
		if strings.EqualFold(name, "") {
			name = hostsTmp
		}
		filename := filepath.Join(WorkDir(), hostsTmp)
		download(url, filename)
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
	rootCmd.Flags().StringVarP(&p.Url, "url", "U", "", "")
	rootCmd.Flags().StringVarP(&p.Name, "name", "N", "", "")
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
