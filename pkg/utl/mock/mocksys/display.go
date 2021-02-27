package mocksys

import (
	"github.com/raspibuddy/rpi"
)

// Display mock
type Display struct {
	ListFn func(readlines []string) (rpi.Display, error)
}

// List mock
func (d Display) List(readlines []string) (rpi.Display, error) {
	return d.ListFn(readlines)
}
