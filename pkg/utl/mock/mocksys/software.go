package mocksys

import (
	"github.com/raspibuddy/rpi"
)

// Software mock
type Software struct {
	ListFn func(
		bool,
		bool,
		bool,
		bool,
		bool,
		bool,
		bool,
	) (rpi.Software, error)
}

// List mock
func (in Software) List(
	isVNC bool,
	isOpenVPN bool,
	isUnzip bool,
	isNordVPN bool,
	isSurfSharkVPN bool,
	isIpVanishVPN bool,
	isVyprVpnVPN bool,
) (rpi.Software, error) {
	return in.ListFn(
		isVNC,
		isOpenVPN,
		isUnzip,
		isNordVPN,
		isSurfSharkVPN,
		isIpVanishVPN,
		isVyprVpnVPN,
	)
}
