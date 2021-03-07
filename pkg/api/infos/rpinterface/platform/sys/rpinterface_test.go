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
		name               string
		readLines          []string
		isStartXElf        bool
		isSSH              bool
		isSSHKeyGenerating bool
		wantedData         rpi.RpInterface
		wantedErr          error
	}{
		{
			name: "success: start_x.elf exists",
			readLines: []string{
				"line1",
				" ",
				"     start_x     =    1  #random bash comment",
				"line3",
			},
			isStartXElf:        false,
			isSSH:              false,
			isSSHKeyGenerating: false,
			wantedData: rpi.RpInterface{
				IsStartXElf:        false,
				IsCamera:           true,
				IsSSH:              false,
				IsSSHKeyGenerating: false,
			},
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := rpinterface.INTSYS(sys.RpInterface{})
			intf, err := s.List(tc.readLines, tc.isStartXElf, tc.isSSH, tc.isSSHKeyGenerating)
			assert.Equal(t, tc.wantedData, intf)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
