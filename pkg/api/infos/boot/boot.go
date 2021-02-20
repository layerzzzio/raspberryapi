package boot

import (
	"github.com/raspibuddy/rpi"
)

// List populates and returns an array of Boot model.
func (b *Boot) List() (rpi.Boot, error) {
	isWaitForNetwork := b.i.IsFileExists("/etc/systemd/system/dhcpcd.service.d/wait.conf")
	return b.boosys.List(isWaitForNetwork)
}
