package mocksys

import (
	"github.com/raspibuddy/rpi"
)

// Boot mock
type Boot struct {
	ListFn func(isWaitForNetwork bool) (rpi.Boot, error)
}

// List mock
func (b Boot) List(isWaitForNetwork bool) (rpi.Boot, error) {
	return b.ListFn(isWaitForNetwork)
}
