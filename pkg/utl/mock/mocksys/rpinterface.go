package mocksys

import (
	"github.com/raspibuddy/rpi"
)

// RpInterface mock
type RpInterface struct {
	ListFn func([]string, bool, bool, bool) (rpi.RpInterface, error)
}

// List mock
func (in RpInterface) List(
	readLines []string,
	isStartXElf bool,
	isSSH bool,
	isSSHKeyGenerating bool,
) (rpi.RpInterface, error) {
	return in.ListFn(readLines, isStartXElf, isSSH, isSSHKeyGenerating)
}
