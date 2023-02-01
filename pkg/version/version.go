//go:generate go run gen_helper.go
//go:generate go fmt

package version

import "fmt"

var (
	ShortVersion   = "dev"
	GitSha1Version = "git-sha1"
	BuildDate      = "2016-01-02"
)

func PrintVersionInfo(printer func(string, ...interface{})) {
	printer("Release OpVersion: %s", ShortVersion)
	printer("Git Commit Hash: %s", GitSha1Version)
	printer("Build Time: %s", BuildDate)
}

func GetVersionString() string {
	return fmt.Sprintf("%s; git: %s; build time: %s", ShortVersion, GitSha1Version, BuildDate)
}
