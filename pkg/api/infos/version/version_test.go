package version_test

import (
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/infos/version"
	"github.com/raspibuddy/rpi/pkg/utl/mock"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/stretchr/testify/assert"
)

func TestListAll(t *testing.T) {
	cases := []struct {
		name       string
		infos      mock.Infos
		vsys       mocksys.Version
		wantedData rpi.Version
		wantedErr  error
	}{
		{
			name: "success: regular version",
			infos: mock.Infos{
				ApiVersionFn: func(string, string) string {
					return "1.0.0"
				},
			},
			vsys: mocksys.Version{
				ListAllFn: func(
					string,
					string,
				) (rpi.Version, error) {
					return rpi.Version{
						RaspiBuddyVersion: "1.0.0",
					}, nil
				},
			},
			wantedData: rpi.Version{
				RaspiBuddyVersion: "1.0.0",
			},
			wantedErr: nil,
		},
		{
			name: "success: empty version",
			infos: mock.Infos{
				ApiVersionFn: func(string, string) string {
					return ""
				},
			},
			vsys: mocksys.Version{
				ListAllFn: func(
					string,
					string,
				) (rpi.Version, error) {
					return rpi.Version{}, nil
				},
			},
			wantedData: rpi.Version{},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := version.New(&tc.vsys, tc.infos)
			intf, err := s.ListAll()
			assert.Equal(t, tc.wantedData, intf)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestListAllApis(t *testing.T) {
	cases := []struct {
		name       string
		infos      mock.Infos
		vsys       mocksys.Version
		wantedData rpi.Version
		wantedErr  error
	}{
		{
			name: "success: regular version",
			infos: mock.Infos{
				ApiVersionFn: func(string, string) string {
					return "1.0.0"
				},
			},
			vsys: mocksys.Version{
				ListAllApisFn: func(
					string,
					string,
				) (rpi.Version, error) {
					return rpi.Version{
						RaspiBuddyVersion: "1.0.0",
					}, nil
				},
			},
			wantedData: rpi.Version{
				RaspiBuddyVersion: "1.0.0",
			},
			wantedErr: nil,
		},
		{
			name: "success: empty version",
			infos: mock.Infos{
				ApiVersionFn: func(string, string) string {
					return ""
				},
			},
			vsys: mocksys.Version{
				ListAllApisFn: func(
					string,
					string,
				) (rpi.Version, error) {
					return rpi.Version{}, nil
				},
			},
			wantedData: rpi.Version{},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := version.New(&tc.vsys, tc.infos)
			intf, err := s.ListAllApis()
			assert.Equal(t, tc.wantedData, intf)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
