package mocksys

import (
	"github.com/raspibuddy/rpi"
)

// Port mock
type Port struct {
	ViewFn func(int32, bool) (rpi.Port, error)
}

// View mock
func (p Port) View(
	port int32,
	isListen bool,
) (rpi.Port, error) {
	return p.ViewFn(
		port,
		isListen,
	)
}
