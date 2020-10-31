package sys

import (
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/actions/deletefile"
	"github.com/raspibuddy/rpi/pkg/utl/actions"

	"github.com/stretchr/testify/assert"
)

func TestExecute(t *testing.T) {
	cases := []struct {
		name       string
		actionName string
		steps      map[int]string
		execs      []rpi.Exec
		startTime  uint64
		endTime    uint64
		wantedData rpi.Action
		wantedErr  error
	}{
		{
			name:       "success",
			actionName: actions.DeleteFile,
			steps:      map[int]string{1: actions.DeleteFile},
			execs: []rpi.Exec{
				{
					Name:       actions.DeleteFile,
					StartTime:  2,
					EndTime:    3,
					ExitStatus: 0,
				},
			},
			startTime: 1,
			endTime:   4,
			wantedData: rpi.Action{
				Name:          "delete_file",
				Steps:         map[int]string{1: actions.DeleteFile},
				NumberOfSteps: 1,
				Executions: []rpi.Exec{
					{
						Name:       actions.DeleteFile,
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
			s := deletefile.DELSYS(DeleteFile{})
			deletefile, err := s.Execute(tc.actionName, tc.steps, tc.execs, tc.startTime, tc.endTime)
			assert.Equal(t, tc.wantedData, deletefile)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
