package mocksys

import (
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/infos/software"
)

// Software mock
type Software struct {
	ListFn func(
		bool,
		bool,
		bool,
		software.NordVPN,
	) (rpi.Software, error)
}

// List mock
func (in Software) List(
	isVNC bool,
	isOpenVPN bool,
	isUnzip bool,
	nordVPN software.NordVPN,
) (rpi.Software, error) {
	return in.ListFn(
		isVNC,
		isOpenVPN,
		isUnzip,
		nordVPN,
	)
}
