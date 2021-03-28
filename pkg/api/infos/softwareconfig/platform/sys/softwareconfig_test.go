package sys_test

import (
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/infos/softwareconfig"
	"github.com/raspibuddy/rpi/pkg/api/infos/softwareconfig/platform/sys"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	cases := []struct {
		name         string
		VPNCountries map[string][]string
		wantedData   rpi.SoftwareConfig
		wantedErr    error
	}{
		{
			name: "success: isNordVPN true",
			VPNCountries: map[string][]string{
				"nordvpn": {"France", "Germany"},
			},
			wantedData: rpi.SoftwareConfig{
				VPNCountries: map[string][]string{
					"nordvpn": {"France", "Germany"},
				},
			},
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := softwareconfig.SOCFSYS(sys.SoftwareConfig{})
			intf, err := s.List(
				tc.VPNCountries,
			)
			assert.Equal(t, tc.wantedData, intf)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
