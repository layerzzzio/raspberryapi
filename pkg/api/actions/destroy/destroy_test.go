package destroy_test

import (
	"testing"
	"time"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/actions/destroy"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
	"github.com/raspibuddy/rpi/pkg/utl/mock"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/stretchr/testify/assert"
)

func TestExecuteDF(t *testing.T) {
	cases := []struct {
		name       string
		path       string
		execs      map[int]rpi.Exec
		actions    *mock.Actions
		dessys     *mocksys.Action
		wantedData rpi.Action
		wantedErr  error
	}{
		{
			name: "success",
			path: "/dummy",
			execs: map[int]rpi.Exec{
				1: {
					Name:       actions.DeleteFile,
					StartTime:  2,
					EndTime:    3,
					ExitStatus: 0,
				},
			},
			actions: &mock.Actions{
				DeleteFileFn: func(path string) rpi.Exec {
					return rpi.Exec{
						Name:       actions.DeleteFile,
						StartTime:  2,
						EndTime:    3,
						ExitStatus: 0,
					}
				},
			},
			dessys: &mocksys.Action{
				ExecuteDFFn: func(map[int]rpi.Exec) (rpi.Action, error) {
					return rpi.Action{
						Name: actions.DeleteFile,
						Steps: map[int]string{
							1: actions.DeleteFile,
						},
						NumberOfSteps: 1,
						Executions: map[int]rpi.Exec{
							1: {
								Name:       actions.DeleteFile,
								StartTime:  2,
								EndTime:    3,
								ExitStatus: 0,
							},
						},
						ExitStatus: 0,
						StartTime:  2,
						EndTime:    uint64(time.Now().Unix()),
					}, nil
				},
			},
			wantedData: rpi.Action{
				Name: actions.DeleteFile,
				Steps: map[int]string{
					1: actions.DeleteFile,
				},
				NumberOfSteps: 1,
				Executions: map[int]rpi.Exec{
					1: {
						Name:       actions.DeleteFile,
						StartTime:  2,
						EndTime:    3,
						ExitStatus: 0,
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

func TestExecuteDU(t *testing.T) {
	cases := []struct {
		name        string
		processname string
		processtype string
		execs       map[int]rpi.Exec
		actions     *mock.Actions
		dessys      *mocksys.Action
		wantedData  rpi.Action
		wantedErr   error
	}{
		{
			name:        "success",
			processname: "dummyprocess",
			processtype: "dummytype",
			execs: map[int]rpi.Exec{
				1: {
					Name:       actions.KillProcessByName,
					StartTime:  2,
					EndTime:    3,
					ExitStatus: 0,
				},
			},
			actions: &mock.Actions{
				KillProcessByNameFn: func(processname string, processtype string) rpi.Exec {
					return rpi.Exec{
						Name:       actions.KillProcessByName,
						StartTime:  2,
						EndTime:    3,
						ExitStatus: 0,
					}
				},
			},
			dessys: &mocksys.Action{
				ExecuteDUFn: func(map[int]rpi.Exec) (rpi.Action, error) {
					return rpi.Action{
						Name: actions.DisconnectUser,
						Steps: map[int]string{
							1: actions.KillProcessByName,
						},
						NumberOfSteps: 1,
						Executions: map[int]rpi.Exec{
							1: {
								Name:       actions.KillProcessByName,
								StartTime:  2,
								EndTime:    3,
								ExitStatus: 0,
							},
						},
						ExitStatus: 0,
						StartTime:  2,
						EndTime:    uint64(time.Now().Unix()),
					}, nil
				},
			},
			wantedData: rpi.Action{
				Name: actions.DisconnectUser,
				Steps: map[int]string{
					1: actions.KillProcessByName,
				},
				NumberOfSteps: 1,
				Executions: map[int]rpi.Exec{
					1: {
						Name:       actions.KillProcessByName,
						StartTime:  2,
						EndTime:    3,
						ExitStatus: 0,
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
			deletefile, err := s.ExecuteDU(tc.processname, tc.processtype)
			assert.Equal(t, tc.wantedData, deletefile)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestExecuteKP(t *testing.T) {
	cases := []struct {
		name       string
		pid        int
		execs      map[int]rpi.Exec
		actions    *mock.Actions
		dessys     *mocksys.Action
		wantedData rpi.Action
		wantedErr  error
	}{
		{
			name: "success",
			pid:  123,
			execs: map[int]rpi.Exec{
				1: {
					Name:       actions.KillProcess,
					StartTime:  2,
					EndTime:    3,
					ExitStatus: 0,
				},
			},
			actions: &mock.Actions{
				KillProcessFn: func(pid string) rpi.Exec {
					return rpi.Exec{
						Name:       actions.KillProcess,
						StartTime:  2,
						EndTime:    3,
						ExitStatus: 0,
					}
				},
			},
			dessys: &mocksys.Action{
				ExecuteKPFn: func(map[int]rpi.Exec) (rpi.Action, error) {
					return rpi.Action{
						Name: actions.KillProcess,
						Steps: map[int]string{
							1: actions.KillProcess,
						},
						NumberOfSteps: 1,
						Executions: map[int]rpi.Exec{
							1: {
								Name:       actions.KillProcess,
								StartTime:  2,
								EndTime:    3,
								ExitStatus: 0,
							},
						},
						ExitStatus: 0,
						StartTime:  2,
						EndTime:    uint64(time.Now().Unix()),
					}, nil
				},
			},
			wantedData: rpi.Action{
				Name: actions.KillProcess,
				Steps: map[int]string{
					1: actions.KillProcess,
				},
				NumberOfSteps: 1,
				Executions: map[int]rpi.Exec{
					1: {
						Name:       actions.KillProcess,
						StartTime:  2,
						EndTime:    3,
						ExitStatus: 0,
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
