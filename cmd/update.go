/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "A Update Command",
	Long: `A Update Command. For example:
			./hosts update [-f]
		`,
	Run: func(cmd *cobra.Command, args []string) {
		updateSource()
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	rootCmd.Flags().BoolVarP(&p.Skip, "skip", "S", false, "")
	rootCmd.Flags().BoolVarP(&p.Force, "force", "F", false, "")
}

func readSource(header, footer *[]string) {
	f, err := os.Open(viper.GetString(hostsPath))
	if err != nil {
		log.Fatalln(err)
	}
	defer func(f *os.File) {
		if err = f.Close(); err != nil {
			log.Println(err)
		}
	}(f)
	start := false
	end := false
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.EqualFold(line, "") {
			continue
		}
		if strings.EqualFold(viper.GetString(tagStart), line) {
			start = true
			continue
		}
		if !start {
			*header = append(*header, line)
			continue
		}
		if strings.EqualFold(viper.GetString(tagEnd), line) {
			end = true
			continue
		}
		if end {
			*footer = append(*footer, line)
		}
	}
}

func updateSource() {
	temp := filepath.Join(WorkDir(), hostsTmp)
	if !p.Skip {
		download(viper.GetString(hostsUrl), temp)
	}
	header := make([]string, 0)
	footer := make([]string, 0)
	body := make([]string, 0)
	f, err := os.Open(temp)
	if err != nil {
		log.Fatalln(err)
	}
	defer func(f *os.File) {
		if err = f.Close(); err != nil {
			log.Println(err)
		}
	}(f)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		body = append(body, scanner.Text())
	}
	if !p.Force {
		readSource(&header, &footer)
	}
	hs := viper.GetString(hostsPath)
	ht, err := os.OpenFile(filepath.Join(filepath.Dir(hs), hostsTmp), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}
	w := bufio.NewWriter(ht)
	if len(header) > 0 {
		writerHosts(w, header)
	}
	if len(body) > 0 {
		writerHosts(w, body)
	}
	if len(footer) > 0 {
		writerHosts(w, footer)
	}
	if err = w.Flush(); err != nil {
		log.Println(err)
	}
	if err = ht.Close(); err != nil {
		log.Println(err)
	}
	// 备份原文件（如果存在）
	if _, err = os.Stat(hs); err == nil {
		backupPath := hs + ".bak"
		if err = os.Rename(hs, backupPath); err != nil {
			log.Fatalln(err)
		}
		defer func(backupPath string) {
			// 成功替换后删除备份
			if err = os.Remove(backupPath); err != nil {
				log.Println(err)
			}
		}(backupPath)
	}
	// 替换原文件
	if err = os.Rename(ht.Name(), hs); err != nil {
		log.Fatalln(err)
	}
}

func writerHosts(w *bufio.Writer, content []string) {
	for _, l := range content {
		if _, err := w.WriteString(l + "\n"); err != nil {
			log.Println(err)
		}
	}
}
