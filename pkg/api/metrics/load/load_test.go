package load_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/metrics/load"
	"github.com/raspibuddy/rpi/pkg/utl/mock"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	lext "github.com/shirou/gopsutil/load"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	cases := []struct {
		name       string
		metrics    *mock.Metrics
		lsys       *mocksys.Load
		wantedData rpi.Load
		wantedErr  error
	}{
		{
			name: "error: temp & procs variables are nil",
			metrics: &mock.Metrics{
				LoadAvgFn: func() (lext.AvgStat, error) {
					return lext.AvgStat{}, errors.New("test error info")
				},
				LoadProcsFn: func() (lext.MiscStat, error) {
					return lext.MiscStat{}, errors.New("test error info")
				},
			},
			wantedData: rpi.Load{},
			wantedErr:  echo.NewHTTPError(http.StatusInternalServerError, "could not list the load metrics"),
		},
		{
			name: "error: temp variable is nil",
			metrics: &mock.Metrics{
				LoadAvgFn: func() (lext.AvgStat, error) {
					return lext.AvgStat{}, errors.New("test error info")
				},
				LoadProcsFn: func() (lext.MiscStat, error) {
					return lext.MiscStat{
						ProcsTotal:   4,
						ProcsBlocked: 3,
						ProcsRunning: 1,
					}, nil
				},
			},
			wantedData: rpi.Load{},
			wantedErr:  echo.NewHTTPError(http.StatusInternalServerError, "could not list the load metrics"),
		},
		{
			name: "error: procs variable is nil",
			metrics: &mock.Metrics{
				LoadAvgFn: func() (lext.AvgStat, error) {
					return lext.AvgStat{
						Load1:  1.1,
						Load5:  5.5,
						Load15: 15.15,
					}, errors.New("test error info")
				},
				LoadProcsFn: func() (lext.MiscStat, error) {
					return lext.MiscStat{}, nil
				},
			},
			wantedData: rpi.Load{},
			wantedErr:  echo.NewHTTPError(http.StatusInternalServerError, "could not list the load metrics"),
		},
		{
			name: "success",
			metrics: &mock.Metrics{
				LoadAvgFn: func() (lext.AvgStat, error) {
					return lext.AvgStat{
						Load1:  1.1,
						Load5:  5.5,
						Load15: 15.15,
					}, nil
				},
				LoadProcsFn: func() (lext.MiscStat, error) {
					return lext.MiscStat{
						ProcsTotal:   4,
						ProcsBlocked: 3,
						ProcsRunning: 1,
					}, nil
				},
			},
			lsys: &mocksys.Load{
				ListFn: func(lext.AvgStat, lext.MiscStat) (rpi.Load, error) {
					return rpi.Load{
						Load1:        1.1,
						Load5:        5.5,
						Load15:       15.15,
						ProcsTotal:   4,
						ProcsBlocked: 3,
						ProcsRunning: 1,
					}, nil
				},
			},
			wantedData: rpi.Load{
				Load1:        1.1,
				Load5:        5.5,
				Load15:       15.15,
				ProcsTotal:   4,
				ProcsBlocked: 3,
				ProcsRunning: 1,
			},
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := load.New(tc.lsys, tc.metrics)
			loads, err := s.List()
			assert.Equal(t, tc.wantedData, loads)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
