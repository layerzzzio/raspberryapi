package appaction_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/actions/appaction"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
	"github.com/raspibuddy/rpi/pkg/utl/mock"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/raspibuddy/rpi/pkg/utl/test_utl"
	"github.com/stretchr/testify/assert"
)

func TestExecuteWOVA(t *testing.T) {
	cases := []struct {
		name               string
		action             string
		vpnName            string
		relativeConfigPath string
		country            string
		username           string
		password           string
		plan               map[int](map[int]actions.Func)
		actions            *mock.Actions
		infos              *mock.Infos
		aacsys             *mocksys.Action
		wantedData         rpi.Action
		wantedErr          error
	}{
		{
			name:   "bad action type",
			action: "connectXXX",
			plan: map[int](map[int]actions.Func){
				1: {
					1: {
						Name:      "FuncA",
						Reference: test_utl.FuncA,
						Argument: []interface{}{
							test_utl.ArgFuncA{
								Arg0: "string0",
								Arg1: "string1",
							},
						},
					},
				},
			},
			actions: &mock.Actions{
				ExecuteBashCommandFn: func(interface{}) (rpi.Exec, error) {
					return rpi.Exec{
						Name:       "FuncA",
						StartTime:  1,
						EndTime:    2,
						ExitStatus: 0,
						Stdout:     "string0-string1",
					}, nil
				},
				KillProcessFn: func(interface{}) (rpi.Exec, error) {
					return rpi.Exec{
						Name:       "FuncA",
						StartTime:  1,
						EndTime:    2,
						ExitStatus: 0,
						Stdout:     "string0-string1",
					}, nil
				},
				ConfirmVPNAuthenticationFn: func(interface{}) (rpi.Exec, error) {
					return rpi.Exec{
						Name:       "FuncA",
						StartTime:  1,
						EndTime:    2,
						ExitStatus: 0,
						Stdout:     "string0-string1",
					}, nil
				},
			},
			infos: &mock.Infos{
				VPNConfigFileFn: func(string, string, string) []string {
					return nil
				},
				ProcessesPidsFn: func(string) []string {
					return nil
				},
			},
			aacsys: &mocksys.Action{
				ExecuteWOVAFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{}, echo.NewHTTPError(http.StatusInternalServerError, "bad action type: connect or disconnect vpn with openvpn failed")
				},
			},
			wantedData: rpi.Action{},
			wantedErr:  echo.NewHTTPError(http.StatusInternalServerError, "bad action type: connect or disconnect vpn with openvpn failed"),
		},
		{
			name:               "success connect",
			action:             "connect",
			vpnName:            "surfshark",
			relativeConfigPath: "",
			country:            "France",
			username:           "loic",
			password:           "pass",
			plan: map[int](map[int]actions.Func){
				1: {
					1: {
						Name:      "FuncA",
						Reference: test_utl.FuncA,
						Argument: []interface{}{
							test_utl.ArgFuncA{
								Arg0: "string0",
								Arg1: "string1",
							},
						},
					},
				},
			},
			actions: &mock.Actions{
				ExecuteBashCommandFn: func(interface{}) (rpi.Exec, error) {
					return rpi.Exec{
						Name:       "FuncA",
						StartTime:  1,
						EndTime:    2,
						ExitStatus: 0,
						Stdout:     "string0-string1",
					}, nil
				},
				KillProcessFn: func(interface{}) (rpi.Exec, error) {
					return rpi.Exec{
						Name:       "FuncA",
						StartTime:  1,
						EndTime:    2,
						ExitStatus: 0,
						Stdout:     "string0-string1",
					}, nil
				},
				ConfirmVPNAuthenticationFn: func(interface{}) (rpi.Exec, error) {
					return rpi.Exec{
						Name:       "FuncA",
						StartTime:  1,
						EndTime:    2,
						ExitStatus: 0,
						Stdout:     "string0-string1",
					}, nil
				},
			},
			infos: &mock.Infos{
				VPNConfigFileFn: func(string, string, string) []string {
					return []string{"france.opvn", "england.opvn"}
				},
				ProcessesPidsFn: func(string) []string {
					return nil
				},
			},
			aacsys: &mocksys.Action{
				ExecuteWOVAFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{
						Name:          "FuncA",
						NumberOfSteps: 1,
						Progress: map[string]rpi.Exec{
							"1": {
								Name:       "FuncA",
								StartTime:  1,
								EndTime:    2,
								ExitStatus: 0,
								Stdout:     "string0-string1",
							},
						},
						ExitStatus: 0,
						StartTime:  2,
						EndTime:    uint64(time.Now().Unix()),
					}, nil
				},
			},
			wantedData: rpi.Action{
				Name:          "FuncA",
				NumberOfSteps: 1,
				Progress: map[string]rpi.Exec{
					"1": {
						Name:       "FuncA",
						StartTime:  1,
						EndTime:    2,
						ExitStatus: 0,
						Stdout:     "string0-string1",
					},
				},
				ExitStatus: 0,
				StartTime:  2,
				EndTime:    uint64(time.Now().Unix()),
			},
			wantedErr: nil,
		},
		{
			name:     "success disconnect",
			action:   "disconnect",
			vpnName:  "",
			country:  "",
			username: "",
			password: "",
			plan: map[int](map[int]actions.Func){
				1: {
					1: {
						Name:      "FuncA",
						Reference: test_utl.FuncA,
						Argument: []interface{}{
							test_utl.ArgFuncA{
								Arg0: "string0",
								Arg1: "string1",
							},
						},
					},
				},
			},
			actions: &mock.Actions{
				ExecuteBashCommandFn: func(interface{}) (rpi.Exec, error) {
					return rpi.Exec{
						Name:       "FuncA",
						StartTime:  1,
						EndTime:    2,
						ExitStatus: 0,
						Stdout:     "string0-string1",
					}, nil
				},
				KillProcessFn: func(interface{}) (rpi.Exec, error) {
					return rpi.Exec{
						Name:       "FuncA",
						StartTime:  1,
						EndTime:    2,
						ExitStatus: 0,
						Stdout:     "string0-string1",
					}, nil
				},
				ConfirmVPNAuthenticationFn: func(interface{}) (rpi.Exec, error) {
					return rpi.Exec{
						Name:       "FuncA",
						StartTime:  1,
						EndTime:    2,
						ExitStatus: 0,
						Stdout:     "string0-string1",
					}, nil
				},
			},
			infos: &mock.Infos{
				VPNConfigFileFn: func(string, string, string) []string {
					return nil
				},
				ProcessesPidsFn: func(string) []string {
					return []string{"122", "222"}
				},
			},
			aacsys: &mocksys.Action{
				ExecuteWOVAFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{
						Name:          "FuncA",
						NumberOfSteps: 1,
						Progress: map[string]rpi.Exec{
							"1": {
								Name:       "FuncA",
								StartTime:  1,
								EndTime:    2,
								ExitStatus: 0,
								Stdout:     "string0-string1",
							},
						},
						ExitStatus: 0,
						StartTime:  2,
						EndTime:    uint64(time.Now().Unix()),
					}, nil
				},
			},
			wantedData: rpi.Action{
				Name:          "FuncA",
				NumberOfSteps: 1,
				Progress: map[string]rpi.Exec{
					"1": {
						Name:       "FuncA",
						StartTime:  1,
						EndTime:    2,
						ExitStatus: 0,
						Stdout:     "string0-string1",
					},
				},
				ExitStatus: 0,
				StartTime:  2,
				EndTime:    uint64(time.Now().Unix()),
			},
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := appaction.New(tc.aacsys, tc.actions, tc.infos)
			vpnAction, err := s.ExecuteWOVA(
				tc.action,
				tc.vpnName,
				tc.relativeConfigPath,
				tc.username,
				tc.password,
				tc.country,
			)
			assert.Equal(t, tc.wantedData, vpnAction)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
