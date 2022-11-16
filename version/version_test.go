package version_test

import (
	"fmt"
	"testing"

	"calf-restful-api/version"
)

func TestVersion(t *testing.T) {
	fmt.Println(version.FullVersion())
}
