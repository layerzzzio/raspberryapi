package cpu

import (
	"errors"
	"net/http"
	"testing"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	cext "github.com/shirou/gopsutil/cpu"
	"github.com/stretchr/testify/assert"
)

func TestListNilResultError(t *testing.T) {
	csys := &mocksys.CPU{
		ListFn: func() ([]cext.InfoStat, []float64, []cext.TimesStat, error) {
			return nil, nil, nil, errors.New("test error")
		},
	}

	wantedErr := echo.NewHTTPError(http.StatusAccepted, "results were not returned as they could not be guaranteed")

	s := New(csys)
	cpus, err := s.List()
	assert.Nil(t, cpus)
	assert.NotNil(t, err)
	assert.Equal(t, err, wantedErr)
}

func TestListLengthDiffError(t *testing.T) {
	cases := []struct {
		name string
		csys *mocksys.CPU
	}{
		{
			name: "length array info greater than percent & time arrays",
			csys: &mocksys.CPU{
				ListFn: func() ([]cext.InfoStat, []float64, []cext.TimesStat, error) {
					return []cext.InfoStat{
							{
								CPU:       0,
								ModelName: "Intel P0",
								Cores:     12,
								Mhz:       35.56,
							},
							{
								CPU:       1,
								ModelName: "Intel P1",
								Cores:     78,
								Mhz:       910.1112,
							},
						},
						[]float64{99.9},
						[]cext.TimesStat{
							{
								CPU:    "total-cpu-0",
								User:   111.1,
								System: 222.2,
								Idle:   333.3,
							},
						}, nil
				},
			},
		},
		{
			name: "length array percent greater than info & time arrays",
			csys: &mocksys.CPU{
				ListFn: func() ([]cext.InfoStat, []float64, []cext.TimesStat, error) {
					return []cext.InfoStat{
							{
								CPU:       0,
								ModelName: "Intel P0",
								Cores:     12,
								Mhz:       34.56,
							},
						},
						[]float64{99.9, 50.0},
						[]cext.TimesStat{
							{
								CPU:    "total-cpu-0",
								User:   111.1,
								System: 222.2,
								Idle:   333.3,
							},
						}, nil
				},
			},
		},
		{
			name: "length array time greater than info & percent arrays",
			csys: &mocksys.CPU{
				ListFn: func() ([]cext.InfoStat, []float64, []cext.TimesStat, error) {
					return []cext.InfoStat{
							{
								CPU:       0,
								ModelName: "Intel P0",
								Cores:     12,
								Mhz:       34.56,
							},
						},
						[]float64{99.9},
						[]cext.TimesStat{
							{
								CPU:    "total-cpu-0",
								User:   111.1,
								System: 222.2,
								Idle:   333.3,
							},
							{
								CPU:    "total-cpu-1",
								User:   444.4,
								System: 555.5,
								Idle:   666.6,
							},
						}, nil
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := New(tc.csys)
			cpus, err := s.List()
			assert.Nil(t, cpus)
			assert.NotNil(t, err)
			assert.EqualValues(t, "results were not returned as they could not be guaranteed", err.(*echo.HTTPError).Message)
			assert.EqualValues(t, http.StatusAccepted, err.(*echo.HTTPError).Code)
		})
	}
}

func TestListSuccess(t *testing.T) {
	csys := &mocksys.CPU{
		ListFn: func() ([]cext.InfoStat, []float64, []cext.TimesStat, error) {
			return []cext.InfoStat{
					{
						CPU:       0,
						ModelName: "test model",
						Cores:     16,
						Mhz:       2300.99,
					},
				},
				[]float64{99.9},
				[]cext.TimesStat{
					{
						CPU:    "total-cpu",
						User:   111.1,
						System: 222.2,
						Idle:   333.3,
					},
				}, nil
		},
	}

	wantedData := rpi.CPU{
		ID:        0,
		ModelName: "test model",
		Cores:     int32(16),
		Mhz:       2300.99,
		Stats: rpi.CPUStats{
			Used:   99.9,
			User:   111.1,
			System: 222.2,
			Idle:   333.3,
		},
	}

	s := New(csys)
	cpus, err := s.List()
	assert.Equal(t, cpus[0], wantedData)
	assert.Nil(t, err)
}
