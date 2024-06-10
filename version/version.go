package version

import "fmt"

const (
	Major = 1
	Minor = 1
	Patch = 0
	Meta  = ""
)

func SemVer() string {
	ver := fmt.Sprintf("%d.%d.%d", Major, Minor, Patch)
	if Meta != "" {
		ver = ver + "-" + Meta
	}
	return ver
}
