package configure_test

import (
	"testing"
	"time"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/actions/configure"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
	"github.com/raspibuddy/rpi/pkg/utl/mock"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/raspibuddy/rpi/pkg/utl/test_utl"
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
						Name:      "FuncA",
						Reference: test_utl.FuncA,
						Argument: []interface{}{
							test_utl.ArgFuncA{
								Arg0: "string0",
								Arg1: "string1",
							},
						},
					},
					2: {
						Name:      "FuncB",
						Reference: test_utl.FuncB,
						Argument: []interface{}{
							test_utl.ArgFuncB{
								Arg2: "string2",
							},
						},
					},
				},
			},
			actions: &mock.Actions{
				ChangeHostnameInHostnameFileFn: func(interface{}) (rpi.Exec, error) {
					return rpi.Exec{
						Name:       "FuncA",
						StartTime:  1,
						EndTime:    2,
						ExitStatus: 0,
						Stdout:     "string0-string1",
					}, nil
				},
				ChangeHostnameInHostsFileFn: func(interface{}) (rpi.Exec, error) {
					return rpi.Exec{
						Name:       "FuncB",
						StartTime:  1,
						EndTime:    2,
						ExitStatus: 0,
						Stdout:     "string2",
					}, nil
				},
			},
			consys: &mocksys.Action{
				ExecuteCHFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
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
							"2": {
								Name:       "FuncB",
								StartTime:  1,
								EndTime:    2,
								ExitStatus: 0,
								Stdout:     "string2",
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
					"2": {
						Name:       "FuncB",
						StartTime:  1,
						EndTime:    2,
						ExitStatus: 0,
						Stdout:     "string2",
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
				ChangePasswordFn: func(interface{}) (rpi.Exec, error) {
					return rpi.Exec{
						Name:       "FuncA",
						StartTime:  1,
						EndTime:    2,
						ExitStatus: 0,
						Stdout:     "string0-string1",
					}, nil
				},
			},
			consys: &mocksys.Action{
				ExecuteCPFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
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
			s := configure.New(tc.consys, tc.actions)
			changePassword, err := s.ExecuteCP(tc.password, tc.username)
			assert.Equal(t, tc.wantedData, changePassword)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
