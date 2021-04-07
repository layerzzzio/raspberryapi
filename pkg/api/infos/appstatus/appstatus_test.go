package appstatus_test

import (
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/infos/appstatus"
	"github.com/raspibuddy/rpi/pkg/utl/mock"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	cases := []struct {
		name       string
		infos      mock.Infos
		apsfsys    mocksys.AppStatus
		wantedData rpi.AppStatus
		wantedErr  error
	}{
		{
			name: "success",
			infos: mock.Infos{
				StatusVPNWithOpenVPNFn: func(string, string) map[string]bool {
					return map[string]bool{
						"nordvpn": true,
					}
				},
			},
			apsfsys: mocksys.AppStatus{
				ListFn: func(
					map[string]bool,
				) (rpi.AppStatus, error) {
					return rpi.AppStatus{
						VPNwithOpenVPN: map[string]bool{
							"nordvpn": true,
						},
					}, nil
				},
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
			s := appstatus.New(&tc.apsfsys, tc.infos)
			intf, err := s.List()
			assert.Equal(t, tc.wantedData, intf)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
