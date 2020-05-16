package sys

import (
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/process"
	"github.com/raspibuddy/rpi/pkg/utl/metrics"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	cases := []struct {
		name       string
		pinfo      []metrics.PInfo
		wantedData []rpi.Process
		wantedErr  error
	}{
		{
			name: "success",
			pinfo: []metrics.PInfo{
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
			wantedData: []rpi.Process{
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
			s := process.PSYS(Process{})
			ps, err := s.List(tc.pinfo)
			assert.Equal(t, tc.wantedData, ps)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}

func TestView(t *testing.T) {
	cases := []struct {
		name       string
		id         int32
		pinfo      []metrics.PInfo
		wantedData rpi.Process
		wantedErr  error
	}{
		{
			name: "success",
			id:   99,
			pinfo: []metrics.PInfo{
				{
					ID:           int32(99),
					Name:         "process_99",
					CPUPercent:   1.1,
					MemPercent:   2.2,
					Username:     "pi",
					CommandLine:  "/cmd/text",
					Status:       "S",
					CreationTime: 1666666,
					Foreground:   true,
					Background:   false,
					IsRunning:    true,
					ParentP:      int32(1),
				},
				{
					ID:           int32(100),
					Name:         "process_100",
					CPUPercent:   1.1,
					MemPercent:   2.2,
					Username:     "pi",
					CommandLine:  "/cmd/text",
					Status:       "S",
					CreationTime: 1666666,
					Foreground:   true,
					Background:   false,
					IsRunning:    true,
					ParentP:      int32(1),
				},
			},
			wantedData: rpi.Process{
				ID:           int32(99),
				Name:         "process_99",
				CPUPercent:   1.1,
				MemPercent:   2.2,
				Username:     "pi",
				CommandLine:  "/cmd/text",
				Status:       "S",
				CreationTime: 1666666,
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
			s := process.PSYS(Process{})
			ps, err := s.View(tc.id, tc.pinfo)
			assert.Equal(t, tc.wantedData, ps)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
