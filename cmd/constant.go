package cmd

import (
	"path/filepath"
	"strings"
)

/*
Copyright © 2025 leig <leigme@gmail.com>
*/

const (
	configPath = "config_path"
	hostsPath  = "hosts_path"
	hostsUrl   = "hosts_url"

	configName = "config.yaml"
	hostsTpl   = "hosts.tpl"
	hostsTmp   = "hosts.tmp"
	tagStart   = "# GitHub IP hosts Start"
	tagEnd     = "# GitHub IP hosts End"
)

func configNamePer() string {
	suf := filepath.Ext(configName)
	return strings.TrimSuffix(configName, suf)
}
