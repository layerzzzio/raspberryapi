package metrics_test

import (
	"testing"
	"time"

	"github.com/raspibuddy/rpi/pkg/utl/metrics"
	"github.com/raspibuddy/rpi/pkg/utl/mock"
	"github.com/shirou/gopsutil/process"
	"github.com/stretchr/testify/assert"
)

func TestProcesses(t *testing.T) {
	cases := []struct {
		name       string
		id         int32
		mps        mock.Metrics
		wantedData metrics.PInfo
		wantedErr  error
	}{
		{
			name: "success: without process id",
			// id -1 simulate a non-existent process id
			id: -1,
			mps: mock.Metrics{
				PsPIDFn: func(p *process.Process, c chan (int32)) {
					c <- 99
				},
				PsNameFn: func(p *process.Process, c chan (string)) {
					c <- "process_99"
				},
				PsCPUPerFn: func(p *process.Process, c chan (float64)) {
					c <- 1.1
				},
				PsMemPerFn: func(p *process.Process, c chan (float32)) {
					c <- 2.2
				},
			},
			wantedData: metrics.PInfo{
				ID:         99,
				Name:       "process_99",
				CPUPercent: 1.1,
				MemPercent: 2.2,
			},
			wantedErr: nil,
		},
		// {
		// 	name:      "error: process id 0",
		// 	id:        0,
		// 	mps:       mock.Metrics{},
		// 	wantedErr: errors.New("process not found"),
		// 	//wantedErr: errors.Is("process not found"),
		// },
		{
			name: "success: with process id",
			// make sure to keep id 1 to make this test work
			id: 1,
			mps: mock.Metrics{
				PsPIDFn: func(p *process.Process, c chan (int32)) {
					c <- 1
				},
				PsNameFn: func(p *process.Process, c chan (string)) {
					c <- "process_1"
				},
				PsCPUPerFn: func(p *process.Process, c chan (float64)) {
					c <- 1.1
				},
				PsMemPerFn: func(p *process.Process, c chan (float32)) {
					c <- 2.2
				},
				PsUsernameFn: func(p *process.Process, c chan (string)) {
					c <- "pi"
				},
				PsCmdLineFn: func(p *process.Process, c chan (string)) {
					c <- "/cmd/test"
				},
				PsStatusFn: func(p *process.Process, c chan (string)) {
					c <- "S"
				},
				PsCreationTimeFn: func(p *process.Process, c chan (time.Time)) {
					c <- time.Time{}.Add(1888)
				},
				PsForegroundFn: func(p *process.Process, c chan (bool)) {
					c <- true
				},
				PsBackgroundFn: func(p *process.Process, c chan (bool)) {
					c <- false
				},
				PsIsRunningFn: func(p *process.Process, c chan (bool)) {
					c <- true
				},
				PsParentFn: func(p *process.Process, c chan (int32)) {
					c <- -1
				},
			},
			wantedData: metrics.PInfo{
				ID:           1,
				Name:         "process_1",
				CPUPercent:   1.1,
				MemPercent:   2.2,
				Username:     "pi",
				CommandLine:  "/cmd/test",
				Status:       "S",
				CreationTime: time.Time{}.Add(1888),
				Foreground:   true,
				Background:   false,
				IsRunning:    true,
				ParentP:      -1,
			},
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := metrics.New(tc.mps)

			var ps []metrics.PInfo
			var err error

			if tc.id >= 0 {
				ps, err = s.Processes(tc.id)
			} else {
				ps, err = s.Processes()
			}

			assert.Equal(t, tc.wantedData, ps[0])
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
