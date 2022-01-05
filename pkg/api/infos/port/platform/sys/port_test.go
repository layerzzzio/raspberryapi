package sys_test

import (
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/infos/port"
	"github.com/raspibuddy/rpi/pkg/api/infos/port/platform/sys"
	"github.com/stretchr/testify/assert"
)

func TestView(t *testing.T) {
	cases := []struct {
		name       string
		port       int32
		isListen   bool
		wantedData rpi.Port
		wantedErr  error
	}{
		{
			name:     "success: port found",
			port:     6666,
			isListen: true,
			wantedData: rpi.Port{
				IsSpecificPortListen: true},
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := port.PSYS(sys.Port{})
			version, err := s.View(
				tc.port,
				tc.isListen,
			)
			assert.Equal(t, tc.wantedData, version)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
