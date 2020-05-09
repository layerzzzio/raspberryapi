package vcore_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/vcore"
	"github.com/raspibuddy/rpi/pkg/utl/mock"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	"github.com/shirou/gopsutil/cpu"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	cases := []struct {
		name       string
		metrics    *mock.Metrics
		vsys       *mocksys.VCore
		wantedData []rpi.VCore
		wantedErr  error
	}{
		{
			name: "error: info & percent arrays are both nil",
			metrics: &mock.Metrics{
				CPUPercentFn: func(time.Duration, bool) ([]float64, error) {
					return nil, errors.New("test error percent")
				},
				CPUTimesFn: func(bool) ([]cpu.TimesStat, error) {
					return nil, errors.New("test error times")
				},
			},
			wantedData: nil,
			wantedErr:  echo.NewHTTPError(http.StatusInternalServerError, "could not retrieve the vcore metrics"),
		},
		{
			name: "error: info array is nil",
			metrics: &mock.Metrics{
				CPUPercentFn: func(time.Duration, bool) ([]float64, error) {
					return []float64{99.9}, nil
				},
				CPUTimesFn: func(bool) ([]cpu.TimesStat, error) {
					return nil, errors.New("test error times")
				},
			},
			wantedData: nil,
			wantedErr:  echo.NewHTTPError(http.StatusInternalServerError, "could not retrieve the vcore metrics"),
		},
		{
			name: "error: percent array is nil",
			metrics: &mock.Metrics{
				CPUPercentFn: func(time.Duration, bool) ([]float64, error) {
					return nil, errors.New("test error percent")
				},
				CPUTimesFn: func(bool) ([]cpu.TimesStat, error) {
					return []cpu.TimesStat{
						{
							CPU:    "cpu0",
							User:   111.1,
							System: 222.2,
							Idle:   333.3,
						},
					}, nil
				},
			},
			wantedData: nil,
			wantedErr:  echo.NewHTTPError(http.StatusInternalServerError, "could not retrieve the vcore metrics"),
		},
		{
			name: "success",
			metrics: &mock.Metrics{
				CPUPercentFn: func(time.Duration, bool) ([]float64, error) {
					return []float64{99.9}, nil
				},
				CPUTimesFn: func(bool) ([]cpu.TimesStat, error) {
					return []cpu.TimesStat{
						{
							CPU:    "cpu0",
							User:   111.1,
							System: 222.2,
							Idle:   333.3,
						},
					}, nil
				},
			},
			vsys: &mocksys.VCore{
				ListFn: func([]float64, []cpu.TimesStat) ([]rpi.VCore, error) {
					return []rpi.VCore{
						{
							ID:     1,
							Used:   99.9,
							User:   111.1,
							System: 222.2,
							Idle:   333.3,
						},
					}, nil
				},
			},
			wantedData: []rpi.VCore{
				{
					ID:     1,
					Used:   99.9,
					User:   111.1,
					System: 222.2,
					Idle:   333.3,
				},
			},
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := vcore.New(tc.vsys, tc.metrics)
			vcores, err := s.List()
			assert.Equal(t, tc.wantedData, vcores)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestView(t *testing.T) {
	cases := []struct {
		name       string
		id         int
		metrics    *mock.Metrics
		vsys       *mocksys.VCore
		wantedData rpi.VCore
		wantedErr  error
	}{
		{
			name: "error: info & percent arrays are both nil",
			metrics: &mock.Metrics{
				CPUPercentFn: func(time.Duration, bool) ([]float64, error) {
					return nil, errors.New("test error percent")
				},
				CPUTimesFn: func(bool) ([]cpu.TimesStat, error) {
					return nil, errors.New("test error times")
				},
			},
			wantedData: rpi.VCore{},
			wantedErr:  echo.NewHTTPError(http.StatusInternalServerError, "could not retrieve the vcore metrics"),
		},
		{
			name: "error: info array is nil",
			metrics: &mock.Metrics{
				CPUPercentFn: func(time.Duration, bool) ([]float64, error) {
					return []float64{99.9}, nil
				},
				CPUTimesFn: func(bool) ([]cpu.TimesStat, error) {
					return nil, errors.New("test error times")
				},
			},
			wantedData: rpi.VCore{},
			wantedErr:  echo.NewHTTPError(http.StatusInternalServerError, "could not retrieve the vcore metrics"),
		},
		{
			name: "error: percent array is nil",
			metrics: &mock.Metrics{
				CPUPercentFn: func(time.Duration, bool) ([]float64, error) {
					return nil, errors.New("test error percent")
				},
				CPUTimesFn: func(bool) ([]cpu.TimesStat, error) {
					return []cpu.TimesStat{
						{
							CPU:    "cpu0",
							User:   111.1,
							System: 222.2,
							Idle:   333.3,
						},
					}, nil
				},
			},
			wantedData: rpi.VCore{},
			wantedErr:  echo.NewHTTPError(http.StatusInternalServerError, "could not retrieve the vcore metrics"),
		},
		{
			name: "success",
			id:   1,
			metrics: &mock.Metrics{
				CPUPercentFn: func(time.Duration, bool) ([]float64, error) {
					return []float64{99.9}, nil
				},
				CPUTimesFn: func(bool) ([]cpu.TimesStat, error) {
					return []cpu.TimesStat{
						{
							CPU:    "cpu0",
							User:   111.1,
							System: 222.2,
							Idle:   333.3,
						},
					}, nil
				},
			},
			vsys: &mocksys.VCore{
				ViewFn: func(int, []float64, []cpu.TimesStat) (rpi.VCore, error) {
					return rpi.VCore{

						ID:     1,
						Used:   99.9,
						User:   111.1,
						System: 222.2,
						Idle:   333.3,
					}, nil
				},
			},
			wantedData: rpi.VCore{
				ID:     1,
				Used:   99.9,
				User:   111.1,
				System: 222.2,
				Idle:   333.3,
			},
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := vcore.New(tc.vsys, tc.metrics)
			vcores, err := s.View(tc.id)
			assert.Equal(t, tc.wantedData, vcores)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
