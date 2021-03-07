package configure_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
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
		infos      *mock.Infos
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
			infos: &mock.Infos{
				GetConfigFilesFn: func() map[string]rpi.ConfigFileDetails {
					return map[string]rpi.ConfigFileDetails{
						"bootconfig": {
							Path: "/dummy/path",
						},
					}
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
			s := configure.New(tc.consys, tc.actions, tc.infos)
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
		infos      *mock.Infos
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
			infos: &mock.Infos{
				GetConfigFilesFn: func() map[string]rpi.ConfigFileDetails {
					return map[string]rpi.ConfigFileDetails{
						"bootconfig": {
							Path: "/dummy/path",
						},
					}
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
			s := configure.New(tc.consys, tc.actions, tc.infos)
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
		infos      *mock.Infos
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
			infos: &mock.Infos{
				GetConfigFilesFn: func() map[string]rpi.ConfigFileDetails {
					return map[string]rpi.ConfigFileDetails{
						"bootconfig": {
							Path: "/dummy/path",
						},
					}
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
			s := configure.New(tc.consys, tc.actions, tc.infos)
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
		infos      *mock.Infos
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
			infos: &mock.Infos{
				GetConfigFilesFn: func() map[string]rpi.ConfigFileDetails {
					return map[string]rpi.ConfigFileDetails{
						"bootconfig": {
							Path: "/dummy/path",
						},
					}
				},
			},
			actions: &mock.Actions{
				DisableOrEnableConfigFn: func(interface{}) (rpi.Exec, error) {
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
			infos: &mock.Infos{
				GetConfigFilesFn: func() map[string]rpi.ConfigFileDetails {
					return map[string]rpi.ConfigFileDetails{
						"bootconfig": {
							Path: "/dummy/path",
						},
					}
				},
			},
			actions: &mock.Actions{
				DisableOrEnableConfigFn: func(interface{}) (rpi.Exec, error) {
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
			s := configure.New(tc.consys, tc.actions, tc.infos)
			overscan, err := s.ExecuteOV(tc.action)
			assert.Equal(t, tc.wantedData, overscan)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestExecuteBL(t *testing.T) {
	cases := []struct {
		name       string
		action     string
		plan       map[int](map[int]actions.Func)
		actions    *mock.Actions
		infos      *mock.Infos
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
						Name:      actions.DisableOrEnableBlanking,
						Reference: actions.DisableOrEnableBlanking,
						Argument: []interface{}{
							actions.TargetDestEnableOrDisableConfig{
								TargetDirOrFilePath:      "path",
								DestinationDirOrFilePath: "destination",
								Action:                   "enable",
							},
						},
					},
				},
			},
			wantedData: rpi.Action{},
			wantedErr:  echo.NewHTTPError(http.StatusInternalServerError, "bad action type: enable or disable blanking failed"),
		},
		{
			name:   "success",
			action: "enable",
			plan: map[int](map[int]actions.Func){
				1: {
					1: {
						Name:      actions.DisableOrEnableBlanking,
						Reference: actions.DisableOrEnableBlanking,
						Argument: []interface{}{
							actions.TargetDestEnableOrDisableConfig{
								TargetDirOrFilePath:      "path",
								DestinationDirOrFilePath: "destination",
								Action:                   "enable",
							},
						},
					},
				},
			},
			actions: &mock.Actions{
				DisableOrEnableConfigFn: func(interface{}) (rpi.Exec, error) {
					return rpi.Exec{
						Name:       actions.DisableOrEnableBlanking,
						StartTime:  1,
						EndTime:    2,
						ExitStatus: 0,
						Stdout:     "path-destination-enable",
					}, nil
				},
			},
			consys: &mocksys.Action{
				ExecuteBLFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{
						Name:          actions.Blanking,
						NumberOfSteps: 1,
						Progress: map[string]rpi.Exec{
							"1": {
								Name:       actions.DisableOrEnableBlanking,
								StartTime:  1,
								EndTime:    2,
								ExitStatus: 0,
								Stdout:     "path-destination-enable",
							},
						},
						ExitStatus: 0,
						StartTime:  2,
						EndTime:    uint64(time.Now().Unix()),
					}, nil
				},
			},
			wantedData: rpi.Action{
				Name:          actions.Blanking,
				NumberOfSteps: 1,
				Progress: map[string]rpi.Exec{
					"1": {
						Name:       actions.DisableOrEnableBlanking,
						StartTime:  1,
						EndTime:    2,
						ExitStatus: 0,
						Stdout:     "path-destination-enable",
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
			s := configure.New(tc.consys, tc.actions, tc.infos)
			blanking, err := s.ExecuteBL(tc.action)
			assert.Equal(t, tc.wantedData, blanking)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestExecuteAUS(t *testing.T) {
	cases := []struct {
		name       string
		username   string
		password   string
		plan       map[int](map[int]actions.Func)
		actions    *mock.Actions
		infos      *mock.Infos
		consys     *mocksys.Action
		wantedData rpi.Action
		wantedErr  error
	}{
		{
			name:     "success",
			username: "username",
			password: "password",
			plan: map[int](map[int]actions.Func){
				1: {
					1: {
						Name:      actions.AddUser,
						Reference: actions.AddUser,
						Argument: []interface{}{
							actions.ADU{
								Username: "username",
								Password: "password",
							},
						},
					},
				},
			},
			actions: &mock.Actions{
				AddUserFn: func(interface{}) (rpi.Exec, error) {
					return rpi.Exec{
						Name:       actions.AddUser,
						StartTime:  1,
						EndTime:    2,
						ExitStatus: 0,
						Stdout:     "username-password-add",
					}, nil
				},
			},
			consys: &mocksys.Action{
				ExecuteAUSFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{
						Name:          actions.AddUser,
						NumberOfSteps: 1,
						Progress: map[string]rpi.Exec{
							"1": {
								Name:       actions.AddUser,
								StartTime:  1,
								EndTime:    2,
								ExitStatus: 0,
								Stdout:     "username-password-add",
							},
						},
						ExitStatus: 0,
						StartTime:  2,
						EndTime:    uint64(time.Now().Unix()),
					}, nil
				},
			},
			wantedData: rpi.Action{
				Name:          actions.AddUser,
				NumberOfSteps: 1,
				Progress: map[string]rpi.Exec{
					"1": {
						Name:       actions.AddUser,
						StartTime:  1,
						EndTime:    2,
						ExitStatus: 0,
						Stdout:     "username-password-add",
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
			s := configure.New(tc.consys, tc.actions, tc.infos)
			user, err := s.ExecuteAUS(tc.username, tc.password)
			assert.Equal(t, tc.wantedData, user)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestExecuteDUS(t *testing.T) {
	cases := []struct {
		name       string
		username   string
		plan       map[int](map[int]actions.Func)
		actions    *mock.Actions
		infos      *mock.Infos
		consys     *mocksys.Action
		wantedData rpi.Action
		wantedErr  error
	}{
		{
			name:     "success",
			username: "username",
			plan: map[int](map[int]actions.Func){
				1: {
					1: {
						Name:      actions.DeleteUser,
						Reference: actions.DeleteUser,
						Argument: []interface{}{
							actions.ADU{
								Username: "username",
							},
						},
					},
				},
			},
			actions: &mock.Actions{
				DeleteUserFn: func(interface{}) (rpi.Exec, error) {
					return rpi.Exec{
						Name:       actions.DeleteUser,
						StartTime:  1,
						EndTime:    2,
						ExitStatus: 0,
						Stdout:     "username-password-add",
					}, nil
				},
			},
			consys: &mocksys.Action{
				ExecuteDUSFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{
						Name:          actions.DeleteUser,
						NumberOfSteps: 1,
						Progress: map[string]rpi.Exec{
							"1": {
								Name:       actions.DeleteUser,
								StartTime:  1,
								EndTime:    2,
								ExitStatus: 0,
								Stdout:     "username-password-add",
							},
						},
						ExitStatus: 0,
						StartTime:  2,
						EndTime:    uint64(time.Now().Unix()),
					}, nil
				},
			},
			wantedData: rpi.Action{
				Name:          actions.DeleteUser,
				NumberOfSteps: 1,
				Progress: map[string]rpi.Exec{
					"1": {
						Name:       actions.DeleteUser,
						StartTime:  1,
						EndTime:    2,
						ExitStatus: 0,
						Stdout:     "username-password-add",
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
			s := configure.New(tc.consys, tc.actions, tc.infos)
			user, err := s.ExecuteDUS(tc.username)
			assert.Equal(t, tc.wantedData, user)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestExecuteCA(t *testing.T) {
	cases := []struct {
		name       string
		action     string
		plan       map[int](map[int]actions.Func)
		actions    *mock.Actions
		infos      *mock.Infos
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
						Name:      actions.DisableOrEnableCameraInterface,
						Reference: actions.DisableOrEnableConfig,
						Argument: []interface{}{
							actions.EODC{
								DirOrFilePath: "path",
								Action:        "enable",
							},
						},
					},
				},
			},
			wantedData: rpi.Action{},
			wantedErr:  echo.NewHTTPError(http.StatusInternalServerError, "bad action type: enable or disable camera failed"),
		},
		{
			name:   "success",
			action: "enable",
			plan: map[int](map[int]actions.Func){
				1: {
					1: {
						Name:      actions.DisableOrEnableCameraInterface,
						Reference: actions.DisableOrEnableConfig,
						Argument: []interface{}{
							actions.EODC{
								DirOrFilePath: "path",
								Action:        "enable",
							},
						},
					},
				},
			},
			infos: &mock.Infos{
				GetConfigFilesFn: func() map[string]rpi.ConfigFileDetails {
					return map[string]rpi.ConfigFileDetails{
						"bootconfig": {
							Path: "/dummy/path",
						},
					}
				},
			},
			actions: &mock.Actions{
				DisableOrEnableConfigFn: func(interface{}) (rpi.Exec, error) {
					return rpi.Exec{
						Name:       actions.DisableOrEnableConfig,
						StartTime:  1,
						EndTime:    2,
						ExitStatus: 0,
						Stdout:     "path-enable",
					}, nil
				},
			},
			consys: &mocksys.Action{
				ExecuteCAFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{
						Name:          actions.CameraInterface,
						NumberOfSteps: 1,
						Progress: map[string]rpi.Exec{
							"1": {
								Name:       actions.DisableOrEnableConfig,
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
				Name:          actions.CameraInterface,
				NumberOfSteps: 1,
				Progress: map[string]rpi.Exec{
					"1": {
						Name:       actions.DisableOrEnableConfig,
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
						Name:      actions.DisableOrEnableCameraRegex,
						Reference: actions.DisableOrEnableConfig,
						Argument: []interface{}{
							actions.EODC{
								DirOrFilePath: "path",
								Action:        "disable",
							},
						},
					},
				},
			},
			infos: &mock.Infos{
				GetConfigFilesFn: func() map[string]rpi.ConfigFileDetails {
					return map[string]rpi.ConfigFileDetails{
						"bootconfig": {
							Path: "/dummy/path",
						},
					}
				},
			},
			actions: &mock.Actions{
				DisableOrEnableConfigFn: func(interface{}) (rpi.Exec, error) {
					return rpi.Exec{
						Name:       actions.DisableOrEnableConfig,
						StartTime:  1,
						EndTime:    2,
						ExitStatus: 0,
						Stdout:     "path-disable",
					}, nil
				},
			},
			consys: &mocksys.Action{
				ExecuteCAFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{
						Name:          actions.CameraInterface,
						NumberOfSteps: 1,
						Progress: map[string]rpi.Exec{
							"1": {
								Name:       actions.DisableOrEnableConfig,
								StartTime:  1,
								EndTime:    2,
								ExitStatus: 0,
								Stdout:     "path-disable",
							},
						},
						ExitStatus: 0,
						StartTime:  2,
						EndTime:    uint64(time.Now().Unix()),
					}, nil
				},
			},
			wantedData: rpi.Action{
				Name:          actions.CameraInterface,
				NumberOfSteps: 1,
				Progress: map[string]rpi.Exec{
					"1": {
						Name:       actions.DisableOrEnableConfig,
						StartTime:  1,
						EndTime:    2,
						ExitStatus: 0,
						Stdout:     "path-disable",
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
			s := configure.New(tc.consys, tc.actions, tc.infos)
			camera, err := s.ExecuteCA(tc.action)
			assert.Equal(t, tc.wantedData, camera)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestExecuteSSH(t *testing.T) {
	cases := []struct {
		name       string
		action     string
		plan       map[int](map[int]actions.Func)
		actions    *mock.Actions
		infos      *mock.Infos
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
						Name:      actions.ExecuteBashCommand,
						Reference: actions.ExecuteBashCommand,
						Argument: []interface{}{
							actions.EBC{
								Command: "command",
							},
						},
					},
				},
			},
			wantedData: rpi.Action{},
			wantedErr:  echo.NewHTTPError(http.StatusInternalServerError, "bad action type: enable or disable ssh failed"),
		},
		{
			name:   "success",
			action: "enable",
			plan: map[int](map[int]actions.Func){
				1: {
					1: {
						Name:      actions.ExecuteBashCommand,
						Reference: actions.ExecuteBashCommand,
						Argument: []interface{}{
							actions.EBC{Command: "command"},
						},
					},
				},
			},
			actions: &mock.Actions{
				ExecuteBashCommandFn: func(interface{}) (rpi.Exec, error) {
					return rpi.Exec{
						Name:       actions.SSH,
						StartTime:  1,
						EndTime:    2,
						ExitStatus: 0,
						Stdout:     "path-enable",
					}, nil
				},
			},
			consys: &mocksys.Action{
				ExecuteSSHFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{
						Name:          actions.SSH,
						NumberOfSteps: 1,
						Progress: map[string]rpi.Exec{
							"1": {
								Name:       actions.ExecuteBashCommand,
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
				Name:          actions.SSH,
				NumberOfSteps: 1,
				Progress: map[string]rpi.Exec{
					"1": {
						Name:       actions.ExecuteBashCommand,
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
						Name:      actions.ExecuteBashCommand,
						Reference: actions.ExecuteBashCommand,
						Argument: []interface{}{
							actions.EBC{
								Command: "command",
							},
						},
					},
				},
			},
			actions: &mock.Actions{
				ExecuteBashCommandFn: func(interface{}) (rpi.Exec, error) {
					return rpi.Exec{
						Name:       actions.SSH,
						StartTime:  1,
						EndTime:    2,
						ExitStatus: 0,
						Stdout:     "path-enable",
					}, nil
				},
			},
			infos: &mock.Infos{
				GetConfigFilesFn: func() map[string]rpi.ConfigFileDetails {
					return map[string]rpi.ConfigFileDetails{
						"bootconfig": {
							Path: "/dummy/path",
						},
					}
				},
			},
			consys: &mocksys.Action{
				ExecuteSSHFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
					return rpi.Action{
						Name:          actions.SSH,
						NumberOfSteps: 1,
						Progress: map[string]rpi.Exec{
							"1": {
								Name:       actions.ExecuteBashCommand,
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
				Name:          actions.SSH,
				NumberOfSteps: 1,
				Progress: map[string]rpi.Exec{
					"1": {
						Name:       actions.ExecuteBashCommand,
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
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := configure.New(tc.consys, tc.actions, tc.infos)
			camera, err := s.ExecuteSSH(tc.action)
			assert.Equal(t, tc.wantedData, camera)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
