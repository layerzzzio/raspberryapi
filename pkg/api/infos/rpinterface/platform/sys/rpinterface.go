package sys

import (
	"regexp"

	"github.com/raspibuddy/rpi"
)

// RpInterface represents a RpInterface entity on the current system.
type RpInterface struct{}

// List returns a list of RpInterface info
func (int RpInterface) List(
	readLines []string,
	isStartXElf bool,
	isSSH bool,
	isSSHKeyGenerating bool,
) (rpi.RpInterface, error) {
	isCamera := false
	// I use a regex here to cover the below cases:
	// start_x=0 (regular)
	// start_x   = 0 (whitespace)
	// start_x=0 #random bash comment (comment)
	re := regexp.MustCompile(`^\s*start_x\s*=\s*1\s*\.*`)

	for _, v := range readLines {
		if re.MatchString(v) {
			isCamera = true
		}
	}

	return rpi.RpInterface{
		IsStartXElf:        isStartXElf,
		IsCamera:           isCamera,
		IsSSH:              isSSH,
		IsSSHKeyGenerating: isSSHKeyGenerating,
	}, nil
}
