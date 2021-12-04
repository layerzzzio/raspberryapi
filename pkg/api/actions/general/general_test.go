package general_test

import (
	"testing"
	"time"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/actions/general"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
	"github.com/raspibuddy/rpi/pkg/utl/mock"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/raspibuddy/rpi/pkg/utl/test_utl"
	"github.com/stretchr/testify/assert"
)

func TestExecuteRBS(t *testing.T) {
	cases := []struct {
		name       string
		option     string
		plan       map[int](map[int]actions.Func)
		actions    *mock.Actions
		gensys     *mocksys.Action
		wantedData rpi.Action
		wantedErr  error
	}{
		{
			name:   "success",
			option: "reboot",
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
			},
			gensys: &mocksys.Action{
				ExecuteRBSFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
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
			s := general.New(tc.gensys, tc.actions)
			rebooSshutdown, err := s.ExecuteRBS(tc.option)
			assert.Equal(t, tc.wantedData, rebooSshutdown)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestExecuteSASO(t *testing.T) {
	cases := []struct {
		name       string
		action     string
		service    string
		plan       map[int](map[int]actions.Func)
		actions    *mock.Actions
		gensys     *mocksys.Action
		wantedData rpi.Action
		wantedErr  error
	}{
		{
			name:    "success",
			action:  "start",
			service: "raspibuddy_deploy",
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
			},
			gensys: &mocksys.Action{
				ExecuteSASOFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
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
			s := general.New(tc.gensys, tc.actions)
			startStop, err := s.ExecuteSASO(tc.action, tc.service)
			assert.Equal(t, tc.wantedData, startStop)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
