package rpinterface_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/infos/rpinterface"
	"github.com/raspibuddy/rpi/pkg/utl/mock"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	cases := []struct {
		name       string
		infos      mock.Infos
		intsys     mocksys.RpInterface
		wantedData rpi.RpInterface
		wantedErr  error
	}{
		{
			name: "error: readLines nil",
			infos: mock.Infos{
				GetConfigFilesFn: func() map[string]rpi.ConfigFileDetails {
					return map[string]rpi.ConfigFileDetails{
						"etcpasswd": {
							Path: "/etc/passwd",
							Name: "passwd",
						},
					}
				},
				ReadFileFn: func(string) ([]string, error) {
					return nil, errors.New("test error readLines")
				},
				IsFileExistsFn: func(string) bool {
					return false
				},
			},
			intsys: mocksys.RpInterface{
				ListFn: func([]string, bool) (rpi.RpInterface, error) {
					return rpi.RpInterface{}, nil
				},
			},
			wantedData: rpi.RpInterface{},
			wantedErr:  echo.NewHTTPError(http.StatusInternalServerError, "could not retrieve the rpinterface details"),
		},
		{
			name: "success",
			infos: mock.Infos{
				GetConfigFilesFn: func() map[string]rpi.ConfigFileDetails {
					return map[string]rpi.ConfigFileDetails{
						"etcpasswd": {
							Path: "/etc/passwd",
							Name: "passwd",
						},
					}
				},
				ReadFileFn: func(string) ([]string, error) {
					return []string{"dummy"}, nil
				},
				IsFileExistsFn: func(string) bool {
					return false
				},
			},
			intsys: mocksys.RpInterface{
				ListFn: func([]string, bool) (rpi.RpInterface, error) {
					return rpi.RpInterface{
						IsStartXElf: true,
						IsCamera:    false,
					}, nil
				},
			},
			wantedData: rpi.RpInterface{
				IsStartXElf: true,
				IsCamera:    false,
			},
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := rpinterface.New(&tc.intsys, tc.infos)
			intf, err := s.List()
			assert.Equal(t, tc.wantedData, intf)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}