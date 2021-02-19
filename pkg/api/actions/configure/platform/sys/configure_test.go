package sys

import (
	"testing"

	"github.com/raspibuddy/rpi/pkg/api/actions/configure"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
	"github.com/raspibuddy/rpi/pkg/utl/test_utl"

	"github.com/stretchr/testify/assert"
)

func TestExecuteCH(t *testing.T) {
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
						Name:      actions.ChangeHostname,
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
			wantedDataName:        "change_hostname",
			wantedDataNumSteps:    1,
			wantedDataStdOutStep1: "string0-string1",
			wantedDataExitStatus:  0,
			wantedErr:             nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := configure.CONSYS(Configure{})
			changeHostname, err := s.ExecuteCH(tc.plan)
			assert.Equal(t, tc.wantedDataName, changeHostname.Name)
			assert.Equal(t, tc.wantedDataNumSteps, changeHostname.NumberOfSteps)
			assert.Equal(t, tc.wantedDataStdOutStep1, changeHostname.Progress["1<|>1"].Stdout)
			assert.Equal(t, tc.wantedDataExitStatus, changeHostname.ExitStatus)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestExecuteCP(t *testing.T) {
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
						Name:      actions.ChangePassword,
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
			wantedDataName:        "change_password",
			wantedDataNumSteps:    1,
			wantedDataStdOutStep1: "string0-string1",
			wantedDataExitStatus:  0,
			wantedErr:             nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := configure.CONSYS(Configure{})
			changePassword, err := s.ExecuteCP(tc.plan)
			assert.Equal(t, tc.wantedDataName, changePassword.Name)
			assert.Equal(t, tc.wantedDataNumSteps, changePassword.NumberOfSteps)
			assert.Equal(t, tc.wantedDataStdOutStep1, changePassword.Progress["1<|>1"].Stdout)
			assert.Equal(t, tc.wantedDataExitStatus, changePassword.ExitStatus)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestExecuteWNB(t *testing.T) {
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
						Name:      actions.WaitForNetworkAtBoot,
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
			wantedDataName:        "wait_for_network_at_boot",
			wantedDataNumSteps:    1,
			wantedDataStdOutStep1: "string0-string1",
			wantedDataExitStatus:  0,
			wantedErr:             nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := configure.CONSYS(Configure{})
			changePassword, err := s.ExecuteWNB(tc.plan)
			assert.Equal(t, tc.wantedDataName, changePassword.Name)
			assert.Equal(t, tc.wantedDataNumSteps, changePassword.NumberOfSteps)
			assert.Equal(t, tc.wantedDataStdOutStep1, changePassword.Progress["1<|>1"].Stdout)
			assert.Equal(t, tc.wantedDataExitStatus, changePassword.ExitStatus)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
