package mocksys

import (
	"github.com/raspibuddy/rpi"
)

// Port mock
type Port struct {
	ViewFn func(bool) (rpi.Port, error)
}

// View mock
func (p Port) View(
	isListen bool,
) (rpi.Port, error) {
	return p.ViewFn(
		isListen,
	)
}
