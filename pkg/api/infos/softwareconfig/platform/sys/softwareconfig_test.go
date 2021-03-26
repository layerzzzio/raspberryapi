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
		name       string
		nordVPN    softwareconfig.NordVPN
		wantedData rpi.SoftwareConfig
		wantedErr  error
	}{
		{
			name: "success: isNordVPN true",
			nordVPN: softwareconfig.NordVPN{
				TCPCountries: []string{"file1"},
				UDPCountries: []string{"file2"},
			},
			wantedData: rpi.SoftwareConfig{
				NordVPN: rpi.NordVPN{
					TCPCountries: []string{"file1"},
					UDPCountries: []string{"file2"},
				},
			},
			wantedErr: nil,
		},
		{
			name: "success: isNordVPN false",
			nordVPN: softwareconfig.NordVPN{
				TCPCountries: []string{"file1"},
				UDPCountries: []string{"file2"},
			},
			wantedData: rpi.SoftwareConfig{
				NordVPN: rpi.NordVPN{
					TCPCountries: []string{"file1"},
					UDPCountries: []string{"file2"},
				},
			},
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := softwareconfig.SOCFSYS(sys.SoftwareConfig{})
			intf, err := s.List(
				tc.nordVPN,
			)
			assert.Equal(t, tc.wantedData, intf)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
