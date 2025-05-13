/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"time"

	"github.com/carlmjohnson/versioninfo"
	"github.com/hacker65536/aws-risp/cmd"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	if version == "dev" {
		version = versioninfo.Version
		commit = versioninfo.Revision
		date = versioninfo.LastCommit.Format(time.RFC3339)
	} else {
		// バージョン文字列に'v'のプレフィックスが既にある場合は追加しない
		if len(version) == 0 || version[0] != 'v' {
			version = "v" + version
		}
	}

	// rootコマンドのバージョン表示用
	cmd.SetVersionInfo(version, commit, date)
	cmd.Execute()
}
