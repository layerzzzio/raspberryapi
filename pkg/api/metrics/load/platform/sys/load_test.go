package sys_test

import (
	"testing"

	"github.com/raspibuddy/rpi"
	"github.com/raspibuddy/rpi/pkg/api/metrics/load"
	"github.com/raspibuddy/rpi/pkg/api/metrics/load/platform/sys"
	lext "github.com/shirou/gopsutil/load"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	cases := []struct {
		name       string
		temp       lext.AvgStat
		procs      lext.MiscStat
		wantedData rpi.Load
		wantedErr  error
	}{
		{
			name: "success: temp is empty",
			procs: lext.MiscStat{
				ProcsTotal:   333,
				ProcsRunning: 222,
				ProcsBlocked: 111,
			},
			wantedData: rpi.Load{
				Load1:        0,
				Load5:        0,
				Load15:       0,
				ProcsTotal:   333,
				ProcsRunning: 222,
				ProcsBlocked: 111,
			},
			wantedErr: nil,
		},
		{
			name: "success: procs is empty",
			temp: lext.AvgStat{
				Load1:  1.1,
				Load5:  5.5,
				Load15: 15.15,
			},
			wantedData: rpi.Load{
				Load1:        1.1,
				Load5:        5.5,
				Load15:       15.15,
				ProcsTotal:   0,
				ProcsRunning: 0,
				ProcsBlocked: 0,
			},
			wantedErr: nil,
		},
		{
			name: "success",
			temp: lext.AvgStat{
				Load1:  1.1,
				Load5:  5.5,
				Load15: 15.15,
			},
			procs: lext.MiscStat{
				ProcsTotal:   333,
				ProcsRunning: 222,
				ProcsBlocked: 111,
			},
			wantedData: rpi.Load{
				Load1:        1.1,
				Load5:        5.5,
				Load15:       15.15,
				ProcsTotal:   333,
				ProcsRunning: 222,
				ProcsBlocked: 111,
			},
			wantedErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := load.LSYS(sys.Load{})
			loads, err := s.List(tc.temp, tc.procs)
			assert.Equal(t, tc.wantedData, loads)
			assert.Equal(t, tc.wantedErr, err)
		})
	}
}
