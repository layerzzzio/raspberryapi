package vcore

import (
	"fmt"
	"testing"

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
			name:       "Testing with one single numerical character in input string",
			input:      "cpu0",
			min:        0,
			max:        9,
			wantedData: []string{"0"},
		},
		{
			name:       "Testing with three numerical characters next to each other in input string",
			input:      "cpu012",
			min:        0,
			max:        9,
			wantedData: []string{"012"},
		},
		{
			name:       "Testing with four numerical characters apart from each other in input string",
			input:      "0cpu12xxx3xxx9",
			min:        0,
			max:        9,
			wantedData: []string{"0", "12", "3", "9"},
		},
		{
			name:       "Testing without numerical characters in input string",
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
			name:       "Error case",
			input:      []string{"X"},
			wantedData: -1,
			wantedErr:  fmt.Errorf("test error"),
		},
		{
			name:       "One single numerical character in input string",
			input:      []string{"0"},
			wantedData: 0,
			wantedErr:  nil,
		},
		{
			name:       "Three numerical characters next to each other in input string",
			input:      []string{"012"},
			wantedData: 12,
			wantedErr:  nil,
		},
		{
			name:       "Four numerical characters next to each other in input string",
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
			name:       "Error case",
			input:      []string{"X"},
			wantedData: -1,
			wantedErr:  fmt.Errorf("test error"),
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

func TestListError(t *testing.T) {
	cases := []struct {
		name string
		vsys *mocksys.VCore
	}{
		{
			name: "Length array info greater than array percent & time",
			vsys: &mocksys.VCore{
				ListFn: func() ([]float64, []cpu.TimesStat, error) {
					return nil, nil, nil
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := New(tc.vsys)
			vcores, err := s.List()
			fmt.Println(vcores)
			fmt.Println(err)
			// assert.Nil(t, vcores)
			// assert.NotNil(t, err)
			// assert.EqualValues(t, "Results were not returned as they could not be guaranteed", err.(*echo.HTTPError).Message)
			// assert.EqualValues(t, http.StatusAccepted, err.(*echo.HTTPError).Code)
		})
	}
}
