package sys_test

import (
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/infos/software"
	"github.com/raspibuddy/rpi/pkg/api/infos/software/platform/sys"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	cases := []struct {
		name           string
		isVNC          bool
		isOpenVPN      bool
		isUnzip        bool
		isNordVPN      bool
		isSurfSharkVPN bool
		isIpVanishVPN  bool
		isVyprVpnVPN   bool
		wantedData     rpi.Software
		wantedErr      error
	}{
		{
			name:           "success: all VPNs true",
			isVNC:          true,
			isOpenVPN:      true,
			isUnzip:        true,
			isNordVPN:      true,
			isSurfSharkVPN: true,
			isIpVanishVPN:  true,
			isVyprVpnVPN:   true,
			wantedData: rpi.Software{
				IsVNCInstalled:          true,
				IsOpenVPNInstalled:      true,
				IsUnzipInstalled:        true,
				IsNordVPNInstalled:      true,
				IsSurfSharkVPNInstalled: true,
				IsIpVanishVPNInstalled:  true,
				IsVyprVpnVPNInstalled:   true,
			},
			wantedErr: nil,
		},
		{
			name:           "success: all VPNs false",
			isVNC:          true,
			isOpenVPN:      true,
			isUnzip:        true,
			isNordVPN:      false,
			isSurfSharkVPN: false,
			isIpVanishVPN:  false,
			isVyprVpnVPN:   false,
			wantedData: rpi.Software{
				IsVNCInstalled:          true,
				IsOpenVPNInstalled:      true,
				IsUnzipInstalled:        true,
				IsNordVPNInstalled:      false,
				IsSurfSharkVPNInstalled: false,
				IsIpVanishVPNInstalled:  false,
				IsVyprVpnVPNInstalled:   false,
			},
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := software.SOFSYS(sys.Software{})
			intf, err := s.List(
				tc.isVNC,
				tc.isOpenVPN,
				tc.isUnzip,
				tc.isNordVPN,
				tc.isSurfSharkVPN,
				tc.isIpVanishVPN,
				tc.isVyprVpnVPN,
			)
			assert.Equal(t, tc.wantedData, intf)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
