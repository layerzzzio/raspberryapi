package sys_test

import (
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/infos/appstatus"
	"github.com/raspibuddy/rpi/pkg/api/infos/appstatus/platform/sys"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	cases := []struct {
		name                 string
		statusVPNWithOpenVPN map[string]bool
		wantedData           rpi.AppStatus
		wantedErr            error
	}{
		{
			name: "success: isNordVPN true",
			statusVPNWithOpenVPN: map[string]bool{
				"nordvpn": true,
			},
			wantedData: rpi.AppStatus{
				VPNwithOpenVPN: map[string]bool{
					"nordvpn": true,
				},
			},
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := appstatus.APSFSYS(sys.AppStatus{})
			intf, err := s.List(
				tc.statusVPNWithOpenVPN,
			)
			assert.Equal(t, tc.wantedData, intf)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
