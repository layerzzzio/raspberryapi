package rpinterface_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
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
				IsQuietGrepFn: func(string, string, string) bool {
					return false
				},
				IsSSHKeyGeneratingFn: func(string) bool {
					return false
				},
				IsDPKGInstalledFn: func(string) bool {
					return false
				},
				IsSPIFn: func(string) bool {
					return false
				},
				IsI2CFn: func(string) bool {
					return false
				},
				IsVariableSetFn: func([]string, string, string) bool {
					return false
				},
				ListWifiInterfacesFn: func(string) []string {
					return []string{}
				},
				IsWpaSupComFn: func() map[string]bool {
					return map[string]bool{}
				},
				ZoneInfoFn: func(string) map[string]string {
					return map[string]string{}
				},
			},
			intsys: mocksys.RpInterface{
				ListFn: func(
					[]string,
					bool,
					bool,
					bool,
					bool,
					bool,
					bool,
					bool,
					bool,
					bool,
					[]string,
					map[string]bool,
					map[string]string,
				) (rpi.RpInterface, error) {
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
				IsQuietGrepFn: func(string, string, string) bool {
					return false
				},
				IsSSHKeyGeneratingFn: func(string) bool {
					return false
				},
				IsDPKGInstalledFn: func(string) bool {
					return false
				},
				IsSPIFn: func(string) bool {
					return false
				},
				IsI2CFn: func(string) bool {
					return false
				},
				IsVariableSetFn: func([]string, string, string) bool {
					return false
				},
				ListWifiInterfacesFn: func(string) []string {
					return []string{}
				},
				IsWpaSupComFn: func() map[string]bool {
					return map[string]bool{}
				},
				ZoneInfoFn: func(string) map[string]string {
					return map[string]string{"FR": "France"}
				},
			},
			intsys: mocksys.RpInterface{
				ListFn: func(
					[]string,
					bool,
					bool,
					bool,
					bool,
					bool,
					bool,
					bool,
					bool,
					bool,
					[]string,
					map[string]bool,
					map[string]string,
				) (rpi.RpInterface, error) {
					return rpi.RpInterface{
						IsStartXElf:        true,
						IsCamera:           false,
						IsSSH:              true,
						IsSSHKeyGenerating: true,
						IsVNC:              true,
						IsSPI:              true,
						IsI2C:              true,
						IsVNCInstalled:     true,
						IsOneWire:          true,
						IsRemoteGpio:       true,
						IsWifiInterfaces:   false,
						IsWpaSupCom:        map[string]bool{"wlan0": true},
						ZoneInfo:           map[string]string{"FR": "France"},
					}, nil
				},
			},
			wantedData: rpi.RpInterface{
				IsStartXElf:        true,
				IsCamera:           false,
				IsSSH:              true,
				IsSSHKeyGenerating: true,
				IsVNC:              true,
				IsSPI:              true,
				IsI2C:              true,
				IsVNCInstalled:     true,
				IsOneWire:          true,
				IsRemoteGpio:       true,
				IsWifiInterfaces:   false,
				IsWpaSupCom:        map[string]bool{"wlan0": true},
				ZoneInfo:           map[string]string{"FR": "France"},
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
