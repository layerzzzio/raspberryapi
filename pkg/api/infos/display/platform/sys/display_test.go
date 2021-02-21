package sys_test

import (
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/infos/display"
	"github.com/raspibuddy/rpi/pkg/api/infos/display/platform/sys"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	cases := []struct {
		name       string
		readLines  []string
		wantedData rpi.Display
		wantedErr  error
	}{
		{
			name: "success: overscan true",
			readLines: []string{
				"line1",
				"disable_overscan=0",
				"line3",
			},
			wantedData: rpi.Display{IsOverscan: true},
			wantedErr:  nil,
		},
		{
			name: "success: overscan false",
			readLines: []string{
				"line1",
				"#disable_overscan=1",
				"line3",
			},
			wantedData: rpi.Display{IsOverscan: false},
			wantedErr:  nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := display.DISSYS(sys.Display{})
			boot, err := s.List(tc.readLines)
			assert.Equal(t, tc.wantedData, boot)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
