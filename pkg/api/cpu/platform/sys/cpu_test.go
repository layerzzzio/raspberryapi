package sys_test

import (
	"net/http"
	"testing"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/cpu"
	"github.com/raspibuddy/rpi/pkg/api/cpu/platform/sys"
	cext "github.com/shirou/gopsutil/cpu"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	cases := []struct {
		name       string
		info       []cext.InfoStat
		percent    []float64
		times      []cext.TimesStat
		wantedData []rpi.CPU
		wantedErr  error
	}{
		{
			name: "error: info array length greater than percent & times array length",
			info: []cext.InfoStat{
				{
					CPU: int32(0),
				},
				{
					CPU: int32(1),
				},
			},
			percent: []float64{99.9},
			times: []cext.TimesStat{
				{
					CPU:    "cpu-total",
					User:   111.1,
					System: 222.2,
					Idle:   333.3,
				},
			},
			wantedData: nil,
			wantedErr:  echo.NewHTTPError(http.StatusNotFound, "results were not returned as they could not be guaranteed"),
		},
		{
			name: "error: percent array length greater than info & times array length",
			info: []cext.InfoStat{
				{
					CPU: int32(0),
				},
			},
			percent: []float64{50.0, 99.9},
			times: []cext.TimesStat{
				{
					CPU:    "cpu-total",
					User:   111.1,
					System: 222.2,
					Idle:   333.3,
				},
			},
			wantedData: nil,
			wantedErr:  echo.NewHTTPError(http.StatusNotFound, "results were not returned as they could not be guaranteed"),
		},
		{
			name: "error: times array length greater than info & percent array length",
			info: []cext.InfoStat{
				{
					CPU: int32(0),
				},
			},
			percent: []float64{99.9},
			times: []cext.TimesStat{
				{
					CPU:    "cpu-total-0",
					User:   111.1,
					System: 222.2,
					Idle:   333.3,
				},
				{
					CPU:    "cpu-total-1",
					User:   444.4,
					System: 555.5,
					Idle:   666.6,
				},
			},
			wantedData: nil,
			wantedErr:  echo.NewHTTPError(http.StatusNotFound, "results were not returned as they could not be guaranteed"),
		},
		{
			name: "success",
			info: []cext.InfoStat{
				{
					CPU:       int32(0),
					Cores:     int32(8),
					ModelName: "intel processor",
					Mhz:       2300.99,
				},
			},
			percent: []float64{99.9},
			times: []cext.TimesStat{
				{
					CPU:    "cpu-total",
					User:   111.1,
					System: 222.2,
					Idle:   333.3,
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
			s := cpu.CSYS(sys.CPU{})
			cpus, err := s.List(tc.info, tc.percent, tc.times)
			assert.Equal(t, tc.wantedData, cpus)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestView(t *testing.T) {
	cases := []struct {
		name       string
		id         int
		info       []cext.InfoStat
		percent    []float64
		times      []cext.TimesStat
		wantedData rpi.CPU
		wantedErr  error
	}{
		{
			name: "error: info array length greater than percent & times array length",
			info: []cext.InfoStat{
				{
					CPU: int32(0),
				},
				{
					CPU: int32(1),
				},
			},
			percent: []float64{99.9},
			times: []cext.TimesStat{
				{
					CPU:    "cpu-total",
					User:   111.1,
					System: 222.2,
					Idle:   333.3,
				},
			},
			wantedData: rpi.CPU{},
			wantedErr:  echo.NewHTTPError(http.StatusNotFound, "results were not returned as they could not be guaranteed"),
		},
		{
			name: "error: percent array length greater than info & times array length",
			info: []cext.InfoStat{
				{
					CPU: int32(0),
				},
			},
			percent: []float64{50.0, 99.9},
			times: []cext.TimesStat{
				{
					CPU:    "cpu-total",
					User:   111.1,
					System: 222.2,
					Idle:   333.3,
				},
			},
			wantedData: rpi.CPU{},
			wantedErr:  echo.NewHTTPError(http.StatusNotFound, "results were not returned as they could not be guaranteed"),
		},
		{
			name: "error: times array length greater than info & percent array length",
			info: []cext.InfoStat{
				{
					CPU: int32(0),
				},
			},
			percent: []float64{99.9},
			times: []cext.TimesStat{
				{
					CPU:    "cpu-total-0",
					User:   111.1,
					System: 222.2,
					Idle:   333.3,
				},
				{
					CPU:    "cpu-total-1",
					User:   111.1,
					System: 222.2,
					Idle:   333.3,
				},
			},
			wantedData: rpi.CPU{},
			wantedErr:  echo.NewHTTPError(http.StatusNotFound, "results were not returned as they could not be guaranteed"),
		},
		{
			name: "error: id value greater than cpu number",
			id:   2,
			info: []cext.InfoStat{
				{
					CPU: int32(0),
				},
			},
			percent: []float64{99.9},
			times: []cext.TimesStat{
				{
					CPU:    "cpu-total",
					User:   111.1,
					System: 222.2,
					Idle:   333.3,
				},
			},
			wantedData: rpi.CPU{},
			wantedErr:  echo.NewHTTPError(http.StatusNotFound, "id out of range"),
		},
		{
			name: "error: id value is negative",
			id:   -1,
			info: []cext.InfoStat{
				{
					CPU: int32(0),
				},
			},
			percent: []float64{99.9},
			times: []cext.TimesStat{
				{
					CPU:    "cpu-total",
					User:   111.1,
					System: 222.2,
					Idle:   333.3,
				},
			},
			wantedData: rpi.CPU{},
			wantedErr:  echo.NewHTTPError(http.StatusNotFound, "id out of range"),
		},
		{
			name: "error: id value equals 0",
			id:   0,
			info: []cext.InfoStat{
				{
					CPU: int32(0),
				},
			},
			percent: []float64{99.9},
			times: []cext.TimesStat{
				{
					CPU:    "cpu-total",
					User:   111.1,
					System: 222.2,
					Idle:   333.3,
				},
			},
			wantedData: rpi.CPU{},
			wantedErr:  echo.NewHTTPError(http.StatusNotFound, "id out of range"),
		},
		{
			name: "success",
			id:   1,
			info: []cext.InfoStat{
				{
					CPU:       int32(0),
					Cores:     int32(8),
					ModelName: "intel processor",
					Mhz:       2300.99,
				},
			},
			percent: []float64{99.9},
			times: []cext.TimesStat{
				{
					CPU:    "cpu-total",
					User:   111.1,
					System: 222.2,
					Idle:   333.3,
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
			s := cpu.CSYS(sys.CPU{})
			cpus, err := s.View(tc.id, tc.info, tc.percent, tc.times)
			assert.Equal(t, tc.wantedData, cpus)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
