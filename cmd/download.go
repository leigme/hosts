/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
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
		destPath := filepath.Join(workDir, hostsTmp)
		if err := downloadExpand(url, destPath); err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	downloadCmd.Flags().IntVarP(&p.DownloadMaxRetries, "max_retries", "r", 3, "")
}

func downloadExpand(url, destPath string) error {
	// 创建目标目录
	if err := os.MkdirAll(filepath.Dir(destPath), os.ModePerm); err != nil {
		return err
	}
	for attempt := 0; attempt < p.DownloadMaxRetries; attempt++ {
		if attempt > 0 {
			waitTime := time.Duration(attempt*2) * time.Second
			fmt.Printf("尝试 %d/%d 失败，等待 %v 后重试...\n", attempt, p.DownloadMaxRetries, waitTime)
			time.Sleep(waitTime)
		}

		// 创建带有超时设置的 HTTP 客户端
		transport := &http.Transport{
			TLSHandshakeTimeout:   60 * time.Second,
			ResponseHeaderTimeout: 60 * time.Second,
		}
		client := &http.Client{
			Transport: transport,
			Timeout:   120 * time.Second,
		}
		if err := download(url, destPath, client); err != nil {
			log.Println(err)
		}
	}

	return fmt.Errorf("达到最大重试次数 (%d)，下载失败", p.DownloadMaxRetries)
}

func download(url, destPath string, client *http.Client) error {
	// 发起HTTP请求
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer func(resp *http.Response) {
		if err = resp.Body.Close(); err != nil {
			log.Println(err)
		}
	}(resp)

	// 检查HTTP状态码
	if resp.StatusCode != http.StatusOK {
		return &DownloadError{
			StatusCode: resp.StatusCode,
			Message:    "HTTP请求失败",
		}
	}

	// 获取文件大小（通过Content-Length头）
	fileSize := resp.ContentLength
	if fileSize <= 0 {
		return &DownloadError{
			StatusCode: resp.StatusCode,
			Message:    "无法获取文件大小",
		}
	}

	// 创建目标文件
	out, err := os.Create(destPath)
	if err != nil {
		return err
	}

	defer func(out *os.File) {
		if err = out.Close(); err != nil {
			log.Println(err)
		}
	}(out)

	// 初始化进度条
	bar := progressbar.NewOptions64(
		fileSize,
		progressbar.OptionSetDescription("下载中..."),
		progressbar.OptionSetWidth(50),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetRenderBlankState(true),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "=",
			SaucerHead:    ">",
			SaucerPadding: ".",
			BarStart:      "[",
			BarEnd:        "]",
		}),
	)

	// 使用 TeeReader 同时写入文件和更新进度条
	_, err = io.Copy(io.MultiWriter(out, bar), resp.Body)
	if err != nil {
		return err
	}
	return nil
}

// DownloadError 自定义下载错误类型
type DownloadError struct {
	StatusCode int
	Message    string
}

func (e *DownloadError) Error() string {
	return fmt.Sprintf("%s, 状态码: %d", e.Message, e.StatusCode)
}
