package cpu

import (
	"errors"
	"net/http"
	"testing"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
	cext "github.com/shirou/gopsutil/cpu"
	"github.com/stretchr/testify/assert"
)

func TestListError(t *testing.T) {
	csys := &mocksys.CPU{
		ListFn: func() ([]cext.InfoStat, []float64, []cext.TimesStat, error) {
			return nil, nil, nil, errors.New("test error")
		},
	}

	s := New(csys)
	cpus, err := s.List()
	assert.Nil(t, cpus)
	assert.NotNil(t, err)
	assert.EqualValues(t, "Results were not returned as they could not be guaranteed", err.(*echo.HTTPError).Message)
	assert.EqualValues(t, http.StatusAccepted, err.(*echo.HTTPError).Code)
}

func TestListLengthDiff(t *testing.T) {
	cases := []struct {
		name string
		csys *mocksys.CPU
	}{
		{
			name: "Length array info greater than array percent & time",
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
			name: "Length array percent greater than array info & time",
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
			name: "Length array time greater than array info & percent",
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
			assert.EqualValues(t, "Results were not returned as they could not be guaranteed", err.(*echo.HTTPError).Message)
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

	s := New(csys)
	cpus, err := s.List()
	assert.EqualValues(t, cpus[0].ID, 0)
	assert.EqualValues(t, cpus[0].ModelName, "test model")
	assert.EqualValues(t, cpus[0].Cores, int32(16))
	assert.EqualValues(t, cpus[0].Mhz, 2300.99)
	assert.EqualValues(t, cpus[0].Stats.Used, 99.9)
	assert.EqualValues(t, cpus[0].Stats.User, 111.1)
	assert.EqualValues(t, cpus[0].Stats.System, 222.2)
	assert.EqualValues(t, cpus[0].Stats.Idle, 333.3)
	assert.Nil(t, err)
}
