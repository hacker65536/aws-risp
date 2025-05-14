package cmd

import (
	"fmt"
	"runtime"
	"runtime/debug"

	"github.com/spf13/cobra"
)

// バージョン情報を格納する変数
var (
	// コンパイル時に -ldflags で設定される変数
	version = ""
	commit  = ""
	date    = "unknown"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display version information",
	Long: `Show detailed version information of the aws-risp CLI tool.
This includes version number, build date, git commit hash,
and Go runtime information.`,
	Run: func(cmd *cobra.Command, args []string) {
		displayVersion()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

// displayVersion は詳細なバージョン情報を表示する
func displayVersion() {
	// バージョン情報が設定されていない場合は、ビルド情報から取得を試みる
	if version == "" && commit == "" {
		displayBuildInfo()
	} else {
		displayStaticInfo()
	}

	fmt.Printf("Go version: %s\n", runtime.Version())
	fmt.Printf("OS/Arch:    %s/%s\n", runtime.GOOS, runtime.GOARCH)
}

func displayBuildInfo() {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		fmt.Println("Build information not available.")
		return
	}

	fmt.Printf("Version:    %s\n", info.Main.Version)
	fmt.Printf("Commit:     %s\n", info.Main.Sum)
}

func displayStaticInfo() {
	fmt.Printf("Version:    %s\n", version)
	fmt.Printf("Commit:     %s\n", commit)
	fmt.Printf("Built:      %s\n", date)
}
