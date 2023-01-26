package sys

import (
	"regexp"

	"github.com/raspibuddy/rpi"
)

// Display represents a Display entity on the current system.
type Display struct{}

// List returns a list of Display info
func (d Display) List(
	readLines []string,
	isXscreenSaverInstalled bool,
	isBlanking bool,
) (rpi.Display, error) {
	isOverscan := false
	// I use a regex here to cover the below cases:
	// disable_overscan=0 (regular)
	// disable_overscan   = 0 (whitespace)
	// disable_overscan=0 #random bash comment (comment)
	re := regexp.MustCompile(`^\s*disable_overscan\s*=\s*0\s*\.*`)

	for _, v := range readLines {
		if re.MatchString(v) {
			isOverscan = true
		}
	}

	return rpi.Display{
		IsOverscan:              isOverscan,
		IsXscreenSaverInstalled: isXscreenSaverInstalled,
		IsBlanking:              isBlanking,
	}, nil
}
