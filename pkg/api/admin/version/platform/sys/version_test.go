package sys_test

import (
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/admin/version"
	"github.com/raspibuddy/rpi/pkg/api/admin/version/platform/sys"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	cases := []struct {
		name       string
		apiVersion string
		wantedData rpi.Version
		wantedErr  error
	}{
		{
			name:       "success: regular version",
			apiVersion: "1.0.0",
			wantedData: rpi.Version{
				ApiVersion: "1.0.0",
			},
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := version.VSYS(sys.Version{})
			version, err := s.List(tc.apiVersion)
			assert.Equal(t, tc.wantedData, version)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
