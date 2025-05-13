package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

// バージョン情報を格納する変数
var (
	// コンパイル時に -ldflags で設定される変数
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version information of aws-risp",
	Long: `Show detailed version information of the aws-risp CLI tool.
This includes version number, build date, git commit hash,
and Go runtime information.`,
	Run: func(cmd *cobra.Command, args []string) {
		displayVersion()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// displayVersion は詳細なバージョン情報を表示する
func displayVersion() {
	fmt.Printf("aws-risp: AWS Reservation Information Service Provider\n")
	fmt.Printf("Version:    %s\n", version)
	fmt.Printf("Commit:     %s\n", commit)
	fmt.Printf("Built:      %s\n", date)
	fmt.Printf("Go version: %s\n", runtime.Version())
	fmt.Printf("OS/Arch:    %s/%s\n", runtime.GOOS, runtime.GOARCH)
}
