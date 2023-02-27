package version

import "fmt"

const (
	Major = 0
	Minor = 0
	Patch = 6
)

func SemVer() string {
	return fmt.Sprintf("%d.%d.%d", Major, Minor, Patch)
}
