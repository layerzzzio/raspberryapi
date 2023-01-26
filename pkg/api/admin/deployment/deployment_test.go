package deployment_test

import (
	"testing"
	"time"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/admin/deployment"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
	"github.com/raspibuddy/rpi/pkg/utl/mock"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/raspibuddy/rpi/pkg/utl/test_utl"
	"github.com/stretchr/testify/assert"
)

func TestExecuteDPTOOL(t *testing.T) {
	cases := []struct {
		name       string
		deployType string
		url        string
		version    string
		plan       map[int](map[int]actions.Func)
		actions    *mock.Actions
		dsys       *mocksys.Action
		wantedData rpi.Action
		wantedErr  error
	}{
		{
			name:       "success: regular version",
			deployType: "full_deploy",
			url:        "https://domain/",
			version:    "1.0.0",
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
			dsys: &mocksys.Action{
				ExecuteDPTOOLFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
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
			s := deployment.New(tc.dsys, tc.actions)
			dp, err := s.ExecuteDPTOOL(tc.deployType, tc.url, tc.version)
			assert.Equal(t, tc.wantedData, dp)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
