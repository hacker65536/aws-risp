package cmd

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"strings"

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

	// バージョンコマンドを登録し、バージョン情報を設定
	if version != "" {
		// 先頭の "v" が重複しないように調整
		if strings.HasPrefix(version, "v") {
			rootCmd.Version = version
		} else {
			rootCmd.Version = "v" + version
		}
	}
}

// displayVersion は詳細なバージョン情報を表示する
func displayVersion() {
	// バージョン情報の取得
	versionInfo := getVersionInfo()

	// バージョン情報の表示
	fmt.Println("AWS Reserved Instance and Savings Plan CLI")
	fmt.Println("------------------------------------------")
	fmt.Printf("Version:    %s\n", versionInfo.version)
	fmt.Printf("Commit:     %s\n", versionInfo.commit)
	fmt.Printf("Built:      %s\n", versionInfo.date)
	fmt.Printf("Go version: %s\n", runtime.Version())
	fmt.Printf("OS/Arch:    %s/%s\n", runtime.GOOS, runtime.GOARCH)
}

// バージョン情報を保持する構造体
type versionInformation struct {
	version string
	commit  string
	date    string
}

// getVersionInfo はビルド情報とコンパイル時情報を組み合わせてバージョン情報を返す
func getVersionInfo() versionInformation {
	// 基本情報を設定
	info := versionInformation{
		version: version,
		commit:  commit,
		date:    date,
	}

	// コンパイル時情報が設定されていない場合はビルド情報から取得
	needBuildInfo := info.version == "" || info.commit == ""

	if needBuildInfo {
		buildInfo, ok := debug.ReadBuildInfo()
		if ok {
			// バージョンが設定されていない場合
			if info.version == "" && buildInfo.Main.Version != "" {
				info.version = buildInfo.Main.Version
			}

			// コミットハッシュが設定されていない場合
			if info.commit == "" {
				for _, setting := range buildInfo.Settings {
					if setting.Key == "vcs.revision" {
						info.commit = setting.Value
						break
					}
				}
			}
		}
	}

	// 空文字列の場合はunknownを設定
	if info.version == "" {
		info.version = "unknown"
	}
	if info.commit == "" {
		info.commit = "unknown"
	}

	return info
}
