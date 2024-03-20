package cmd

import (
	_ "embed"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"log/slog"
)

const (
	author = "Mystery0"
)

var (
	//go:embed version.txt
	versionTxt     string
	GitCommitHex   string
	GitCommitCount string
	BuildTime      string
	GoVersion      string
)

var version = fmt.Sprintf(`%s.%s-%s`, versionTxt, GitCommitCount, GitCommitHex)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version information",
	Run: func(cmd *cobra.Command, args []string) {
		x := color.RedString("♥")
		a := color.CyanString(author)
		slog.Info("CLI Version:", color.RedString("[%s]", version))
		slog.Info("Go Version:", color.GreenString("[%s]", GoVersion))
		slog.Info("Build Time:", color.BlueString("[%s]", BuildTime))
		slog.Info("  —— Made with", x, "by", a)
	},
}
