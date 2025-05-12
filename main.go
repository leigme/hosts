/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
	"syscall"
)

var workDir string

func main() {
	if strings.EqualFold(workDir, "") {
		// 获取当前用户信息
		currentUser, err := user.Current()
		if err != nil {
			log.Fatalf("无法获取当前用户信息: %v", err)
		}

		// 获取当前用户的主目录
		homeDir := os.Getenv("HOME")

		if homeDir == "" {
			homeDir = currentUser.HomeDir
		}

		fmt.Println("==> " + homeDir)

		ex, err := os.Executable()
		if err != nil {
			log.Fatalln(err)
		}

		workDir = filepath.Join(homeDir, "."+filepath.Base(ex))

		fmt.Println("---> " + workDir)
	}
	// 检查是否有 root 权限
	if syscall.Geteuid() != 0 {
		fmt.Println("需要 root 权限，尝试使用 sudo 运行...")

		// 使用 sudo 重新执行当前程序
		command := exec.Command("sudo", os.Args...)
		command.Stdin = os.Stdin
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr

		err := command.Run()
		if err != nil {
			log.Fatalf("sudo 执行失败: %v", err)
		}
		return
	}
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
