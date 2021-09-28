package sys_test

import (
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/infos/appconfig"
	"github.com/raspibuddy/rpi/pkg/api/infos/appconfig/platform/sys"

	"github.com/stretchr/testify/assert"
)

func TestListVPN(t *testing.T) {
	cases := []struct {
		name         string
		VPNCountries map[string]map[string]string
		wantedData   rpi.AppConfigVPNWithOvpn
		wantedErr    error
	}{
		{
			name: "success: isNordVPN true",
			VPNCountries: map[string]map[string]string{
				"nordvpn": {"France": "file1", "Germany": "file2"},
			},
			wantedData: rpi.AppConfigVPNWithOvpn{
				VPNCountries: map[string]map[string]string{
					"nordvpn": {"France": "file1", "Germany": "file2"},
				},
			},
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := appconfig.APCFVPNSYS(sys.AppConfigVPNWithOvpn{})
			intf, err := s.ListVPN(
				tc.VPNCountries,
			)
			assert.Equal(t, tc.wantedData, intf)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
