package software_test

import (
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/infos/software"
	"github.com/raspibuddy/rpi/pkg/utl/mock"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	cases := []struct {
		name       string
		infos      mock.Infos
		sofsys     mocksys.Software
		wantedData rpi.Software
		wantedErr  error
	}{
		{
			name: "success",
			infos: mock.Infos{
				IsDPKGInstalledFn: func(string) bool {
					return false
				},
				HasDirectoryAtLeastOneFileFn: func(string, bool) bool {
					return false
				},
			},
			sofsys: mocksys.Software{
				ListFn: func(
					bool,
					bool,
					bool,
					bool,
					bool,
					bool,
					bool,
				) (rpi.Software, error) {
					return rpi.Software{
						IsVNCInstalled:     false,
						IsOpenVPNInstalled: false,
						IsUnzipInstalled:   true,
						IsNordVPNInstalled: true,
					}, nil
				},
			},
			wantedData: rpi.Software{
				IsVNCInstalled:     false,
				IsOpenVPNInstalled: false,
				IsUnzipInstalled:   true,
				IsNordVPNInstalled: true,
			},
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := software.New(&tc.sofsys, tc.infos)
			intf, err := s.List()
			assert.Equal(t, tc.wantedData, intf)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
