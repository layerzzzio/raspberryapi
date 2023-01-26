package sys_test

import (
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/infos/boot"
	"github.com/raspibuddy/rpi/pkg/api/infos/boot/platform/sys"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	cases := []struct {
		name             string
		isWaitForNetwork bool
		wantedData       rpi.Boot
		wantedErr        error
	}{
		{
			name:             "success",
			isWaitForNetwork: true,
			wantedData:       rpi.Boot{IsWaitForNetwork: true},
			wantedErr:        nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := boot.BOOSYS(sys.Boot{})
			boot, err := s.List(tc.isWaitForNetwork)
			assert.Equal(t, tc.wantedData, boot)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
