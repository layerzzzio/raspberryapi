package sys_test

import (
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/infos/rpinterface"
	"github.com/raspibuddy/rpi/pkg/api/infos/rpinterface/platform/sys"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	cases := []struct {
		name        string
		readLines   []string
		isStartXElf bool
		wantedData  rpi.RpInterface
		wantedErr   error
	}{
		{
			name: "success: start_x.elf exists",
			readLines: []string{
				"line1",
				" ",
				"     start_x     =    1  #random bash comment",
				"line3",
			},
			isStartXElf: false,
			wantedData: rpi.RpInterface{
				IsStartXElf: false,
				IsCamera:    true,
			},
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := rpinterface.INTSYS(sys.RpInterface{})
			intf, err := s.List(tc.readLines, tc.isStartXElf)
			assert.Equal(t, tc.wantedData, intf)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
