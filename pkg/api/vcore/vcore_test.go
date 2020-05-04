package vcore

import (
	"errors"
	"net/http"
	"testing"

	"github.com/labstack/echo"
	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/utl/mock/mocksys"
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
		vsys       *mocksys.VCore
		wantedData []rpi.VCore
		wantedErr  error
	}{
		{
			name: "error: info & percent array are nil",
			vsys: &mocksys.VCore{
				ListFn: func() ([]float64, []cpu.TimesStat, error) {
					return nil, nil, echo.NewHTTPError(http.StatusAccepted, "results were not returned as they could not be guaranteed")
				},
			},
			wantedData: nil,
			wantedErr:  echo.NewHTTPError(http.StatusAccepted, "results were not returned as they could not be guaranteed"),
		},
		{
			name: "error: time array length greater than percent array length",
			vsys: &mocksys.VCore{
				ListFn: func() ([]float64, []cpu.TimesStat, error) {
					return []float64{50.0},
						[]cpu.TimesStat{
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
						nil
				},
			},
			wantedData: nil,
			wantedErr:  echo.NewHTTPError(http.StatusAccepted, "results were not returned as they could not be guaranteed"),
		},
		{
			name: "error: percent array length greater than time array length",
			vsys: &mocksys.VCore{
				ListFn: func() ([]float64, []cpu.TimesStat, error) {
					return []float64{50.0, 99.9},
						[]cpu.TimesStat{
							{
								CPU:    "cpu0",
								User:   111.1,
								System: 222.2,
								Idle:   333.3,
							},
						},
						nil
				},
			},
			wantedData: nil,
			wantedErr:  echo.NewHTTPError(http.StatusAccepted, "results were not returned as they could not be guaranteed"),
		},
		{
			name: "error: parsing ID failed",
			vsys: &mocksys.VCore{
				ListFn: func() ([]float64, []cpu.TimesStat, error) {
					return []float64{99.9},
						[]cpu.TimesStat{
							{
								CPU:    "cpu",
								User:   111.1,
								System: 222.2,
								Idle:   333.3,
							},
						},
						nil
				},
			},
			wantedData: nil,
			wantedErr:  echo.NewHTTPError(http.StatusInternalServerError, "parsing id was unsuccessful"),
		},
		{
			name: "success",
			vsys: &mocksys.VCore{
				ListFn: func() ([]float64, []cpu.TimesStat, error) {
					return []float64{99.9},
						[]cpu.TimesStat{
							{
								CPU:    "cpu0",
								User:   111.1,
								System: 222.2,
								Idle:   333.3,
							},
						},
						nil
				},
			},
			wantedData: []rpi.VCore{
				{
					ID:     0,
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
			s := New(tc.vsys)
			vcores, err := s.List()
			assert.Equal(t, tc.wantedData, vcores)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestView(t *testing.T) {
	cases := []struct {
		name       string
		vsys       *mocksys.VCore
		id         int
		wantedData *rpi.VCore
		wantedErr  error
	}{
		{
			name: "error: info & percent array are nil",
			vsys: &mocksys.VCore{
				ListFn: func() ([]float64, []cpu.TimesStat, error) {
					return nil, nil, echo.NewHTTPError(http.StatusAccepted, "results were not returned as they could not be guaranteed")
				},
			},
			wantedData: nil,
			wantedErr:  echo.NewHTTPError(http.StatusAccepted, "results were not returned as they could not be guaranteed"),
		},
		{
			name: "error: time array length greater than percent array length",
			vsys: &mocksys.VCore{
				ListFn: func() ([]float64, []cpu.TimesStat, error) {
					return []float64{50.0},
						[]cpu.TimesStat{
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
						nil
				},
			},
			id:         2,
			wantedData: nil,
			wantedErr:  echo.NewHTTPError(http.StatusAccepted, "results were not returned as they could not be guaranteed"),
		},
		{
			name: "error: percent array length greater than time array length",
			vsys: &mocksys.VCore{
				ListFn: func() ([]float64, []cpu.TimesStat, error) {
					return []float64{50.0, 99.9},
						[]cpu.TimesStat{
							{
								CPU:    "cpu0",
								User:   111.1,
								System: 222.2,
								Idle:   333.3,
							},
						},
						nil
				},
			},
			id:         2,
			wantedData: nil,
			wantedErr:  echo.NewHTTPError(http.StatusAccepted, "results were not returned as they could not be guaranteed"),
		},
		{
			name: "error: id value greater than percent array length",
			vsys: &mocksys.VCore{
				ListFn: func() ([]float64, []cpu.TimesStat, error) {
					return []float64{50.0},
						[]cpu.TimesStat{
							{
								CPU:    "cpu0",
								User:   111.1,
								System: 222.2,
								Idle:   333.3,
							},
						},
						nil
				},
			},
			id:         999,
			wantedData: nil,
			wantedErr:  echo.NewHTTPError(http.StatusNotFound, "id out of range : this system has  1 vcores; count starts from 0"),
		},
		{
			name: "error: negative id",
			vsys: &mocksys.VCore{
				ListFn: func() ([]float64, []cpu.TimesStat, error) {
					return []float64{50.0},
						[]cpu.TimesStat{
							{
								CPU:    "cpu0",
								User:   111.1,
								System: 222.2,
								Idle:   333.3,
							},
						},
						nil
				},
			},
			id:         -999,
			wantedData: nil,
			wantedErr:  echo.NewHTTPError(http.StatusNotFound, "id out of range : this system has  1 vcores; count starts from 0"),
		},
		{
			name: "error: parsing id failed",
			vsys: &mocksys.VCore{
				ListFn: func() ([]float64, []cpu.TimesStat, error) {
					return []float64{99.9},
						[]cpu.TimesStat{
							{
								CPU:    "cpu",
								User:   111.1,
								System: 222.2,
								Idle:   333.3,
							},
						},
						nil
				},
			},
			id:         1,
			wantedData: nil,
			wantedErr:  echo.NewHTTPError(http.StatusInternalServerError, "parsing id was unsuccessful"),
		},
		{
			name: "success",
			vsys: &mocksys.VCore{
				ListFn: func() ([]float64, []cpu.TimesStat, error) {
					return []float64{99.9},
						[]cpu.TimesStat{
							{
								CPU:    "cpu0",
								User:   111.1,
								System: 222.2,
								Idle:   333.3,
							},
						},
						nil
				},
			},
			id: 0,
			wantedData: &rpi.VCore{
				ID:     0,
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
			s := New(tc.vsys)
			vcores, err := s.View(tc.id)
			assert.Equal(t, tc.wantedData, vcores)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
