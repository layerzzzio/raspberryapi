package cpu_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/metrics/cpu"
	"github.com/raspibuddy/rpi/pkg/utl/mock"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	cext "github.com/shirou/gopsutil/cpu"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	cases := []struct {
		name       string
		metrics    mock.Metrics
		csys       mocksys.CPU
		wantedData []rpi.CPU
		wantedErr  error
	}{
		{
			name: "error: info & percent & times arrays are nil",
			metrics: mock.Metrics{
				CPUInfoFn: func() ([]cext.InfoStat, error) {
					return nil, errors.New("test error info")
				},
				CPUPercentFn: func(time.Duration, bool) ([]float64, error) {
					return nil, errors.New("test error percent")
				},
				CPUTimesFn: func(bool) ([]cext.TimesStat, error) {
					return nil, errors.New("test error times")
				},
			},
			wantedData: nil,
			wantedErr:  echo.NewHTTPError(http.StatusInternalServerError, "could not retrieve the cpu metrics"),
		},
		{
			name: "error: percent & times arrays are both nil",
			metrics: mock.Metrics{
				CPUInfoFn: func() ([]cext.InfoStat, error) {
					return []cext.InfoStat{
						{
							CPU: int32(10),
						},
					}, nil
				},
				CPUPercentFn: func(time.Duration, bool) ([]float64, error) {
					return nil, errors.New("test error percent")
				},
				CPUTimesFn: func(bool) ([]cext.TimesStat, error) {
					return nil, errors.New("test error times")
				},
			},
			wantedData: nil,
			wantedErr:  echo.NewHTTPError(http.StatusInternalServerError, "could not retrieve the cpu metrics"),
		},
		{
			name: "error: info & times arrays are both nil",
			metrics: mock.Metrics{
				CPUInfoFn: func() ([]cext.InfoStat, error) {
					return nil, errors.New("test error info")
				},
				CPUPercentFn: func(time.Duration, bool) ([]float64, error) {
					return []float64{99.9}, nil
				},
				CPUTimesFn: func(bool) ([]cext.TimesStat, error) {
					return nil, errors.New("test error times")
				},
			},
			wantedData: nil,
			wantedErr:  echo.NewHTTPError(http.StatusInternalServerError, "could not retrieve the cpu metrics"),
		},
		{
			name: "error: info & percent arrays are both nil",
			metrics: mock.Metrics{
				CPUInfoFn: func() ([]cext.InfoStat, error) {
					return nil, errors.New("test error info")
				},
				CPUPercentFn: func(time.Duration, bool) ([]float64, error) {
					return []float64{99.9}, nil
				},
				CPUTimesFn: func(bool) ([]cext.TimesStat, error) {
					return []cext.TimesStat{
						{
							CPU: "cpu-total",
						},
					}, nil
				},
			},
			wantedData: nil,
			wantedErr:  echo.NewHTTPError(http.StatusInternalServerError, "could not retrieve the cpu metrics"),
		},
		{
			name: "success",
			metrics: mock.Metrics{
				CPUInfoFn: func() ([]cext.InfoStat, error) {
					return []cext.InfoStat{
						{
							CPU:       0,
							ModelName: "intel processor",
							Cores:     int32(8),
							Mhz:       2300.99,
						},
					}, nil
				},
				CPUPercentFn: func(time.Duration, bool) ([]float64, error) {
					return []float64{99.9}, nil
				},
				CPUTimesFn: func(bool) ([]cext.TimesStat, error) {
					return []cext.TimesStat{
						{
							User:   111.1,
							System: 222.2,
							Idle:   333.3,
						},
					}, nil
				},
			},
			csys: mocksys.CPU{
				ListFn: func([]cext.InfoStat, []float64, []cext.TimesStat) ([]rpi.CPU, error) {
					return []rpi.CPU{
						{
							ID:        1,
							Cores:     int32(8),
							ModelName: "intel processor",
							Mhz:       2300.99,
							Stats: rpi.CPUStats{
								Used:   99.9,
								User:   111.1,
								System: 222.2,
								Idle:   333.3,
							},
						},
					}, nil
				},
			},
			wantedData: []rpi.CPU{
				{
					ID:        1,
					Cores:     int32(8),
					ModelName: "intel processor",
					Mhz:       2300.99,
					Stats: rpi.CPUStats{
						Used:   99.9,
						User:   111.1,
						System: 222.2,
						Idle:   333.3,
					},
				},
			},
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := cpu.New(&tc.csys, tc.metrics)
			cpus, err := s.List()
			assert.Equal(t, tc.wantedData, cpus)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestView(t *testing.T) {
	cases := []struct {
		name       string
		id         int
		metrics    mock.Metrics
		csys       mocksys.CPU
		wantedData rpi.CPU
		wantedErr  error
	}{
		{
			name: "error: info & percent & times arrays are nil",
			metrics: mock.Metrics{
				CPUInfoFn: func() ([]cext.InfoStat, error) {
					return nil, errors.New("test error info")
				},
				CPUPercentFn: func(time.Duration, bool) ([]float64, error) {
					return nil, errors.New("test error percent")
				},
				CPUTimesFn: func(bool) ([]cext.TimesStat, error) {
					return nil, errors.New("test error times")
				},
			},
			wantedData: rpi.CPU{},
			wantedErr:  echo.NewHTTPError(http.StatusInternalServerError, "could not retrieve the cpu metrics"),
		},
		{
			name: "error: percent & times arrays are both nil",
			metrics: mock.Metrics{
				CPUInfoFn: func() ([]cext.InfoStat, error) {
					return []cext.InfoStat{
						{
							CPU: int32(10),
						},
					}, nil
				},
				CPUPercentFn: func(time.Duration, bool) ([]float64, error) {
					return nil, errors.New("test error percent")
				},
				CPUTimesFn: func(bool) ([]cext.TimesStat, error) {
					return nil, errors.New("test error times")
				},
			},
			wantedData: rpi.CPU{},
			wantedErr:  echo.NewHTTPError(http.StatusInternalServerError, "could not retrieve the cpu metrics"),
		},
		{
			name: "error: info & times arrays are both nil",
			metrics: mock.Metrics{
				CPUInfoFn: func() ([]cext.InfoStat, error) {
					return nil, errors.New("test error info")
				},
				CPUPercentFn: func(time.Duration, bool) ([]float64, error) {
					return []float64{99.9}, nil
				},
				CPUTimesFn: func(bool) ([]cext.TimesStat, error) {
					return nil, errors.New("test error times")
				},
			},
			wantedData: rpi.CPU{},
			wantedErr:  echo.NewHTTPError(http.StatusInternalServerError, "could not retrieve the cpu metrics"),
		},
		{
			name: "error: info & percent arrays are both nil",
			metrics: mock.Metrics{
				CPUInfoFn: func() ([]cext.InfoStat, error) {
					return nil, errors.New("test error info")
				},
				CPUPercentFn: func(time.Duration, bool) ([]float64, error) {
					return []float64{99.9}, nil
				},
				CPUTimesFn: func(bool) ([]cext.TimesStat, error) {
					return []cext.TimesStat{
						{
							CPU: "cpu-total",
						},
					}, nil
				},
			},
			wantedData: rpi.CPU{},
			wantedErr:  echo.NewHTTPError(http.StatusInternalServerError, "could not retrieve the cpu metrics"),
		},
		{
			name: "success",
			id:   1,
			metrics: mock.Metrics{
				CPUInfoFn: func() ([]cext.InfoStat, error) {
					return []cext.InfoStat{
						{
							CPU:       0,
							ModelName: "intel processor",
							Cores:     int32(8),
							Mhz:       2300.99,
						},
					}, nil
				},
				CPUPercentFn: func(time.Duration, bool) ([]float64, error) {
					return []float64{99.9}, nil
				},
				CPUTimesFn: func(bool) ([]cext.TimesStat, error) {
					return []cext.TimesStat{
						{
							User:   111.1,
							System: 222.2,
							Idle:   333.3,
						},
					}, nil
				},
			},
			csys: mocksys.CPU{
				ViewFn: func(int, []cext.InfoStat, []float64, []cext.TimesStat) (rpi.CPU, error) {
					return rpi.CPU{
						ID:        1,
						Cores:     int32(8),
						ModelName: "intel processor",
						Mhz:       2300.99,
						Stats: rpi.CPUStats{
							Used:   99.9,
							User:   111.1,
							System: 222.2,
							Idle:   333.3,
						},
					}, nil
				},
			},
			wantedData: rpi.CPU{
				ID:        1,
				Cores:     int32(8),
				ModelName: "intel processor",
				Mhz:       2300.99,
				Stats: rpi.CPUStats{
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
			s := cpu.New(&tc.csys, tc.metrics)
			cpus, err := s.View(tc.id)
			assert.Equal(t, tc.wantedData, cpus)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
