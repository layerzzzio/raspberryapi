package sys

import (
	"testing"

	"github.com/raspibuddy/rpi/pkg/api/actions/general"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
	"github.com/raspibuddy/rpi/pkg/utl/test_utl"

	"github.com/stretchr/testify/assert"
)

func TestExecuteRBS(t *testing.T) {
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
			name: "success",
			plan: map[int](map[int]actions.Func){
				1: {
					1: {
						Name:      actions.Reboot,
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
			wantedDataName:        "RebootShutdown",
			wantedDataNumSteps:    1,
			wantedDataStdOutStep1: "string0-string1",
			wantedDataExitStatus:  0,
			wantedErr:             nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := general.GENSYS(General{})
			rebootShutdown, err := s.ExecuteRBS(tc.plan)

			assert.Equal(t, tc.wantedDataName, rebootShutdown.Name)
			assert.Equal(t, tc.wantedDataNumSteps, rebootShutdown.NumberOfSteps)
			assert.Equal(t, tc.wantedDataStdOutStep1, rebootShutdown.Progress["1<|>1"].Stdout)
			assert.Equal(t, tc.wantedDataExitStatus, rebootShutdown.ExitStatus)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestExecuteSASO(t *testing.T) {
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
			name: "success",
			plan: map[int](map[int]actions.Func){
				1: {
					1: {
						Name:      actions.Reboot,
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
			wantedDataName:        "StartStop",
			wantedDataNumSteps:    1,
			wantedDataStdOutStep1: "string0-string1",
			wantedDataExitStatus:  0,
			wantedErr:             nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := general.GENSYS(General{})
			rebootShutdown, err := s.ExecuteSASO(tc.plan)

			assert.Equal(t, tc.wantedDataName, rebootShutdown.Name)
			assert.Equal(t, tc.wantedDataNumSteps, rebootShutdown.NumberOfSteps)
			assert.Equal(t, tc.wantedDataStdOutStep1, rebootShutdown.Progress["1<|>1"].Stdout)
			assert.Equal(t, tc.wantedDataExitStatus, rebootShutdown.ExitStatus)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
