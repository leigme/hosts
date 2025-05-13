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
	tagStart   = "tag_start"
	tagEnd     = "tag_end"

	configName = "config.yaml"
	hostsTmp   = "hosts.tmp"

	defaultHostsPath = "/etc/hosts"
	defaultHostsUrl  = "https://github.com/ittuann/GitHub-IP-hosts/raw/refs/heads/main/hosts"
	defaultTagStart  = "# GitHub IP hosts Start"
	defaultTagEnd    = "# GitHub IP hosts End"
)

func configNamePer() string {
	suf := filepath.Ext(configName)
	return strings.TrimSuffix(configName, suf)
}
