package sys_test

import (
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/infos/software"
	"github.com/raspibuddy/rpi/pkg/api/infos/software/platform/sys"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	cases := []struct {
		name       string
		isVNC      bool
		isOpenVPN  bool
		isUnzip    bool
		wantedData rpi.Software
		wantedErr  error
	}{
		{
			name:      "success",
			isVNC:     true,
			isOpenVPN: true,
			isUnzip:   true,
			wantedData: rpi.Software{
				IsVNC:     true,
				IsOpenVPN: true,
				IsUnzip:   true,
			},
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := software.SOFSYS(sys.Software{})
			intf, err := s.List(
				tc.isVNC,
				tc.isOpenVPN,
				tc.isUnzip,
			)
			assert.Equal(t, tc.wantedData, intf)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
