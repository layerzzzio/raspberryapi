package process_test

import (
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/process"
	"github.com/raspibuddy/rpi/pkg/utl/metrics"
	"github.com/raspibuddy/rpi/pkg/utl/mock"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	cases := []struct {
		name       string
		metrics    mock.Metrics
		psys       mocksys.Process
		wantedData []rpi.ProcessSummary
		wantedErr  error
	}{
		{
			name: "error: pinfo is nil",
			metrics: mock.Metrics{
				ProcessesFn: func(id ...int32) ([]metrics.PInfo, error) {
					return nil, errors.New("test error pinfo")
				},
			},
			wantedData: nil,
			wantedErr:  echo.NewHTTPError(http.StatusInternalServerError, "could not list the process metrics"),
		},
		{
			name: "success",
			metrics: mock.Metrics{
				ProcessesFn: func(id ...int32) ([]metrics.PInfo, error) {
					return []metrics.PInfo{
						{
							ID:         int32(1),
							Name:       "process_1",
							CPUPercent: 1.1,
							MemPercent: 2.2,
						},
						{
							ID:         int32(2),
							Name:       "process_2",
							CPUPercent: 3.3,
							MemPercent: 4.4,
						},
					}, nil
				},
			},
			psys: mocksys.Process{
				ListFn: func([]metrics.PInfo) ([]rpi.ProcessSummary, error) {
					return []rpi.ProcessSummary{
						{
							ID:         int32(1),
							Name:       "process_1",
							CPUPercent: 1.1,
							MemPercent: 2.2,
						},
						{
							ID:         int32(2),
							Name:       "process_2",
							CPUPercent: 3.3,
							MemPercent: 4.4,
						},
					}, nil
				},
			},
			wantedData: []rpi.ProcessSummary{
				{
					ID:         int32(1),
					Name:       "process_1",
					CPUPercent: 1.1,
					MemPercent: 2.2,
				},
				{
					ID:         int32(2),
					Name:       "process_2",
					CPUPercent: 3.3,
					MemPercent: 4.4,
				},
			},
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := process.New(&tc.psys, tc.metrics)
			ps, err := s.List()
			assert.Equal(t, tc.wantedData, ps)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestView(t *testing.T) {
	cases := []struct {
		name       string
		id         int32
		metrics    mock.Metrics
		psys       mocksys.Process
		wantedData rpi.ProcessDetails
		wantedErr  error
	}{
		{
			name: "error: pinfo is nil",
			id:   int32(99),
			metrics: mock.Metrics{
				ProcessesFn: func(id ...int32) ([]metrics.PInfo, error) {
					return nil, errors.New("test error pinfo")
				},
			},
			wantedData: rpi.ProcessDetails{},
			wantedErr:  echo.NewHTTPError(http.StatusInternalServerError, "could not view the process metrics"),
		},
		{
			name: "error: process not found",
			id:   int32(66),
			metrics: mock.Metrics{
				ProcessesFn: func(id ...int32) ([]metrics.PInfo, error) {
					return nil, errors.New("process not found")
				},
			},
			wantedData: rpi.ProcessDetails{},
			wantedErr:  echo.NewHTTPError(http.StatusInternalServerError, "process 66 does not exist"),
		},
		{
			name: "success",
			id:   int32(99),
			metrics: mock.Metrics{
				ProcessesFn: func(id ...int32) ([]metrics.PInfo, error) {
					return []metrics.PInfo{
						{
							ID:           int32(99),
							Name:         "process_99",
							CPUPercent:   1.1,
							MemPercent:   2.2,
							Username:     "pi",
							CommandLine:  "/cmd/text",
							Status:       "S",
							CreationTime: time.Time{}.Add(1666666),
							Foreground:   true,
							Background:   false,
							IsRunning:    true,
							ParentP:      1,
						},
					}, nil
				},
			},
			psys: mocksys.Process{
				ViewFn: func(int32, []metrics.PInfo) (rpi.ProcessDetails, error) {
					return rpi.ProcessDetails{
						ID:           int32(99),
						Name:         "process_99",
						CPUPercent:   1.1,
						MemPercent:   2.2,
						Username:     "pi",
						CommandLine:  "/cmd/text",
						Status:       "S",
						CreationTime: time.Time{}.Add(1666666),
						Foreground:   true,
						Background:   false,
						IsRunning:    true,
						ParentP:      int32(1),
					}, nil
				},
			},
			wantedData: rpi.ProcessDetails{
				ID:           int32(99),
				Name:         "process_99",
				CPUPercent:   1.1,
				MemPercent:   2.2,
				Username:     "pi",
				CommandLine:  "/cmd/text",
				Status:       "S",
				CreationTime: time.Time{}.Add(1666666),
				Foreground:   true,
				Background:   false,
				IsRunning:    true,
				ParentP:      int32(1),
			},
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := process.New(&tc.psys, tc.metrics)
			ps, err := s.View(tc.id)
			assert.Equal(t, tc.wantedData, ps)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
