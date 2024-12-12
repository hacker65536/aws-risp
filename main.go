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
		// Goreleaser doesn't prefix with a `v`, which we expect
		version = "v" + version
	}
	cmd.SetVersionInfo(version, commit, date)
	cmd.Execute()
}
