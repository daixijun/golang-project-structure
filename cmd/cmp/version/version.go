package version

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	version   = "master"
	gitCommit = "unknown"
	buildDate = "unknown"
)

var (
	fullVersion bool
	versionText = `Cloud Manager Platform
Version:   %s
GoVersion: %s
GitCommit: %s
BuildDate: %s
Compiler:  %s
Platform:  %s/%s
`
)

// NewCommand 创建命令
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Show build and version",
		Run: func(_ *cobra.Command, _ []string) {
			if fullVersion {
				fmt.Printf(
					versionText,
					version,
					runtime.Version(),
					gitCommit,
					buildDate,
					runtime.Compiler,
					runtime.GOOS,
					runtime.GOARCH,
				)
			} else {
				fmt.Printf("Version: %s", version)

			}
		},
	}
	cmd.Flags().BoolVarP(&fullVersion, "full", "f", false, "Show full version")

	return cmd
}
