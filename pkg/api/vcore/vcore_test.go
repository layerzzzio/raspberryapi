package vcore

import (
	"errors"
	"fmt"
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
			name:       "testing with one single numerical character in input string",
			input:      "cpu0",
			min:        0,
			max:        9,
			wantedData: []string{"0"},
		},
		{
			name:       "testing with three numerical characters next to each other in input string",
			input:      "cpu012",
			min:        0,
			max:        9,
			wantedData: []string{"012"},
		},
		{
			name:       "testing with four numerical characters apart from each other in input string",
			input:      "0cpu12xxx3xxx9",
			min:        0,
			max:        9,
			wantedData: []string{"0", "12", "3", "9"},
		},
		{
			name:       "testing without numerical characters in input string",
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

func TestConcatIDSuccess(t *testing.T) {
	cases := []struct {
		name       string
		input      []string
		wantedData int
		wantedErr  error
	}{
		{
			name:       "one single numerical character in input string",
			input:      []string{"0"},
			wantedData: 0,
			wantedErr:  nil,
		},
		{
			name:       "three numerical characters next to each other in input string",
			input:      []string{"012"},
			wantedData: 12,
			wantedErr:  nil,
		},
		{
			name:       "four numerical characters next to each other in input string",
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

func TestConcatIDError(t *testing.T) {
	cases := []struct {
		name       string
		input      []string
		wantedData int
		wantedErr  error
	}{
		{
			name:       "invalid argument type",
			input:      []string{"X"},
			wantedData: -1,
			wantedErr:  errors.New("invalid syntax"),
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

func TestListNilMetricsError(t *testing.T) {
	vsys := &mocksys.VCore{
		ListFn: func() ([]float64, []cpu.TimesStat, error) {
			return nil, nil, echo.NewHTTPError(http.StatusAccepted, "results were not returned as they could not be guaranteed")
		},
	}

	wantedErr := echo.NewHTTPError(http.StatusAccepted, "results were not returned as they could not be guaranteed")

	s := New(vsys)
	vcores, err := s.List()
	assert.Nil(t, vcores)
	assert.NotNil(t, err)
	assert.Equal(t, err, wantedErr)
}

func TestListLengthDiffError(t *testing.T) {
	cases := []struct {
		name string
		vsys *mocksys.VCore
	}{
		{
			name: "length array percent greater than time array",
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
		},
		{
			name: "length array percent greater than time array",
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
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := New(tc.vsys)
			vcores, err := s.List()
			assert.Nil(t, vcores)
			assert.NotNil(t, err)
			assert.EqualValues(t, "results were not returned as they could not be guaranteed", err.(*echo.HTTPError).Message)
			assert.EqualValues(t, http.StatusAccepted, err.(*echo.HTTPError).Code)
		})
	}
}

func TestListParseIdError(t *testing.T) {
	vsys := &mocksys.VCore{
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
	}

	wantedErr := echo.NewHTTPError(http.StatusInternalServerError, "parsing id was unsuccessful")

	s := New(vsys)
	vcores, err := s.List()
	assert.Nil(t, vcores)
	assert.NotNil(t, err)
	assert.Equal(t, err, wantedErr)
}

func TestListSuccess(t *testing.T) {
	vsys := &mocksys.VCore{
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
	}

	wantedData := rpi.VCore{
		ID:     0,
		Used:   99.9,
		User:   111.1,
		System: 222.2,
		Idle:   333.3,
	}

	s := New(vsys)
	vcores, err := s.List()
	assert.Nil(t, err)
	assert.NotNil(t, vcores)
	assert.Equal(t, vcores[0], wantedData)
}

func TestViewNilMetricsError(t *testing.T) {
	vsys := &mocksys.VCore{
		ListFn: func() ([]float64, []cpu.TimesStat, error) {
			return nil, nil, echo.NewHTTPError(http.StatusAccepted, "results were not returned as they could not be guaranteed")
		},
	}

	wantedErr := echo.NewHTTPError(http.StatusAccepted, "results were not returned as they could not be guaranteed")

	s := New(vsys)
	id := 2
	vcores, err := s.View(id)
	assert.Nil(t, vcores)
	assert.NotNil(t, err)
	assert.Equal(t, err, wantedErr)
}

func TestViewLengthDiffError(t *testing.T) {
	cases := []struct {
		name string
		vsys *mocksys.VCore
	}{
		{
			name: "length array percent greater than time array",
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
		},
		{
			name: "length array percent greater than time array",
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
		},
	}

	id := 2
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := New(tc.vsys)
			vcores, err := s.View(id)
			assert.Nil(t, vcores)
			assert.NotNil(t, err)
			assert.EqualValues(t, "results were not returned as they could not be guaranteed", err.(*echo.HTTPError).Message)
			assert.EqualValues(t, http.StatusAccepted, err.(*echo.HTTPError).Code)
		})
	}
}

func TestViewSuccess(t *testing.T) {
	vsys := &mocksys.VCore{
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
	}

	wantedData := &rpi.VCore{
		ID:     0,
		Used:   99.9,
		User:   111.1,
		System: 222.2,
		Idle:   333.3,
	}

	s := New(vsys)
	id := 0
	vcores, err := s.View(id)
	assert.Nil(t, err)
	fmt.Println(vcores)
	assert.NotNil(t, vcores)
	assert.Equal(t, vcores, wantedData)
}
