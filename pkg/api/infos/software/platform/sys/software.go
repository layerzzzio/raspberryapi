package sys

import (
	"github.com/raspibuddy/rpi"
)

// Software represents a Software entity on the current system.
type Software struct{}

// List returns a list of Software info
func (int Software) List(
	isVNC bool,
	isOpenVPN bool,
) (rpi.Software, error) {
	return rpi.Software{
		IsVNC:     isVNC,
		IsOpenVPN: isOpenVPN,
	}, nil
}
