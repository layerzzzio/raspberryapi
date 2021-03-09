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
		name                string
		readLines           []string
		isStartXElf         bool
		isSSH               bool
		isSSHKeyGenerating  bool
		isVNC               bool
		isVNCInstalledCheck bool
		isSPI               bool
		isI2C               bool
		wantedData          rpi.RpInterface
		wantedErr           error
	}{
		{
			name: "success: start_x.elf exists",
			readLines: []string{
				"line1",
				" ",
				"     start_x     =    1  #random bash comment",
				"line3",
			},
			isStartXElf:         false,
			isSSH:               false,
			isSSHKeyGenerating:  false,
			isVNC:               true,
			isVNCInstalledCheck: true,
			isSPI:               true,
			isI2C:               true,
			wantedData: rpi.RpInterface{
				IsStartXElf:        false,
				IsCamera:           true,
				IsSSH:              false,
				IsSSHKeyGenerating: false,
				IsVNC:              true,
				IsVNCInstalled:     true,
				IsSPI:              true,
				IsI2C:              true,
			},
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := rpinterface.INTSYS(sys.RpInterface{})
			intf, err := s.List(
				tc.readLines,
				tc.isStartXElf,
				tc.isSSH,
				tc.isSSHKeyGenerating,
				tc.isVNC,
				tc.isVNCInstalledCheck,
				tc.isSPI,
				tc.isI2C,
			)
			assert.Equal(t, tc.wantedData, intf)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
