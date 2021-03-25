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
				ListNameFilesInDirectoryFn: func(string) []string {
					return []string{"file1"}
				},
			},
			socfsys: mocksys.SoftwareConfig{
				ListFn: func(
					softwareconfig.NordVPN,
				) (rpi.SoftwareConfig, error) {
					return rpi.SoftwareConfig{
						NordVPN: rpi.NordVPN{
							TCPFiles: []string{"file1"},
							UDPFiles: []string{"file2"},
						},
					}, nil
				},
			},
			wantedData: rpi.SoftwareConfig{
				NordVPN: rpi.NordVPN{
					TCPFiles: []string{"file1"},
					UDPFiles: []string{"file2"},
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
