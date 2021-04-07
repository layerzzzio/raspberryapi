package appconfig_test

import (
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/infos/appconfig"
	"github.com/raspibuddy/rpi/pkg/utl/mock"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	cases := []struct {
		name       string
		infos      mock.Infos
		apcfsys    mocksys.AppConfig
		wantedData rpi.AppConfig
		wantedErr  error
	}{
		{
			name: "success",
			infos: mock.Infos{
				VPNCountriesFn: func(string) map[string][]string {
					return map[string][]string{
						"nordvpn": {"France", "Germany"},
					}
				},
			},
			apcfsys: mocksys.AppConfig{
				ListFn: func(
					map[string][]string,
				) (rpi.AppConfig, error) {
					return rpi.AppConfig{
						VPNCountries: map[string][]string{
							"nordvpn": {"France", "Germany"},
						},
					}, nil
				},
			},
			wantedData: rpi.AppConfig{
				VPNCountries: map[string][]string{
					"nordvpn": {"France", "Germany"},
				},
			},
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := appconfig.New(&tc.apcfsys, tc.infos)
			intf, err := s.List()
			assert.Equal(t, tc.wantedData, intf)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
