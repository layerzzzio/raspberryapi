package destroy_test

import (
	"testing"
	"time"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/actions/destroy"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
	"github.com/raspibuddy/rpi/pkg/utl/mock"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/raspibuddy/rpi/pkg/utl/test_utl"
	"github.com/stretchr/testify/assert"
)

func TestExecuteDF(t *testing.T) {
	cases := []struct {
		name       string
		path       string
		plan       map[int](map[int]actions.Func)
		actions    *mock.Actions
		dessys     *mocksys.Action
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
				},
			},
			actions: &mock.Actions{
				DeleteFileFn: func(interface{}) (rpi.Exec, error) {
					return rpi.Exec{
						Name:       "FuncA",
						StartTime:  1,
						EndTime:    2,
						ExitStatus: 0,
						Stdout:     "string0-string1",
					}, nil
				},
			},
			dessys: &mocksys.Action{
				ExecuteDFFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
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
			s := destroy.New(tc.dessys, tc.actions)
			deletefile, err := s.ExecuteDF(tc.path)
			assert.Equal(t, tc.wantedData, deletefile)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestExecuteSUS(t *testing.T) {
	cases := []struct {
		name        string
		processname string
		plan        map[int](map[int]actions.Func)
		actions     *mock.Actions
		dessys      *mocksys.Action
		wantedData  rpi.Action
		wantedErr   error
	}{
		{
			name:        "success",
			processname: "pts/2",
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
				KillProcessByNameFn: func(interface{}) (rpi.Exec, error) {
					return rpi.Exec{
						Name:       "FuncA",
						StartTime:  1,
						EndTime:    2,
						ExitStatus: 0,
						Stdout:     "string0-string1",
					}, nil
				},
			},
			dessys: &mocksys.Action{
				ExecuteSUSFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
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
			s := destroy.New(tc.dessys, tc.actions)
			deletefile, err := s.ExecuteSUS(tc.processname, "terminal")
			assert.Equal(t, tc.wantedData, deletefile)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestExecuteKP(t *testing.T) {
	cases := []struct {
		name       string
		pid        int
		plan       map[int](map[int]actions.Func)
		actions    *mock.Actions
		dessys     *mocksys.Action
		wantedData rpi.Action
		wantedErr  error
	}{
		{
			name: "success",
			pid:  12345,
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
				KillProcessFn: func(interface{}) (rpi.Exec, error) {
					return rpi.Exec{
						Name:       "FuncA",
						StartTime:  1,
						EndTime:    2,
						ExitStatus: 0,
						Stdout:     "string0-string1",
					}, nil
				},
			},
			dessys: &mocksys.Action{
				ExecuteKPFn: func(map[int](map[int]actions.Func)) (rpi.Action, error) {
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
			s := destroy.New(tc.dessys, tc.actions)
			deletefile, err := s.ExecuteKP(tc.pid)
			assert.Equal(t, tc.wantedData, deletefile)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

// func TestExecuteSUS(t *testing.T) {
// 	cases := []struct {
// 		name        string
// 		processname string
// 		processtype string
// 		execs       map[int]rpi.Exec
// 		actions     *mock.Actions
// 		dessys      *mocksys.Action
// 		wantedData  rpi.Action
// 		wantedErr   error
// 	}{
// 		{
// 			name:        "success",
// 			processname: "dummyprocess",
// 			processtype: "dummytype",
// 			execs: map[int]rpi.Exec{
// 				1: {
// 					Name:       actions.KillProcessByName,
// 					StartTime:  2,
// 					EndTime:    3,
// 					ExitStatus: 0,
// 				},
// 			},
// 			actions: &mock.Actions{
// 				KillProcessByNameFn: func(processname string, processtype string) rpi.Exec {
// 					return rpi.Exec{
// 						Name:       actions.KillProcessByName,
// 						StartTime:  2,
// 						EndTime:    3,
// 						ExitStatus: 0,
// 					}
// 				},
// 			},
// 			dessys: &mocksys.Action{
// 				ExecuteSUSFn: func(map[int]rpi.Exec) (rpi.Action, error) {
// 					return rpi.Action{
// 						Name: actions.StopUserSession,
// 						Steps: map[int]string{
// 							1: actions.KillProcessByName,
// 						},
// 						NumberOfSteps: 1,
// 						Executions: map[int]rpi.Exec{
// 							1: {
// 								Name:       actions.KillProcessByName,
// 								StartTime:  2,
// 								EndTime:    3,
// 								ExitStatus: 0,
// 							},
// 						},
// 						ExitStatus: 0,
// 						StartTime:  2,
// 						EndTime:    uint64(time.Now().Unix()),
// 					}, nil
// 				},
// 			},
// 			wantedData: rpi.Action{
// 				Name: actions.StopUserSession,
// 				Steps: map[int]string{
// 					1: actions.KillProcessByName,
// 				},
// 				NumberOfSteps: 1,
// 				Executions: map[int]rpi.Exec{
// 					1: {
// 						Name:       actions.KillProcessByName,
// 						StartTime:  2,
// 						EndTime:    3,
// 						ExitStatus: 0,
// 					},
// 				},
// 				ExitStatus: 0,
// 				StartTime:  2,
// 				EndTime:    uint64(time.Now().Unix()),
// 			},
// 			wantedErr: nil,
// 		},
// 	}

// 	for _, tc := range cases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			s := destroy.New(tc.dessys, tc.actions)
// 			deletefile, err := s.ExecuteSUS(tc.processname, tc.processtype)
// 			assert.Equal(t, tc.wantedData, deletefile)
// 			assert.Equal(t, tc.wantedErr, err)
// 		})
// 	}
// }

// func TestExecuteKP(t *testing.T) {
// 	cases := []struct {
// 		name       string
// 		pid        int
// 		execs      map[int]rpi.Exec
// 		actions    *mock.Actions
// 		dessys     *mocksys.Action
// 		wantedData rpi.Action
// 		wantedErr  error
// 	}{
// 		{
// 			name: "success",
// 			pid:  123,
// 			execs: map[int]rpi.Exec{
// 				1: {
// 					Name:       actions.KillProcess,
// 					StartTime:  2,
// 					EndTime:    3,
// 					ExitStatus: 0,
// 				},
// 			},
// 			actions: &mock.Actions{
// 				KillProcessFn: func(pid string) rpi.Exec {
// 					return rpi.Exec{
// 						Name:       actions.KillProcess,
// 						StartTime:  2,
// 						EndTime:    3,
// 						ExitStatus: 0,
// 					}
// 				},
// 			},
// 			dessys: &mocksys.Action{
// 				ExecuteKPFn: func(map[int]rpi.Exec) (rpi.Action, error) {
// 					return rpi.Action{
// 						Name: actions.KillProcess,
// 						Steps: map[int]string{
// 							1: actions.KillProcess,
// 						},
// 						NumberOfSteps: 1,
// 						Executions: map[int]rpi.Exec{
// 							1: {
// 								Name:       actions.KillProcess,
// 								StartTime:  2,
// 								EndTime:    3,
// 								ExitStatus: 0,
// 							},
// 						},
// 						ExitStatus: 0,
// 						StartTime:  2,
// 						EndTime:    uint64(time.Now().Unix()),
// 					}, nil
// 				},
// 			},
// 			wantedData: rpi.Action{
// 				Name: actions.KillProcess,
// 				Steps: map[int]string{
// 					1: actions.KillProcess,
// 				},
// 				NumberOfSteps: 1,
// 				Executions: map[int]rpi.Exec{
// 					1: {
// 						Name:       actions.KillProcess,
// 						StartTime:  2,
// 						EndTime:    3,
// 						ExitStatus: 0,
// 					},
// 				},
// 				ExitStatus: 0,
// 				StartTime:  2,
// 				EndTime:    uint64(time.Now().Unix()),
// 			},
// 			wantedErr: nil,
// 		},
// 	}

// 	for _, tc := range cases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			s := destroy.New(tc.dessys, tc.actions)
// 			deletefile, err := s.ExecuteKP(tc.pid)
// 			assert.Equal(t, tc.wantedData, deletefile)
// 			assert.Equal(t, tc.wantedErr, err)
// 		})
// 	}
// }
