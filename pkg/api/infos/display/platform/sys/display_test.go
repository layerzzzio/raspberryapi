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
		name                    string
		readLines               []string
		isXscreenSaverInstalled bool
		isBlanking              bool
		wantedData              rpi.Display
		wantedErr               error
	}{
		{
			name: "success: overscan true",
			readLines: []string{
				"line1",
				" ",
				"     disable_overscan     =    0  #random bash comment",
				"line3",
			},
			isXscreenSaverInstalled: false,
			isBlanking:              false,
			wantedData: rpi.Display{
				IsOverscan:              true,
				IsXscreenSaverInstalled: false,
				IsBlanking:              false,
			},
			wantedErr: nil,
		},
		{
			name: "success: overscan false",
			readLines: []string{
				"line1",
				"#disable_overscan=1",
				"line3",
			},
			isXscreenSaverInstalled: true,
			isBlanking:              true,
			wantedData: rpi.Display{
				IsOverscan:              false,
				IsXscreenSaverInstalled: true,
				IsBlanking:              true,
			},
			wantedErr: nil,
		},
		{
			name: "success: overscan false with whitespaces",
			readLines: []string{
				"line1",
				"  #disable_overscan =   1 ",
				"line3",
			},
			isXscreenSaverInstalled: false,
			isBlanking:              false,
			wantedData: rpi.Display{
				IsOverscan:              false,
				IsXscreenSaverInstalled: false,
				IsBlanking:              false,
			},
			wantedErr: nil,
		},
		{
			name:                    "success: arg is nil",
			readLines:               nil,
			isXscreenSaverInstalled: true,
			isBlanking:              true,
			wantedData: rpi.Display{
				IsOverscan:              false,
				IsXscreenSaverInstalled: true,
				IsBlanking:              true,
			},
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := display.DISSYS(sys.Display{})
			boot, err := s.List(tc.readLines, tc.isXscreenSaverInstalled, tc.isBlanking)
			assert.Equal(t, tc.wantedData, boot)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
