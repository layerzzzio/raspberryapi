package sys

import (
	"github.com/raspibuddy/rpi"
)

// Display represents a Display entity on the current system.
type Display struct{}

// List returns a list of Display info
func (d Display) List(readLines []string) (rpi.Display, error) {
	isOverscan := false

	for _, v := range readLines {
		if v == "disable_overscan=0" {
			isOverscan = true
		}
	}

	return rpi.Display{
		IsOverscan: isOverscan,
	}, nil
}
