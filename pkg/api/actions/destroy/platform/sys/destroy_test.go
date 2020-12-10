package sys

import (
	"testing"

	"github.com/raspibuddy/rpi/pkg/api/actions/destroy"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
	"github.com/raspibuddy/rpi/pkg/utl/test_utl"

	"github.com/stretchr/testify/assert"
)

func TestExecuteDF(t *testing.T) {
	cases := []struct {
		name                  string
		plan                  map[int](map[int]actions.Func)
		wantedDataName        string
		wantedDataNumSteps    uint16
		wantedDataStdOutStep1 string
		wantedDataExitStatus  uint8
		wantedErr             error
	}{
		{
			name: "success : action two steps action failed",
			plan: map[int](map[int]actions.Func){
				1: {
					1: {
						Name:      actions.KillProcessByName,
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
			wantedDataName:        "delete_file",
			wantedDataNumSteps:    1,
			wantedDataStdOutStep1: "string0-string1",
			wantedDataExitStatus:  0,
			wantedErr:             nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := destroy.DESSYS(Destroy{})
			deletefile, err := s.ExecuteDF(tc.plan)

			assert.Equal(t, tc.wantedDataName, deletefile.Name)
			assert.Equal(t, tc.wantedDataNumSteps, deletefile.NumberOfSteps)
			assert.Equal(t, tc.wantedDataStdOutStep1, deletefile.Progress["1<|>1"].Stdout)
			assert.Equal(t, tc.wantedDataExitStatus, deletefile.ExitStatus)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestExecuteSUS(t *testing.T) {
	cases := []struct {
		name                  string
		plan                  map[int](map[int]actions.Func)
		wantedDataName        string
		wantedDataNumSteps    uint16
		wantedDataStdOutStep1 string
		wantedDataExitStatus  uint8
		wantedErr             error
	}{
		{
			name: "success : action two steps action failed",
			plan: map[int](map[int]actions.Func){
				1: {
					1: {
						Name:      actions.StopUserSession,
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
			wantedDataName:        "stop_user_sessions",
			wantedDataNumSteps:    1,
			wantedDataStdOutStep1: "string0-string1",
			wantedDataExitStatus:  0,
			wantedErr:             nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := destroy.DESSYS(Destroy{})
			deletefile, err := s.ExecuteSUS(tc.plan)

			assert.Equal(t, tc.wantedDataName, deletefile.Name)
			assert.Equal(t, tc.wantedDataNumSteps, deletefile.NumberOfSteps)
			assert.Equal(t, tc.wantedDataStdOutStep1, deletefile.Progress["1<|>1"].Stdout)
			assert.Equal(t, tc.wantedDataExitStatus, deletefile.ExitStatus)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestExecuteKP(t *testing.T) {
	cases := []struct {
		name                  string
		plan                  map[int](map[int]actions.Func)
		wantedDataName        string
		wantedDataNumSteps    uint16
		wantedDataStdOutStep1 string
		wantedDataExitStatus  uint8
		wantedErr             error
	}{
		{
			name: "success : action two steps action failed",
			plan: map[int](map[int]actions.Func){
				1: {
					1: {
						Name:      actions.KillProcess,
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
			wantedDataName:        "kill_process",
			wantedDataNumSteps:    1,
			wantedDataStdOutStep1: "string0-string1",
			wantedDataExitStatus:  0,
			wantedErr:             nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := destroy.DESSYS(Destroy{})
			deletefile, err := s.ExecuteKP(tc.plan)

			assert.Equal(t, tc.wantedDataName, deletefile.Name)
			assert.Equal(t, tc.wantedDataNumSteps, deletefile.NumberOfSteps)
			assert.Equal(t, tc.wantedDataStdOutStep1, deletefile.Progress["1<|>1"].Stdout)
			assert.Equal(t, tc.wantedDataExitStatus, deletefile.ExitStatus)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
