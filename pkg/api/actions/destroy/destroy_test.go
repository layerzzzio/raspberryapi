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
