package softwareconfig_test

import (
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/infos/softwareconfig"
	"github.com/raspibuddy/rpi/pkg/utl/mock"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	cases := []struct {
		name       string
		infos      mock.Infos
		socfsys    mocksys.SoftwareConfig
		wantedData rpi.SoftwareConfig
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
			socfsys: mocksys.SoftwareConfig{
				ListFn: func(
					map[string][]string,
				) (rpi.SoftwareConfig, error) {
					return rpi.SoftwareConfig{
						VPNCountries: map[string][]string{
							"nordvpn": {"France", "Germany"},
						},
					}, nil
				},
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
			s := softwareconfig.New(&tc.socfsys, tc.infos)
			intf, err := s.List()
			assert.Equal(t, tc.wantedData, intf)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
