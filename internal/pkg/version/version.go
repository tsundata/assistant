package version

import "fmt"

var (
	Version   = "unknown"
	GitCommit = "unknown"
	BuildTime = "unknown"
	GoVersion = "unknown"
)

func Info() string {
	return fmt.Sprintf("Version: %s\nGit Commit: %s\nBuilt: %s\nGo version: %s\n", Version, GitCommit, BuildTime, GoVersion)
}