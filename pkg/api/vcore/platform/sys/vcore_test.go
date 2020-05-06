package sys

import (
	"errors"
	"net/http"
	"testing"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/vcore"
	"github.com/shirou/gopsutil/cpu"
	"github.com/stretchr/testify/assert"
)

func TestNumExtract(t *testing.T) {
	cases := []struct {
		name       string
		input      string
		min        int
		max        int
		wantedData []string
	}{
		{
			name:       "success: input string containing a single numerical character",
			input:      "cpu0",
			min:        0,
			max:        9,
			wantedData: []string{"0"},
		},
		{
			name:       "success: input string containing multiple consecutive numerical characters",
			input:      "cpu012",
			min:        0,
			max:        9,
			wantedData: []string{"012"},
		},
		{
			name:       "success: input string containing multiple non-consecutive numerical characters",
			input:      "0cpu12xxx3xxx9",
			min:        0,
			max:        9,
			wantedData: []string{"0", "12", "3", "9"},
		},
		{
			name:       "success: input string without any numerical characters",
			input:      "cpu",
			min:        0,
			max:        9,
			wantedData: []string(nil),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			out := extractNum(tc.input, tc.min, tc.max)
			assert.EqualValues(t, out, tc.wantedData)
		})
	}
}

func TestConcatID(t *testing.T) {
	cases := []struct {
		name       string
		input      []string
		wantedData int
		wantedErr  error
	}{
		{
			name:       "error: invalid argument type",
			input:      []string{"X"},
			wantedData: -1,
			wantedErr:  errors.New("invalid syntax"),
		},
		{
			name:       "success: input string containing a single numerical character",
			input:      []string{"0"},
			wantedData: 0,
			wantedErr:  nil,
		},
		{
			name:       "success: input string containing multiple consecutive numerical characters",
			input:      []string{"012"},
			wantedData: 12,
			wantedErr:  nil,
		},
		{
			name:       "success: input string containing multiple non-consecutive numerical characters",
			input:      []string{"0", "1", "2", "3"},
			wantedData: 123,
			wantedErr:  nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			out, err := concatID(tc.input)
			assert.EqualValues(t, out, tc.wantedData)
			assert.EqualValues(t, err, tc.wantedErr)
		})
	}
}

func TestList(t *testing.T) {
	cases := []struct {
		name       string
		percent    []float64
		times      []cpu.TimesStat
		wantedData []rpi.VCore
		wantedErr  error
	}{
		{
			name:    "error: times array length greater than percent array length",
			percent: []float64{99.9},
			times: []cpu.TimesStat{
				{
					CPU:    "cpu0",
					User:   111.1,
					System: 222.2,
					Idle:   333.3,
				},
				{
					CPU:    "cpu1",
					User:   444.4,
					System: 555.5,
					Idle:   666.6,
				},
			},
			wantedData: nil,
			wantedErr:  echo.NewHTTPError(http.StatusNotFound, "results were not returned as they could not be guaranteed"),
		},
		{
			name:    "error: times array length greater than percent array length",
			percent: []float64{50.0, 99.9},
			times: []cpu.TimesStat{
				{
					CPU:    "cpu0",
					User:   111.1,
					System: 222.2,
					Idle:   333.3,
				},
			},
			wantedData: nil,
			wantedErr:  echo.NewHTTPError(http.StatusNotFound, "results were not returned as they could not be guaranteed"),
		},
		{
			name:    "error: parsing ID failed",
			percent: []float64{99.9},
			times: []cpu.TimesStat{
				{
					CPU:    "cpu",
					User:   111.1,
					System: 222.2,
					Idle:   333.3,
				},
			},
			wantedData: nil,
			wantedErr:  echo.NewHTTPError(http.StatusNotFound, "parsing id was unsuccessful"),
		},
		{
			name:    "success",
			percent: []float64{99.9},
			times: []cpu.TimesStat{
				{
					CPU:    "cpu0",
					User:   111.1,
					System: 222.2,
					Idle:   333.3,
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
			s := vcore.VSYS(VCore{})
			vcores, err := s.List(tc.percent, tc.times)
			assert.Equal(t, tc.wantedData, vcores)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestView(t *testing.T) {
	cases := []struct {
		name       string
		id         int
		percent    []float64
		times      []cpu.TimesStat
		wantedData rpi.VCore
		wantedErr  error
	}{
		{
			name:    "error: times array length greater than percent array length",
			percent: []float64{99.9},
			times: []cpu.TimesStat{
				{
					CPU:    "cpu0",
					User:   111.1,
					System: 222.2,
					Idle:   333.3,
				},
				{
					CPU:    "cpu1",
					User:   444.4,
					System: 555.5,
					Idle:   666.6,
				},
			},
			wantedData: rpi.VCore{},
			wantedErr:  echo.NewHTTPError(http.StatusNotFound, "results were not returned as they could not be guaranteed"),
		},
		{
			name:    "error: times array length greater than percent array length",
			percent: []float64{50.0, 99.9},
			times: []cpu.TimesStat{
				{
					CPU:    "cpu0",
					User:   111.1,
					System: 222.2,
					Idle:   333.3,
				},
			},
			wantedData: rpi.VCore{},
			wantedErr:  echo.NewHTTPError(http.StatusNotFound, "results were not returned as they could not be guaranteed"),
		},
		{
			name:    "error: id value greater than vcore number",
			id:      2,
			percent: []float64{99.9},
			times: []cpu.TimesStat{
				{
					CPU:    "cpu0",
					User:   111.1,
					System: 222.2,
					Idle:   333.3,
				},
			},
			wantedData: rpi.VCore{},
			wantedErr:  echo.NewHTTPError(http.StatusNotFound, "id out of range"),
		},
		{
			name:    "error: id value is negative",
			id:      -1,
			percent: []float64{99.9},
			times: []cpu.TimesStat{
				{
					CPU:    "cpu0",
					User:   111.1,
					System: 222.2,
					Idle:   333.3,
				},
			},
			wantedData: rpi.VCore{},
			wantedErr:  echo.NewHTTPError(http.StatusNotFound, "id out of range"),
		},
		{
			name:    "error: id value equals 0",
			id:      0,
			percent: []float64{99.9},
			times: []cpu.TimesStat{
				{
					CPU:    "cpu0",
					User:   111.1,
					System: 222.2,
					Idle:   333.3,
				},
			},
			wantedData: rpi.VCore{},
			wantedErr:  echo.NewHTTPError(http.StatusNotFound, "id out of range"),
		},
		{
			name:    "error: parsing ID failed",
			id:      1,
			percent: []float64{99.9},
			times: []cpu.TimesStat{
				{
					CPU:    "cpu",
					User:   111.1,
					System: 222.2,
					Idle:   333.3,
				},
			},
			wantedData: rpi.VCore{},
			wantedErr:  echo.NewHTTPError(http.StatusNotFound, "parsing id was unsuccessful"),
		},
		{
			name:    "success",
			id:      1,
			percent: []float64{99.9},
			times: []cpu.TimesStat{
				{
					CPU:    "cpu0",
					User:   111.1,
					System: 222.2,
					Idle:   333.3,
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
			s := vcore.VSYS(VCore{})
			vcores, err := s.View(tc.id, tc.percent, tc.times)
			assert.Equal(t, tc.wantedData, vcores)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
