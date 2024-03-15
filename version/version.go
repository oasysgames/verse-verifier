package version

import "fmt"

const (
	Major = 0
	Minor = 0
	Patch = 11
	Meta  = "dev1"
)

func SemVer() string {
	ver := fmt.Sprintf("%d.%d.%d", Major, Minor, Patch)
	if Meta != "" {
		ver = ver + "-" + Meta
	}
	return ver
}
