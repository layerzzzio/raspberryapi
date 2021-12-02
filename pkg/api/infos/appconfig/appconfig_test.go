package appconfig_test

import (
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/infos/appconfig"
	"github.com/raspibuddy/rpi/pkg/utl/mock"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/stretchr/testify/assert"
)

func TestListVPN(t *testing.T) {
	cases := []struct {
		name       string
		infos      mock.Infos
		apcfsys    mocksys.AppConfigVPNWithOvpn
		wantedData rpi.AppConfigVPNWithOvpn
		wantedErr  error
	}{
		{
			name: "success",
			infos: mock.Infos{
				VPNCountriesFn: func(string) map[string](map[string]string) {
					return map[string](map[string]string){
						"nordvpn": {"France": "file1", "Germany": "file2"},
					}
				},
			},
			apcfsys: mocksys.AppConfigVPNWithOvpn{
				ListVPNFn: func(
					map[string](map[string]string),
				) (rpi.AppConfigVPNWithOvpn, error) {
					return rpi.AppConfigVPNWithOvpn{
						VPNCountries: map[string]map[string]string{
							"nordvpn": {"France": "file1", "Germany": "file2"},
						},
					}, nil
				},
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
			s := appconfig.New(&tc.apcfsys, tc.infos)
			intf, err := s.ListVPN()
			assert.Equal(t, tc.wantedData, intf)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
