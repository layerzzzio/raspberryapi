package boot_test

import (
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/infos/boot"
	"github.com/raspibuddy/rpi/pkg/utl/mock"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	cases := []struct {
		name       string
		infos      mock.Infos
		boosys     mocksys.Boot
		wantedData rpi.Boot
		wantedErr  error
	}{
		{
			name: "file exists",
			infos: mock.Infos{
				IsFileExistsFn: func(string) bool {
					return false
				},
			},
			boosys: mocksys.Boot{
				ListFn: func(bool) (rpi.Boot, error) {
					return rpi.Boot{IsWaitForNetwork: false}, nil
				},
			},
			wantedData: rpi.Boot{IsWaitForNetwork: false},
			wantedErr:  nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := boot.New(&tc.boosys, tc.infos)
			boot, err := s.List()
			assert.Equal(t, tc.wantedData, boot)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
