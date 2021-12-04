package sys_test

import (
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/infos/version"
	"github.com/raspibuddy/rpi/pkg/api/infos/version/platform/sys"
	"github.com/stretchr/testify/assert"
)

func TestListAll(t *testing.T) {
	cases := []struct {
		name                    string
		raspibuddyVersion       string
		raspibuddyDeployVersion string
		wantedData              rpi.Version
		wantedErr               error
	}{
		{
			name:                    "success: regular version",
			raspibuddyVersion:       "1.0.0",
			raspibuddyDeployVersion: "1.1.1",
			wantedData: rpi.Version{
				RaspiBuddyVersion:       "1.0.0",
				RaspiBuddyDeployVersion: "1.1.1",
			},
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := version.VSYS(sys.Version{})
			version, err := s.ListAll(
				tc.raspibuddyVersion,
				tc.raspibuddyDeployVersion,
			)
			assert.Equal(t, tc.wantedData, version)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestListAllVersion(t *testing.T) {
	cases := []struct {
		name                    string
		raspibuddyVersion       string
		raspibuddyDeployVersion string
		wantedData              rpi.Version
		wantedErr               error
	}{
		{
			name:                    "success: regular version",
			raspibuddyVersion:       "1.0.0",
			raspibuddyDeployVersion: "1.1.1",
			wantedData: rpi.Version{
				RaspiBuddyVersion:       "1.0.0",
				RaspiBuddyDeployVersion: "1.1.1",
			},
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := version.VSYS(sys.Version{})
			version, err := s.ListAllApis(
				tc.raspibuddyVersion,
				tc.raspibuddyDeployVersion,
			)
			assert.Equal(t, tc.wantedData, version)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
