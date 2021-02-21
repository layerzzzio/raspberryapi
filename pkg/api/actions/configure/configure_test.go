package configure_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/actions/configure"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
	"github.com/raspibuddy/rpi/pkg/utl/mock"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/stretchr/testify/assert"
)

func TestExecuteCH(t *testing.T) {
	cases := []struct {
		name       string
		path       string
		plan       map[int](map[int]actions.Func)
		actions    *mock.Actions
		consys     *mocksys.Action
		wantedData rpi.Action
		wantedErr  error
	}{
		{
			name: "success",
			path: "/dummy",
			plan: map[int](map[int]actions.Func){
				1: {
					1: {
						Name:      actions.ChangeHostnameInHostsFile,
						Reference: actions.ChangeHostnameInHostsFile,
						Argument: []interface{}{
							actions.DataToFile{
								TargetFile: "path",
								Data:       "data",
							},
						},
					},
					2: {
						Name:      actions.ChangeHostnameInHostnameFile,
						Reference: actions.ChangeHostnameInHostnameFile,
						Argument: []interface{}{
							actions.DataToFile{
								TargetFile: "path",
								Data:       "data",
							},
						},
					},
				},
			},
			actions: &mock.Actions{
				ChangeHostnameInHostnameFileFn: func(interface{}) (rpi.Exec, error) {
					return rpi.Exec{
						Name:       actions.ChangeHostnameInHostsFile,
						StartTime:  1,
						EndTime:    2,
						ExitStatus: 0,
						Stdout:     "path-data",
					}, nil
				},
				ChangeHostnameInHostsFileFn: func(interface{}) (rpi.Exec, error) {
					return rpi.Exec{
						Name:       actions.ChangeHostnameInHostnameFile,
						StartTime:  1,
						EndTime:    2,
						ExitStatus: 0,
						Stdout:     "path-data",
					}, nil
				},
			},
			consys: &mocksys.Action{
				ExecuteCHFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{
						Name:          actions.ChangeHostname,
						NumberOfSteps: 1,
						Progress: map[string]rpi.Exec{
							"1": {
								Name:       actions.ChangeHostnameInHostsFile,
								StartTime:  1,
								EndTime:    2,
								ExitStatus: 0,
								Stdout:     "path-data",
							},
							"2": {
								Name:       actions.ChangeHostnameInHostnameFile,
								StartTime:  1,
								EndTime:    2,
								ExitStatus: 0,
								Stdout:     "path-data",
							},
						},
						ExitStatus: 0,
						StartTime:  2,
						EndTime:    uint64(time.Now().Unix()),
					}, nil
				},
			},
			wantedData: rpi.Action{
				Name:          actions.ChangeHostname,
				NumberOfSteps: 1,
				Progress: map[string]rpi.Exec{
					"1": {
						Name:       actions.ChangeHostnameInHostsFile,
						StartTime:  1,
						EndTime:    2,
						ExitStatus: 0,
						Stdout:     "path-data",
					},
					"2": {
						Name:       actions.ChangeHostnameInHostnameFile,
						StartTime:  1,
						EndTime:    2,
						ExitStatus: 0,
						Stdout:     "path-data",
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
			s := configure.New(tc.consys, tc.actions)
			changeHostname, err := s.ExecuteCH(tc.path)
			assert.Equal(t, tc.wantedData, changeHostname)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestExecuteCP(t *testing.T) {
	cases := []struct {
		name       string
		password   string
		username   string
		plan       map[int](map[int]actions.Func)
		actions    *mock.Actions
		consys     *mocksys.Action
		wantedData rpi.Action
		wantedErr  error
	}{
		{
			name:     "success",
			password: "dummypassword",
			username: "dummyusername",
			plan: map[int](map[int]actions.Func){
				1: {
					1: {
						Name:      actions.ChangePassword,
						Reference: actions.ChangePassword,
						Argument: []interface{}{
							actions.CP{
								Password: "password",
								Username: "username",
							},
						},
					},
				},
			},
			actions: &mock.Actions{
				ChangePasswordFn: func(interface{}) (rpi.Exec, error) {
					return rpi.Exec{
						Name:       actions.ChangePassword,
						StartTime:  1,
						EndTime:    2,
						ExitStatus: 0,
						Stdout:     "password-username",
					}, nil
				},
			},
			consys: &mocksys.Action{
				ExecuteCPFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{
						Name:          actions.ChangePassword,
						NumberOfSteps: 1,
						Progress: map[string]rpi.Exec{
							"1": {
								Name:       actions.ChangePassword,
								StartTime:  1,
								EndTime:    2,
								ExitStatus: 0,
								Stdout:     "password-username",
							},
						},
						ExitStatus: 0,
						StartTime:  2,
						EndTime:    uint64(time.Now().Unix()),
					}, nil
				},
			},
			wantedData: rpi.Action{
				Name:          actions.ChangePassword,
				NumberOfSteps: 1,
				Progress: map[string]rpi.Exec{
					"1": {
						Name:       actions.ChangePassword,
						StartTime:  1,
						EndTime:    2,
						ExitStatus: 0,
						Stdout:     "password-username",
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
			s := configure.New(tc.consys, tc.actions)
			changePassword, err := s.ExecuteCP(tc.password, tc.username)
			assert.Equal(t, tc.wantedData, changePassword)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestExecuteWNB(t *testing.T) {
	cases := []struct {
		name       string
		action     string
		plan       map[int](map[int]actions.Func)
		actions    *mock.Actions
		consys     *mocksys.Action
		wantedData rpi.Action
		wantedErr  error
	}{
		{
			name:   "success",
			action: "dummyaction",
			plan: map[int](map[int]actions.Func){
				1: {
					1: {
						Name:      actions.WaitForNetworkAtBoot,
						Reference: actions.WaitForNetworkAtBoot,
						Argument: []interface{}{
							actions.EnableOrDisableConfig{
								DirOrFilePath: "directory",
								Action:        "enable",
							},
						},
					},
				},
			},
			actions: &mock.Actions{
				WaitForNetworkAtBootFn: func(interface{}) (rpi.Exec, error) {
					return rpi.Exec{
						Name:       actions.WaitForNetworkAtBoot,
						StartTime:  1,
						EndTime:    2,
						ExitStatus: 0,
						Stdout:     "directory-enable",
					}, nil
				},
			},
			consys: &mocksys.Action{
				ExecuteWNBFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{
						Name:          actions.WaitForNetworkAtBoot,
						NumberOfSteps: 1,
						Progress: map[string]rpi.Exec{
							"1": {
								Name:       actions.WaitForNetworkAtBoot,
								StartTime:  1,
								EndTime:    2,
								ExitStatus: 0,
								Stdout:     "directory-enable",
							},
						},
						ExitStatus: 0,
						StartTime:  2,
						EndTime:    uint64(time.Now().Unix()),
					}, nil
				},
			},
			wantedData: rpi.Action{
				Name:          actions.WaitForNetworkAtBoot,
				NumberOfSteps: 1,
				Progress: map[string]rpi.Exec{
					"1": {
						Name:       actions.WaitForNetworkAtBoot,
						StartTime:  1,
						EndTime:    2,
						ExitStatus: 0,
						Stdout:     "directory-enable",
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
			s := configure.New(tc.consys, tc.actions)
			changePassword, err := s.ExecuteWNB(tc.action)
			assert.Equal(t, tc.wantedData, changePassword)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestExecuteOV(t *testing.T) {
	cases := []struct {
		name       string
		action     string
		plan       map[int](map[int]actions.Func)
		actions    *mock.Actions
		consys     *mocksys.Action
		wantedData rpi.Action
		wantedErr  error
	}{
		{
			name:   "error",
			action: "enable-xxx",
			plan: map[int](map[int]actions.Func){
				1: {
					1: {
						Name:      actions.DisableOrEnableOverscan,
						Reference: actions.DisableOrEnableOverscan,
						Argument: []interface{}{
							actions.EnableOrDisableConfig{
								DirOrFilePath: "path",
								Action:        "enable",
							},
						},
					},
				},
			},
			wantedData: rpi.Action{},
			wantedErr:  echo.NewHTTPError(http.StatusInternalServerError, "bad action type: enable or disable overscan failed"),
		},
		{
			name:   "success",
			action: "enable",
			plan: map[int](map[int]actions.Func){
				1: {
					1: {
						Name:      actions.DisableOrEnableOverscan,
						Reference: actions.DisableOrEnableOverscan,
						Argument: []interface{}{
							actions.EnableOrDisableConfig{
								DirOrFilePath: "path",
								Action:        "enable",
							},
						},
					},
				},
			},
			actions: &mock.Actions{
				DisableOrEnableOverscanFn: func(interface{}) (rpi.Exec, error) {
					return rpi.Exec{
						Name:       actions.DisableOrEnableOverscan,
						StartTime:  1,
						EndTime:    2,
						ExitStatus: 0,
						Stdout:     "path-enable",
					}, nil
				},
			},
			consys: &mocksys.Action{
				ExecuteOVFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{
						Name:          actions.Overscan,
						NumberOfSteps: 1,
						Progress: map[string]rpi.Exec{
							"1": {
								Name:       actions.DisableOrEnableOverscan,
								StartTime:  1,
								EndTime:    2,
								ExitStatus: 0,
								Stdout:     "path-enable",
							},
						},
						ExitStatus: 0,
						StartTime:  2,
						EndTime:    uint64(time.Now().Unix()),
					}, nil
				},
			},
			wantedData: rpi.Action{
				Name:          actions.Overscan,
				NumberOfSteps: 1,
				Progress: map[string]rpi.Exec{
					"1": {
						Name:       actions.DisableOrEnableOverscan,
						StartTime:  1,
						EndTime:    2,
						ExitStatus: 0,
						Stdout:     "path-enable",
					},
				},
				ExitStatus: 0,
				StartTime:  2,
				EndTime:    uint64(time.Now().Unix()),
			},
			wantedErr: nil,
		},
		{
			name:   "success",
			action: "disable",
			plan: map[int](map[int]actions.Func){
				1: {
					1: {
						Name:      actions.DisableOrEnableOverscan,
						Reference: actions.DisableOrEnableOverscan,
						Argument: []interface{}{
							actions.EnableOrDisableConfig{
								DirOrFilePath: "path",
								Action:        "enable",
							},
						},
					},
				},
				2: {
					1: {
						Name:      actions.CommentOverscan,
						Reference: actions.CommentOverscan,
						Argument: []interface{}{
							actions.CommentOrUncommentConfig{
								DirOrFilePath: "path",
								Action:        "comment",
							},
						},
					},
				},
			},
			actions: &mock.Actions{
				DisableOrEnableOverscanFn: func(interface{}) (rpi.Exec, error) {
					return rpi.Exec{
						Name:       actions.DisableOrEnableOverscan,
						StartTime:  1,
						EndTime:    2,
						ExitStatus: 0,
						Stdout:     "path-enable",
					}, nil
				},
				CommentOverscanFn: func(interface{}) (rpi.Exec, error) {
					return rpi.Exec{
						Name:       actions.CommentOverscan,
						StartTime:  1,
						EndTime:    2,
						ExitStatus: 0,
						Stdout:     "path-comment",
					}, nil
				},
			},
			consys: &mocksys.Action{
				ExecuteOVFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{
						Name:          actions.Overscan,
						NumberOfSteps: 1,
						Progress: map[string]rpi.Exec{
							"1": {
								Name:       actions.DisableOrEnableOverscan,
								StartTime:  1,
								EndTime:    2,
								ExitStatus: 0,
								Stdout:     "path-enable",
							},
							"2": {
								Name:       actions.CommentOverscan,
								StartTime:  1,
								EndTime:    2,
								ExitStatus: 0,
								Stdout:     "path-comment",
							},
						},
						ExitStatus: 0,
						StartTime:  2,
						EndTime:    uint64(time.Now().Unix()),
					}, nil
				},
			},
			wantedData: rpi.Action{
				Name:          actions.Overscan,
				NumberOfSteps: 1,
				Progress: map[string]rpi.Exec{
					"1": {
						Name:       actions.DisableOrEnableOverscan,
						StartTime:  1,
						EndTime:    2,
						ExitStatus: 0,
						Stdout:     "path-enable",
					},
					"2": {
						Name:       actions.CommentOverscan,
						StartTime:  1,
						EndTime:    2,
						ExitStatus: 0,
						Stdout:     "path-comment",
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
			s := configure.New(tc.consys, tc.actions)
			overscan, err := s.ExecuteOV(tc.action)
			assert.Equal(t, tc.wantedData, overscan)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
