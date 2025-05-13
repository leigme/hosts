/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/leigme/hosts/cmd"
	"log"
)

func main() {
	cmd.Execute()
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
