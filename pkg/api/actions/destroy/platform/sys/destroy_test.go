package sys

import (
	"net/http"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/actions/destroy"
	"github.com/raspibuddy/rpi/pkg/utl/actions"

	"github.com/stretchr/testify/assert"
)

func TestExecuteDF(t *testing.T) {
	cases := []struct {
		name       string
		execs      map[int]rpi.Exec
		wantedData rpi.Action
		wantedErr  error
	}{
		{
			name:       "length steps & execs are different",
			execs:      map[int]rpi.Exec{},
			wantedData: rpi.Action{},
			wantedErr:  echo.NewHTTPError(http.StatusInternalServerError, "length steps & execs are different"),
		},
		{
			name: "length steps & execs are different",
			execs: map[int]rpi.Exec{
				0: {
					Name:       actions.DeleteFile,
					StartTime:  2,
					EndTime:    3,
					ExitStatus: 0,
				},
			},
			wantedData: rpi.Action{},
			wantedErr:  echo.NewHTTPError(http.StatusInternalServerError, "first key in execs map is not equal 1"),
		},
		{
			name: "success",
			execs: map[int]rpi.Exec{
				1: {
					Name:       actions.DeleteFile,
					StartTime:  2,
					EndTime:    3,
					ExitStatus: 0,
				},
			},
			wantedData: rpi.Action{
				Name:          actions.DeleteFile,
				Steps:         map[int]string{1: actions.DeleteFile},
				NumberOfSteps: 1,
				Executions: map[int]rpi.Exec{
					1: {
						Name:       actions.DeleteFile,
						StartTime:  2,
						EndTime:    3,
						ExitStatus: 0,
					}},
				ExitStatus: 0,
				StartTime:  2,
				EndTime:    uint64(time.Now().Unix()),
			},
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := destroy.DESSYS(Destroy{})
			deletefile, err := s.ExecuteDF(tc.execs)
			assert.Equal(t, tc.wantedData, deletefile)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
