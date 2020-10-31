package deletefile_test

import (
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/actions/deletefile"
	"github.com/raspibuddy/rpi/pkg/utl/actions"
	"github.com/raspibuddy/rpi/pkg/utl/mock"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/stretchr/testify/assert"
)

func TestExecute(t *testing.T) {
	cases := []struct {
		name       string
		actionName string
		path       string
		startTime  uint64
		endTime    uint64
		steps      map[int]string
		actions    *mock.Actions
		delsys     *mocksys.Action
		wantedData rpi.Action
		wantedErr  error
	}{
		{
			name:       "execs count different from steps count",
			actionName: actions.DeleteFile,
			path:       "/dummy",
			startTime:  1,
			endTime:    4,
			steps: map[int]string{
				1: "dummy_step_1",
				2: "dummy_step_2",
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
			delsys: &mocksys.Action{
				ExecuteFn: func(string, map[int]string, []rpi.Exec, uint64, uint64) (rpi.Action, error) {
					return rpi.Action{}, nil
				},
			},
			wantedData: rpi.Action{},
			wantedErr:  nil,
		},
		{
			name:       "deletefile returns 0",
			actionName: actions.DeleteFile,
			path:       "/dummy",
			startTime:  1,
			endTime:    4,
			steps: map[int]string{
				1: actions.DeleteFile,
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
			delsys: &mocksys.Action{
				ExecuteFn: func(string, map[int]string, []rpi.Exec, uint64, uint64) (rpi.Action, error) {
					return rpi.Action{
						Name:          "delete_file",
						Steps:         map[int]string{1: actions.DeleteFile},
						NumberOfSteps: 1,
						Executions: []rpi.Exec{
							{
								Name:       "delete_file",
								StartTime:  2,
								EndTime:    3,
								ExitStatus: 0,
							}},
						ExitStatus: 0,
						StartTime:  1,
						EndTime:    4,
					}, nil
				},
			},
			wantedData: rpi.Action{
				Name:          "delete_file",
				Steps:         map[int]string{1: actions.DeleteFile},
				NumberOfSteps: 1,
				Executions: []rpi.Exec{
					{
						Name:       "delete_file",
						StartTime:  2,
						EndTime:    3,
						ExitStatus: 0,
					}},
				ExitStatus: 0,
				StartTime:  1,
				EndTime:    4,
			},
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := deletefile.New(tc.delsys, tc.actions)
			deletefile, err := s.Execute(tc.path)
			assert.Equal(t, tc.wantedData, deletefile)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
