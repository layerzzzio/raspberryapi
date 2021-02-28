package boot

import (
	"github.com/raspibuddy/rpi"
)

// List populates and returns an array of Boot model.
func (b *Boot) List() (rpi.Boot, error) {
	isWaitForNetworkPath := b.i.GetConfigFiles()["waitfornetwork"].Path
	isWaitForNetwork := b.i.IsFileExists(isWaitForNetworkPath)
	return b.boosys.List(isWaitForNetwork)
}
